package nazuna

import (
	"fmt"
	"net/url"
	"regexp"
	"sync"

	"github.com/callummance/nazuna/messages"
	"github.com/callummance/nazuna/restclient"
	"github.com/callummance/nazuna/webhooklistener"
	"github.com/sirupsen/logrus"
)

type NazunaOpts struct {
	WebhookPath    string
	ListenOn       string
	ClientID       string
	ClientSecret   string
	Scopes         []string
	Secret         string
	ServerHostname string
}

type EventsubClient struct {
	listener      webhooklistener.Listener
	restClient    restclient.Client
	handlersLock  sync.RWMutex
	handlers      []webhooklistener.WebhookHandler
	transportOpts messages.TransportOpts
}

func NewClient(opts NazunaOpts) (*EventsubClient, error) {
	//Create listener
	var listener *webhooklistener.Listener
	var err error
	if opts.Secret == "" {
		listener, err = webhooklistener.NewListener()
		opts.Secret = listener.Secret()
	} else {
		listener, err = webhooklistener.NewListenerWithSecret(opts.Secret)
	}
	if err != nil {
		return nil, err
	}

	//Create REST client
	restclient := restclient.InitClient(opts.ClientID, opts.ClientSecret, opts.Scopes)

	//Build transport definition
	callbackURL, err := url.Parse(opts.ServerHostname)
	if err != nil {
		return nil, err
	}
	callbackURL.Path = opts.WebhookPath
	transport := messages.TransportOpts{
		Method:   "webhook",
		Callback: callbackURL.String(),
		Secret:   opts.Secret,
	}

	client := EventsubClient{
		listener:      *listener,
		restClient:    *restclient,
		transportOpts: transport,
	}

	go client.dispatchMessages()
	err = client.listener.Listen(opts.WebhookPath, opts.ListenOn)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

//RegisterHandler adds a handler function to the handlers slice
func (c *EventsubClient) RegisterHandler(handler interface{}) {
	switch v := handler.(type) {
	case func(*messages.Subscription, *messages.ChannelUpdateEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelUpdateHandler(v))
	case func(*messages.Subscription, *messages.ChannelFollowEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelFollowHandler(v))
	case func(*messages.Subscription, *messages.ChannelSubscribeEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelSubscribeHandler(v))
	case func(*messages.Subscription, *messages.ChannelCheerEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelCheerHandler(v))
	case func(*messages.Subscription, *messages.ChannelRaidEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelRaidHandler(v))
	case func(*messages.Subscription, *messages.ChannelBanEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelBanHandler(v))
	case func(*messages.Subscription, *messages.ChannelUnbanEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelUnbanHandler(v))
	case func(*messages.Subscription, *messages.ChannelPointsCustomRewardAddEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelPointsCustomRewardAddHandler(v))
	case func(*messages.Subscription, *messages.ChannelPointsCustomRewardUpdateEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelPointsCustomRewardUpdateHandler(v))
	case func(*messages.Subscription, *messages.ChannelPointsCustomRewardRemoveEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelPointsCustomRewardRemoveHandler(v))
	case func(*messages.Subscription, *messages.ChannelPointsCustomRewardRedemptionAddEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelPointsCustomRewardRedemptionAddHandler(v))
	case func(*messages.Subscription, *messages.ChannelPointsCustomRewardRedemptionUpdateEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelPointsCustomRewardRedemptionUpdateHandler(v))
	case func(*messages.Subscription, *messages.ChannelHypeTrainBeginEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelHypeTrainBeginHandler(v))
	case func(*messages.Subscription, *messages.ChannelHypeTrainProgressEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelHypeTrainProgressHandler(v))
	case func(*messages.Subscription, *messages.ChannelHypeTrainEndEvent):
		c.handlers = append(c.handlers, webhooklistener.ChannelHypeTrainEndHandler(v))
	case func(*messages.Subscription, *messages.StreamOnlineEvent):
		c.handlers = append(c.handlers, webhooklistener.StreamOnlineHandler(v))
	case func(*messages.Subscription, *messages.StreamOfflineEvent):
		c.handlers = append(c.handlers, webhooklistener.StreamOfflineHandler(v))
	case func(*messages.Subscription, *messages.UserAuthorizationRevokeEvent):
		c.handlers = append(c.handlers, webhooklistener.UserAuthorizationRevokeHandler(v))
	case func(*messages.Subscription, *messages.UserUpdateEvent):
		c.handlers = append(c.handlers, webhooklistener.UserUpdateHandler(v))
	}
}

func (c *EventsubClient) CreateSubscription(condition interface{}) (*messages.SubscriptionRequestStatus, error) {
	return c.restClient.CreateSubscription(condition, c.transportOpts)
}

func (c *EventsubClient) Subscriptions(filters restclient.SubscriptionsParams) chan restclient.SubscriptionResult {
	return c.restClient.Subscriptions(&filters)
}

func (c *EventsubClient) DeleteSubscription(subscriptionID string) error {
	return c.restClient.DeleteSubscription(subscriptionID)
}

func (c *EventsubClient) ClearSubscriptions() error {
	for res := range c.Subscriptions(restclient.SubscriptionsParams{}) {
		if res.Err != nil {
			return res.Err
		}
		c.DeleteSubscription(res.Subscription.ID)
	}
	return nil
}

func (c *EventsubClient) GetUsers(ids, names []string) ([]restclient.TwitchUser, error) {
	return c.restClient.GetUsers(ids, names)
}

var broadcasterUrlRegex = regexp.MustCompile(`(?:https?://)?(?:(?:www|go|m)\.)?twitch\.tv/(?P<username>[a-zA-Z0-9_]{4,25})`)

func (c *EventsubClient) GetBroadcaster(urlOrName string) (*restclient.TwitchUser, error) {
	matches := broadcasterUrlRegex.FindStringSubmatch(urlOrName)
	switch {
	case matches == nil:
		//Regex did not match, so assume we have a username directly
		users, err := c.GetUsers([]string{}, []string{urlOrName})
		if err != nil {
			return nil, fmt.Errorf("%v does not appear to be a twitch url, so assuming it is a username; fetching user data failed due to %v", urlOrName, err)
		}
		return &users[0], nil
	case matches[1] != "":
		//Regex matches, so we have a url
		username := matches[1]
		users, err := c.GetUsers([]string{}, []string{username})
		if err != nil {
			return nil, fmt.Errorf("extracted username %v from the provided twitch url; fetching user data failed due to %v", username, err)
		}
		return &users[0], nil
	default:
		return nil, fmt.Errorf("regex matching failed whilst trying to get broadcaster data for %v", urlOrName)
	}
}

func (c *EventsubClient) dispatchMessages() {
	for {
		select {
		case msg, open := <-c.listener.NotificationsChannel():
			if open {
				logrus.Debugf("Dispatching message %v", msg)
				c.dispatchMessage(msg)
			} else {
				logrus.Info("Stopping message dispatch due to closed channel")
				return
			}
		}
	}
}

func (c *EventsubClient) dispatchMessage(message messages.EventNotificationMessage) {
	c.handlersLock.RLock()
	defer c.handlersLock.RUnlock()
	for _, handler := range c.handlers {
		go handler.Handle(message)
	}
}
