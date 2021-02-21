package webhooklistener

import "github.com/callummance/nazuna/messages"

type WebhookHandler interface {
	Type() string
	Handle(msg messages.EventNotificationMessage)
}

//ChannelUpdateHandler represents a handler for webhook messages of type ChannelUpdateEvent
type ChannelUpdateHandler func(*messages.Subscription, *messages.ChannelUpdateEvent)

//Type returns the string representing ChannelUpdate's event name on the Twitch API
func (h ChannelUpdateHandler) Type() string {
	return messages.SubscriptionChannelUpdate
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelUpdateHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelUpdateEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelFollowHandler represents a handler for webhook messages of type ChannelFollowEvent
type ChannelFollowHandler func(*messages.Subscription, *messages.ChannelFollowEvent)

//Type returns the string representing ChannelFollow's event name on the Twitch API
func (h ChannelFollowHandler) Type() string {
	return messages.SubscriptionChannelFollow
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelFollowHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelFollowEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelSubscribeHandler represents a handler for webhook messages of type ChannelSubscribeEvent
type ChannelSubscribeHandler func(*messages.Subscription, *messages.ChannelSubscribeEvent)

//Type returns the string representing ChannelSubscribe's event name on the Twitch API
func (h ChannelSubscribeHandler) Type() string {
	return messages.SubscriptionChannelSubscribe
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelSubscribeHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelSubscribeEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelCheerHandler represents a handler for webhook messages of type ChannelCheerEvent
type ChannelCheerHandler func(*messages.Subscription, *messages.ChannelCheerEvent)

//Type returns the string representing ChannelCheer's event name on the Twitch API
func (h ChannelCheerHandler) Type() string {
	return messages.SubscriptionChannelCheer
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelCheerHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelCheerEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelRaidHandler represents a handler for webhook messages of type ChannelRaidEvent
type ChannelRaidHandler func(*messages.Subscription, *messages.ChannelRaidEvent)

//Type returns the string representing ChannelRaid's event name on the Twitch API
func (h ChannelRaidHandler) Type() string {
	return messages.SubscriptionChannelRaid
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelRaidHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelRaidEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelBanHandler represents a handler for webhook messages of type ChannelBanEvent
type ChannelBanHandler func(*messages.Subscription, *messages.ChannelBanEvent)

//Type returns the string representing ChannelBan's event name on the Twitch API
func (h ChannelBanHandler) Type() string {
	return messages.SubscriptionChannelBan
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelBanHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelBanEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelUnbanHandler represents a handler for webhook messages of type ChannelUnbanEvent
type ChannelUnbanHandler func(*messages.Subscription, *messages.ChannelUnbanEvent)

//Type returns the string representing ChannelUnban's event name on the Twitch API
func (h ChannelUnbanHandler) Type() string {
	return messages.SubscriptionChannelUnban
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelUnbanHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelUnbanEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelPointsCustomRewardAddHandler represents a handler for webhook messages of type ChannelPointsCustomRewardAddEvent
type ChannelPointsCustomRewardAddHandler func(*messages.Subscription, *messages.ChannelPointsCustomRewardAddEvent)

//Type returns the string representing ChannelPointsCustomRewardAdd's event name on the Twitch API
func (h ChannelPointsCustomRewardAddHandler) Type() string {
	return messages.SubscriptionChannelPointsCustomRewardAdd
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelPointsCustomRewardAddHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelPointsCustomRewardAddEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelPointsCustomRewardUpdateHandler represents a handler for webhook messages of type ChannelPointsCustomRewardUpdateEvent
type ChannelPointsCustomRewardUpdateHandler func(*messages.Subscription, *messages.ChannelPointsCustomRewardUpdateEvent)

//Type returns the string representing ChannelPointsCustomRewardUpdate's event name on the Twitch API
func (h ChannelPointsCustomRewardUpdateHandler) Type() string {
	return messages.SubscriptionChannelPointsCustomRewardUpdate
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelPointsCustomRewardUpdateHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelPointsCustomRewardUpdateEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelPointsCustomRewardRemoveHandler represents a handler for webhook messages of type ChannelPointsCustomRewardRemoveEvent
type ChannelPointsCustomRewardRemoveHandler func(*messages.Subscription, *messages.ChannelPointsCustomRewardRemoveEvent)

//Type returns the string representing ChannelPointsCustomRewardRemove's event name on the Twitch API
func (h ChannelPointsCustomRewardRemoveHandler) Type() string {
	return messages.SubscriptionChannelPointsCustomRewardRemove
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelPointsCustomRewardRemoveHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelPointsCustomRewardRemoveEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelPointsCustomRewardRedemptionAddHandler represents a handler for webhook messages of type ChannelPointsCustomRewardRedemptionAddEvent
type ChannelPointsCustomRewardRedemptionAddHandler func(*messages.Subscription, *messages.ChannelPointsCustomRewardRedemptionAddEvent)

//Type returns the string representing ChannelPointsCustomRewardRedemptionAdd's event name on the Twitch API
func (h ChannelPointsCustomRewardRedemptionAddHandler) Type() string {
	return messages.SubscriptionChannelPointsCustomRewardRedemptionAdd
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelPointsCustomRewardRedemptionAddHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelPointsCustomRewardRedemptionAddEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelPointsCustomRewardRedemptionUpdateHandler represents a handler for webhook messages of type ChannelPointsCustomRewardRedemptionUpdateEvent
type ChannelPointsCustomRewardRedemptionUpdateHandler func(*messages.Subscription, *messages.ChannelPointsCustomRewardRedemptionUpdateEvent)

//Type returns the string representing ChannelPointsCustomRewardRedemptionUpdate's event name on the Twitch API
func (h ChannelPointsCustomRewardRedemptionUpdateHandler) Type() string {
	return messages.SubscriptionChannelPointsCustomRewardRedemptionUpdate
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelPointsCustomRewardRedemptionUpdateHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelPointsCustomRewardRedemptionUpdateEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelHypeTrainBeginHandler represents a handler for webhook messages of type ChannelHypeTrainBeginEvent
type ChannelHypeTrainBeginHandler func(*messages.Subscription, *messages.ChannelHypeTrainBeginEvent)

//Type returns the string representing ChannelHypeTrainBegin's event name on the Twitch API
func (h ChannelHypeTrainBeginHandler) Type() string {
	return messages.SubscriptionChannelHypeTrainBegin
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelHypeTrainBeginHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelHypeTrainBeginEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelHypeTrainProgressHandler represents a handler for webhook messages of type ChannelHypeTrainProgressEvent
type ChannelHypeTrainProgressHandler func(*messages.Subscription, *messages.ChannelHypeTrainProgressEvent)

//Type returns the string representing ChannelHypeTrainProgress's event name on the Twitch API
func (h ChannelHypeTrainProgressHandler) Type() string {
	return messages.SubscriptionChannelHypeTrainProgress
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelHypeTrainProgressHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelHypeTrainProgressEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//ChannelHypeTrainEndHandler represents a handler for webhook messages of type ChannelHypeTrainEndEvent
type ChannelHypeTrainEndHandler func(*messages.Subscription, *messages.ChannelHypeTrainEndEvent)

//Type returns the string representing ChannelHypeTrainEnd's event name on the Twitch API
func (h ChannelHypeTrainEndHandler) Type() string {
	return messages.SubscriptionChannelHypeTrainEnd
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h ChannelHypeTrainEndHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.ChannelHypeTrainEndEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//StreamOnlineHandler represents a handler for webhook messages of type StreamOnlineEvent
type StreamOnlineHandler func(*messages.Subscription, *messages.StreamOnlineEvent)

//Type returns the string representing StreamOnline's event name on the Twitch API
func (h StreamOnlineHandler) Type() string {
	return messages.SubscriptionStreamOnline
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h StreamOnlineHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.StreamOnlineEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//StreamOfflineHandler represents a handler for webhook messages of type StreamOfflineEvent
type StreamOfflineHandler func(*messages.Subscription, *messages.StreamOfflineEvent)

//Type returns the string representing StreamOffline's event name on the Twitch API
func (h StreamOfflineHandler) Type() string {
	return messages.SubscriptionStreamOffline
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h StreamOfflineHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.StreamOfflineEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//UserAuthorizationRevokeHandler represents a handler for webhook messages of type UserAuthorizationRevokeEvent
type UserAuthorizationRevokeHandler func(*messages.Subscription, *messages.UserAuthorizationRevokeEvent)

//Type returns the string representing UserAuthorizationRevoke's event name on the Twitch API
func (h UserAuthorizationRevokeHandler) Type() string {
	return messages.SubscriptionUserAuthorizationRevoke
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h UserAuthorizationRevokeHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.UserAuthorizationRevokeEvent); ok {
		h(&msg.Subscription, ev)
	}
}

//UserUpdateHandler represents a handler for webhook messages of type UserUpdateEvent
type UserUpdateHandler func(*messages.Subscription, *messages.UserUpdateEvent)

//Type returns the string representing UserUpdate's event name on the Twitch API
func (h UserUpdateHandler) Type() string {
	return messages.SubscriptionUserUpdate
}

//Handle passes on a message to the handler function if it is of the correct type.
func (h UserUpdateHandler) Handle(msg messages.EventNotificationMessage) {
	if ev, ok := msg.Event.(*messages.UserUpdateEvent); ok {
		h(&msg.Subscription, ev)
	}
}
