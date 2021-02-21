package messages

import "time"

const (
	SubscriptionChannelUpdate                             = "channel.update"
	SubscriptionChannelFollow                             = "channel.follow"
	SubscriptionChannelSubscribe                          = "channel.subscribe"
	SubscriptionChannelCheer                              = "channel.cheer"
	SubscriptionChannelRaid                               = "channel.raid"
	SubscriptionChannelBan                                = "channel.ban"
	SubscriptionChannelUnban                              = "channel.unban"
	SubscriptionChannelPointsCustomRewardAdd              = "channel.channel_points_custom_reward.add"
	SubscriptionChannelPointsCustomRewardUpdate           = "channel.channel_points_custom_reward.update"
	SubscriptionChannelPointsCustomRewardRemove           = "channel.channel_points_custom_reward.remove"
	SubscriptionChannelPointsCustomRewardRedemptionAdd    = "channel.channel_points_custom_reward_redemption.add"
	SubscriptionChannelPointsCustomRewardRedemptionUpdate = "channel.channel_points_custom_reward_redemption.update"
	SubscriptionChannelHypeTrainBegin                     = "channel.hype_train.begin"
	SubscriptionChannelHypeTrainProgress                  = "channel.hype_train.progress"
	SubscriptionChannelHypeTrainEnd                       = "channel.hype_train.end"
	SubscriptionStreamOnline                              = "stream.online"
	SubscriptionStreamOffline                             = "stream.offline"
	SubscriptionUserAuthorizationRevoke                   = "user.authorization.revoke"
	SubscriptionUserUpdate                                = "user.update"
)

type Subscription struct {
	ID        string        `json:"id,omitempty"`
	Status    string        `json:"status,omitempty"`
	Type      string        `json:"type"`
	Version   string        `json:"version"`
	Condition interface{}   `json:"condition"`
	Transport TransportOpts `json:"transport"`
	CreatedAt *time.Time    `json:"created_at,omitempty"`
}

type TransportOpts struct {
	Method   string `json:"method"`
	Callback string `json:"callback"`
	Secret   string `json:"secret"`
}

type SubscriptionRequestStatus struct {
	Data  []Subscription `json:"data"`
	Total int            `json:"total"`
	Limit int            `json:"limit"`
}

type VerificationMessage struct {
	Challenge    string       `json:"challenge"`
	Subscription Subscription `json:"subscription"`
}

type ConditionChannelUpdate struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelFollow struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelSubscribe struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelCheer struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelRaid struct {
	FromBroadcasterUID string `json:"from_broadcaster_user_id,omitempty"`
	ToBroadcasterUID   string `json:"to_broadcaster_user_id,omitempty"`
}

type ConditionChannelBan struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelUnban struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelPointsCustomRewardAdd struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelPointsCustomRewardUpdate struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
	RewardID       string `json:"reward_id,omitempty"`
}

type ConditionChannelPointsCustomRewardRemove struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
	RewardID       string `json:"reward_id,omitempty"`
}

type ConditionChannelPointsCustomRewardRedemptionAdd struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
	RewardID       string `json:"reward_id,omitempty"`
}

type ConditionChannelPointsCustomRewardRedemptionUpdate struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
	RewardID       string `json:"reward_id,omitempty"`
}

type ConditionChannelHypeTrainBegin struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelHypeTrainProgress struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionChannelHypeTrainEnd struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionStreamOnline struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionStreamOffline struct {
	BroadcasterUID string `json:"broadcaster_user_id"`
}

type ConditionUserAuthorizationRevoke struct {
	ClientID string `json:"client_id"`
}

type ConditionUserUpdate struct {
	UserID string `json:"user_id"`
}
