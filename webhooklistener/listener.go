package webhooklistener

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/callummance/nazuna/messages"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var messageExpiry, _ = time.ParseDuration("10m")

const (
	messageIDCacheExpiry  = 24 * time.Hour
	messageIDCacheCleanup = time.Hour
)

type Listener struct {
	processedMessages    *cache.Cache
	secret               string
	notificationsChannel chan messages.EventNotificationMessage
	closeChannel         chan interface{}
	permissive           bool
}

func NewListenerWithSecret(secret string, permissive bool) (*Listener, error) {
	messageIDs := cache.New(messageIDCacheExpiry, messageIDCacheCleanup)
	notificationChannel := make(chan messages.EventNotificationMessage)
	closeChannel := make(chan interface{})
	return &Listener{
		processedMessages:    messageIDs,
		secret:               secret,
		notificationsChannel: notificationChannel,
		closeChannel:         closeChannel,
		permissive:           permissive,
	}, nil
}

func NewListener(permissive bool) (*Listener, error) {
	secret, err := GenSecret()
	if err != nil {
		return nil, err
	}
	return NewListenerWithSecret(secret, permissive)
}

func GenSecret() (string, error) {
	secretBytes := make([]byte, 100)
	_, err := rand.Read(secretBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(secretBytes)[0:49], nil
}

//Listen starts listening for incoming webhook calls at the pattern `webhookPath` on interface and port `listenOk`
func (l *Listener) Listen(webhookPath string, listenOn string) error {
	logrus.Infof("Starting server to listen for webhooks at path %v on address:port %v.", webhookPath, listenOn)
	listener, err := net.Listen("tcp4", listenOn)
	if err != nil {
		logrus.Errorf("Failed to start listening for webhooks due to error %v", err)
		return err
	}
	go func() {
		//TODO: listen on closeChannel to exit server when signalled
		mux := http.NewServeMux()
		mux.HandleFunc(webhookPath, l.handleWebhook)
		server := http.Server{
			Addr:    listenOn,
			Handler: mux,
		}
		logrus.Fatal(server.Serve(listener))
	}()
	return nil
}

//NotificationsChannel returns the channel upon which messages are returned.
func (l *Listener) NotificationsChannel() chan messages.EventNotificationMessage {
	return l.notificationsChannel
}

//NotificationsChannel returns the channel upon which messages are returned.
func (l *Listener) Secret() string {
	return l.secret
}

func (l *Listener) handleWebhook(w http.ResponseWriter, r *http.Request) {
	//Verify message is from twitch and get body
	body := l.verifyMessage(&w, r, l.secret)
	if body == nil {
		return
	}
	//Take note of the fact the message has been recieved
	msgID := strings.Join(r.Header["Twitch-Eventsub-Message-Id"], "")
	l.processedMessages.Set(msgID, nil, cache.DefaultExpiration)

	//Branch based on message type
	msgType := strings.Join(r.Header["Twitch-Eventsub-Message-Type"], "")
	switch msgType {
	case "webhook_callback_verification":
		//Verification message
		logrus.Debugf("Recieved verification message from twitch")
		logrus.Tracef("Recieved verification message from twitch: %q", body)
		var message messages.VerificationMessage
		err := json.Unmarshal(body, &message)
		if err != nil {
			logrus.Warnf("Failed to unmarshal webhook verification message from twitch")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logrus.Infof("Responding to twitch callback verification for subscription %v.", message.Subscription)
		challenge := message.Challenge
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s", challenge)
		return
	case "notification":
		//Actual notification message
		logrus.Tracef("Recieved notification from twitch: %q", body)
		subscriptionType := strings.Join(r.Header["Twitch-Eventsub-Subscription-Type"], "")
		message, err := decodeNotification(&body, subscriptionType)
		if err != nil {
			logrus.Warnf("Discarding message.")
			w.WriteHeader(http.StatusOK)
			return
		}
		l.notificationsChannel <- *message
		w.WriteHeader(http.StatusOK)
		return
	default:
		//Unknown message type
		logrus.Warnf("Recieved message with unknown message type %v from twitch: %v", msgType, body)
		return
	}
}

//Attempts to verify the signature, send time and unique ID of a message, returning the body contents iff successful.
func (l *Listener) verifyMessage(w *http.ResponseWriter, r *http.Request, secret string) []byte {
	msgID := strings.Join(r.Header["Twitch-Eventsub-Message-Id"], "")
	msgTimestamp := strings.Join(r.Header["Twitch-Eventsub-Message-Timestamp"], "")

	//Check message time
	oldestValidTime := time.Now().Add(-messageExpiry)
	messageTime, err := time.Parse(time.RFC3339, msgTimestamp)
	if err != nil {
		logrus.Warnf("Failed to decode message timestamp %v due to error %v", msgTimestamp, err)
		http.Error(*w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	if messageTime.Before(oldestValidTime) && !l.permissive {
		//Message is too old
		logrus.Infof("Discarded message because it was sent more than %v ago.", messageExpiry)
		http.Error(*w, fmt.Errorf("message was sent at %v, wheras only messages sent since %v are currently acceptible", messageTime, oldestValidTime).Error(), http.StatusBadRequest)
		return nil
	}

	//Check if we have seen message before
	_, seenBefore := l.processedMessages.Get(msgID)
	if seenBefore && !l.permissive {
		//Message is seen before
		logrus.Infof("Discarded message because it was recieved before.", messageExpiry)
		(*w).WriteHeader(http.StatusOK)
		return nil
	}

	var hmacBuf bytes.Buffer
	var bodyBuf bytes.Buffer
	hmacBuf.WriteString(msgID)
	hmacBuf.WriteString(msgTimestamp)
	io.Copy(&hmacBuf, io.TeeReader(r.Body, &bodyBuf))

	hasher := hmac.New(sha256.New, []byte(secret))
	_, err = hasher.Write(hmacBuf.Bytes())
	if err != nil {
		logrus.Warnf("Failed to copy bytes from request body to HMAC buf due to error %v", err)
		return nil
	}

	calculatedHash := hasher.Sum(nil)
	headerSig := strings.Join(r.Header["Twitch-Eventsub-Message-Signature"], "")
	providedHash, err := hex.DecodeString(strings.TrimPrefix(headerSig, "sha256="))
	if err != nil {
		logrus.Warnf("Provided HMAC signature '%v' was not valid hexadecimal: %v", headerSig, err)
	}

	if l.permissive || hmac.Equal(calculatedHash, providedHash) {
		return bodyBuf.Bytes()
	}
	http.Error(*w, "Hash did not match", http.StatusForbidden)
	return nil
}

type intermediateNotification struct {
	Subscription messages.Subscription `json:"subscription"`
	Event        json.RawMessage       `json:"event"`
}

func decodeNotification(body *[]byte, subscriptionType string) (*messages.EventNotificationMessage, error) {
	var intermediate intermediateNotification
	err := json.Unmarshal(*body, &intermediate)
	if err != nil {
		logrus.Warnf("Failed to unmarshal JSON message from Twitch %v due to error %v.", body, err)
		return nil, err
	}
	res := messages.EventNotificationMessage{
		Subscription: intermediate.Subscription,
		Event:        nil,
	}
	switch subscriptionType {
	case messages.SubscriptionChannelUpdate:
		var ev messages.ChannelUpdateEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelUpdate event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelFollow:
		var ev messages.ChannelFollowEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelFollow event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelSubscribe:
		var ev messages.ChannelSubscribeEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelSubscribe event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelCheer:
		var ev messages.ChannelCheerEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelCheer event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelRaid:
		var ev messages.ChannelRaidEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelRaid event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelBan:
		var ev messages.ChannelBanEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelBan event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelUnban:
		var ev messages.ChannelUnbanEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelUnban event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelPointsCustomRewardAdd:
		var ev messages.ChannelPointsCustomRewardAddEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelPointsCustomRewardAdd event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelPointsCustomRewardUpdate:
		var ev messages.ChannelPointsCustomRewardUpdateEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelPointsCustomRewardUpdate event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelPointsCustomRewardRemove:
		var ev messages.ChannelPointsCustomRewardRemoveEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelPointsCustomRewardRemove event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelPointsCustomRewardRedemptionAdd:
		var ev messages.ChannelPointsCustomRewardRedemptionAddEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelPointsCustomRewardRedemptionAdd event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelPointsCustomRewardRedemptionUpdate:
		var ev messages.ChannelPointsCustomRewardRedemptionUpdateEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelPointsCustomRewardRedemptionUpdate event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelHypeTrainBegin:
		var ev messages.ChannelHypeTrainBeginEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelHypeTrainBegin event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelHypeTrainProgress:
		var ev messages.ChannelHypeTrainProgressEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelHypeTrainProgress event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionChannelHypeTrainEnd:
		var ev messages.ChannelHypeTrainEndEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal ChannelHypeTrainEnd event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionStreamOnline:
		var ev messages.StreamOnlineEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal StreamOnline event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionStreamOffline:
		var ev messages.StreamOfflineEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal StreamOffline event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionUserAuthorizationRevoke:
		var ev messages.UserAuthorizationRevokeEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal UserAuthorizationRevoke event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	case messages.SubscriptionUserUpdate:
		var ev messages.UserUpdateEvent
		err := json.Unmarshal(intermediate.Event, &ev)
		if err != nil {
			logrus.Warnf("Failed to unmarshal UserUpdate event %v due to error %v", intermediate.Event, err)
			return nil, err
		}
		res.Event = &ev
	}
	return &res, nil
}
