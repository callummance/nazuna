package messages

import "time"

type EventNotificationMessage struct {
	Subscription Subscription `json:"subscription"`
	Event        interface{}  `json:"event"`
}

//ChannelUpdateEvent represents a channel.update event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelupdate`
//The channel.update subscription type sends notifications when a broadcaster updates the category, title, mature flag, or broadcast language for their channel.
//No authorization required.
type ChannelUpdateEvent struct {
	BroadcasterUID       string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Language             string `json:"lanugage"`
	CategoryID           string `json:"category_id"`
	CategoryName         string `json:"category_name"`
	IsMature             bool   `json:"is_mature"`
}

//ChannelFollowEvent represents a channel.follow event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelfollow`
//The channel.follow subscription type sends a notification when a specified channel receives a follow.
//No authorization required.
type ChannelFollowEvent struct {
	UserUID              string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUID       string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

//ChannelSubscribeEvent represents a channel.subscribe event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelsubscribe`
//The channel.subscribe subscription type sends a notification when a specified channel receives a subscriber. This does not include resubscribes.
//Must have channel:read:subscriptions scope.
type ChannelSubscribeEvent struct {
	UserUID              string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUID       string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Tier                 string `json:"tier"`
	IsGift               bool   `json:"is_gift"`
}

//ChannelCheerEvent represents a channel.cheer event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelcheer`
//The channel.cheer subscription type sends a notification when a user cheers on the specified channel.
//Must have bits:read scope.
type ChannelCheerEvent struct {
	IsAnonymous          bool   `json:"is_anonymous"`
	UserUID              string `json:"user_id,omitempty"`
	UserLogin            string `json:"user_login,omitempty"`
	UserName             string `json:"user_name,omitempty"`
	BroadcasterUID       string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	Message              string `json:"message"`
	Bits                 int    `json:"bits"`
}

//ChannelRaidEvent represents a channel.raid event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelraid-beta`
//The channel.raid subscription type sends a notification when a broadcaster raids another broadcaster’s channel.
//No authorization required.
type ChannelRaidEvent struct {
	FromBroadcasterUID       string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterUID         string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
	Viewers                  int    `json:"viewers"`
}

//ChannelBanEvent represents a channel.ban event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelban`
//The channel.ban subscription type sends a notification when a viewer is timed out or banned from the specified channel.
//Must have channel:moderate scope.
type ChannelBanEvent struct {
	UserUID              string    `json:"user_id"`
	UserLogin            string    `json:"user_login"`
	UserName             string    `json:"user_name"`
	BroadcasterUID       string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	ModeratorUID         string    `json:"moderator_user_id"`
	ModeratorUserLogin   string    `json:"moderator_user_login"`
	ModeratorUserName    string    `json:"moderator_user_name"`
	Reason               string    `json:"reason"`
	EndsAt               time.Time `json:"ends_at"`
	IsPermanent          bool      `json:"is_permanent"`
}

//ChannelUnbanEvent represents a channel.unban event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelunban`
//The channel.unban subscription type sends a notification when a viewer is unbanned from the specified channel.
//Must have channel:moderate scope.
type ChannelUnbanEvent struct {
	UserUID              string `json:"user_id"`
	UserLogin            string `json:"user_login"`
	UserName             string `json:"user_name"`
	BroadcasterUID       string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
	ModeratorUID         string `json:"moderator_user_id"`
	ModeratorUserLogin   string `json:"moderator_user_login"`
	ModeratorUserName    string `json:"moderator_user_name"`
}

//ChannelPointsCustomRewardAddEvent represents a channel.channel_points_custom_reward.add event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelchannel_points_custom_rewardadd`
//The channel.channel_points_custom_reward.add subscription type sends a notification when a custom channel points reward has been created for the specified channel.
//Must have channel:read:redemptions or channel:manage:redemptions scope.
type ChannelPointsCustomRewardAddEvent struct {
	ID                                string               `json:"id"`
	BroadcasterUID                    string               `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	IsEnabled                         bool                 `json:"is_enabled"`
	IsPaused                          bool                 `json:"is_paused"`
	IsInStock                         bool                 `json:"is_in_stock"`
	Title                             string               `json:"title"`
	Cost                              int                  `json:"cost"`
	Prompt                            string               `json:"prompt"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	CooldownExpires                   time.Time            `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  int                  `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      PointsRewardLimit    `json:"max_per_stream"`
	MaxPerUserPerStream               PointsRewardLimit    `json:"max_per_user_per_stream"`
	GlobalCooldown                    PointsRewardCooldown `json:"global_cooldown"`
	BackgroundColour                  string               `json:"background_color"`
	Image                             map[string]string    `json:"image"`
	DefaultImage                      map[string]string    `json:"default_image"`
}

type PointsRewardLimit struct {
	IsEnabled bool `json:"is_enabled"`
	Value     int  `json:"value"`
}

type PointsRewardCooldown struct {
	IsEnabled bool `json:"is_enabled"`
	Seconds   int  `json:"seconds"`
}

//ChannelPointsCustomRewardUpdateEvent represents a channel.channel_points_custom_reward.update event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelchannel_points_custom_rewardupdate`
//The channel.channel_points_custom_reward.update subscription type sends a notification when a custom channel points reward has been updated for the specified channel.
//Must have channel:read:redemptions or channel:manage:redemptions scope.
type ChannelPointsCustomRewardUpdateEvent struct {
	ID                                string               `json:"id"`
	BroadcasterUID                    string               `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	IsEnabled                         bool                 `json:"is_enabled"`
	IsPaused                          bool                 `json:"is_paused"`
	IsInStock                         bool                 `json:"is_in_stock"`
	Title                             string               `json:"title"`
	Cost                              int                  `json:"cost"`
	Prompt                            string               `json:"prompt"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	CooldownExpires                   time.Time            `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  int                  `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      PointsRewardLimit    `json:"max_per_stream"`
	MaxPerUserPerStream               PointsRewardLimit    `json:"max_per_user_per_stream"`
	GlobalCooldown                    PointsRewardCooldown `json:"global_cooldown"`
	BackgroundColour                  string               `json:"background_color"`
	Image                             map[string]string    `json:"image"`
	DefaultImage                      map[string]string    `json:"default_image"`
}

//ChannelPointsCustomRewardRemoveEvent represents a channel.channel_points_custom_reward.remove event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelchannel_points_custom_rewardremove`
//The channel.channel_points_custom_reward.remove subscription type sends a notification when a custom channel points reward has been removed from the specified channel.
//Must have channel:read:redemptions or channel:manage:redemptions scope.
type ChannelPointsCustomRewardRemoveEvent struct {
	ID                                string               `json:"id"`
	BroadcasterUID                    string               `json:"broadcaster_user_id"`
	BroadcasterUserLogin              string               `json:"broadcaster_user_login"`
	BroadcasterUserName               string               `json:"broadcaster_user_name"`
	IsEnabled                         bool                 `json:"is_enabled"`
	IsPaused                          bool                 `json:"is_paused"`
	IsInStock                         bool                 `json:"is_in_stock"`
	Title                             string               `json:"title"`
	Cost                              int                  `json:"cost"`
	Prompt                            string               `json:"prompt"`
	IsUserInputRequired               bool                 `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                 `json:"should_redemptions_skip_request_queue"`
	CooldownExpires                   time.Time            `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  int                  `json:"redemptions_redeemed_current_stream"`
	MaxPerStream                      PointsRewardLimit    `json:"max_per_stream"`
	MaxPerUserPerStream               PointsRewardLimit    `json:"max_per_user_per_stream"`
	GlobalCooldown                    PointsRewardCooldown `json:"global_cooldown"`
	BackgroundColour                  string               `json:"background_color"`
	Image                             map[string]string    `json:"image"`
	DefaultImage                      map[string]string    `json:"default_image"`
}

//ChannelPointsCustomRewardRedemptionAddEvent represents a channel.channel_points_custom_reward_redemption.add event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelchannel_points_custom_reward_redemptionadd`
//The channel.channel_points_custom_reward_redemption.add subscription type sends a notification when a viewer has redeemed a custom channel points reward on the specified channel.
//Must have channel:read:redemptions or channel:manage:redemptions scope.
type ChannelPointsCustomRewardRedemptionAddEvent struct {
	ID                   string             `json:"id"`
	BroadcasterUID       string             `json:"broadcaster_user_id"`
	BroadcasterUserLogin string             `json:"broadcaster_user_login"`
	BroadcasterUserName  string             `json:"broadcaster_user_name"`
	UserUID              string             `json:"user_id"`
	UserLogin            string             `json:"user_login"`
	UserName             string             `json:"user_name"`
	UserInput            string             `json:"user_input"`
	Status               string             `json:"status"`
	Reward               PointsCustomReward `json:"reward"`
	RedeemedAt           time.Time          `json:"redeemed_at"`
}

type PointsCustomReward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

//ChannelPointsCustomRewardRedemptionUpdateEvent represents a channel.channel_points_custom_reward_redemption.update event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelchannel_points_custom_reward_redemptionupdate`
//The channel.channel_points_custom_reward_redemption.update subscription type sends a notification when a redemption of a channel points custom reward has been updated for the specified channel.
//Must have channel:read:redemptions or channel:manage:redemptions scope.
type ChannelPointsCustomRewardRedemptionUpdateEvent struct {
	ID                   string             `json:"id"`
	BroadcasterUID       string             `json:"broadcaster_user_id"`
	BroadcasterUserLogin string             `json:"broadcaster_user_login"`
	BroadcasterUserName  string             `json:"broadcaster_user_name"`
	UserUID              string             `json:"user_id"`
	UserLogin            string             `json:"user_login"`
	UserName             string             `json:"user_name"`
	UserInput            string             `json:"user_input"`
	Status               string             `json:"status"`
	Reward               PointsCustomReward `json:"reward"`
	RedeemedAt           time.Time          `json:"redeemed_at"`
}

//ChannelHypeTrainBeginEvent represents a channel.hype_train.begin event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelhype_trainbegin`
//The channel.hype_train.begin subscription type sends a notification when a hype train begins on the specified channel. In addition to a channel.hype_train.begin event, one channel.hype_train.progress event will be sent for each contribution that caused the hype train to begin. EventSub does not make strong assurances about the order of message delivery, so it is possible to receive channel.hype_train.progress notifications before you receive the corresponding channel.hype_train.begin notification.
//After the hype train begins, any additional cheers or subscriptions on the channel will cause channel.hype_train.progress notifications to be sent. When the hype train is over, channel.hype_train.end is emitted.
//Must have channel:read:hype_train scope.
type ChannelHypeTrainBeginEvent struct {
	BroadcasterUID       string                 `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                 `json:"broadcaster_user_login"`
	BroadcasterUserName  string                 `json:"broadcaster_user_name"`
	Total                int                    `json:"total"`
	Progress             int                    `json:"progress"`
	Goal                 int                    `json:"goal"`
	TopContributions     []HypeTrainContributor `json:"top_contributions"`
	LastContribution     HypeTrainContributor   `json:"last_contribution"`
	StartedAt            time.Time              `json:"started_at"`
	ExpiresAt            time.Time              `json:"expires_at"`
}

type HypeTrainContributor struct {
	UserUID   string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
	Type      string `json:"type"`
	Total     int    `json:"total"`
}

//ChannelHypeTrainProgressEvent represents a channel.hype_train.progress event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelhype_trainprogress`
//The channel.hype_train.progress subscription type sends a notification when a hype train makes progress on the specified channel. When a hype train starts, one channel.hype_train.progress event will be sent for each contribution that caused the hype train to begin (in addition to the channel.hype_train.begin event). EventSub does not make strong assurances about the order of message delivery, so it is possible to receive channel.hype_train.progress before you receive the corresponding channel.hype_train.begin.
//After a hype train begins, any additional cheers or subscriptions on the channel will cause channel.hype_train.progress notifications to be sent. When the hype train is over, channel.hype_train.end is emitted.
//Must have channel:read:hype_train scope.
type ChannelHypeTrainProgressEvent struct {
	BroadcasterUID       string                 `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                 `json:"broadcaster_user_login"`
	BroadcasterUserName  string                 `json:"broadcaster_user_name"`
	Level                int                    `json:"level"`
	Total                int                    `json:"total"`
	Progress             int                    `json:"progress"`
	Goal                 int                    `json:"goal"`
	TopContributions     []HypeTrainContributor `json:"top_contributions"`
	LastContribution     HypeTrainContributor   `json:"last_contribution"`
	StartedAt            time.Time              `json:"started_at"`
	ExpiresAt            time.Time              `json:"expires_at"`
}

//ChannelHypeTrainEndEvent represents a channel.hype_train.end event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#channelhype_trainend`
//The channel.hype_train.end subscription type sends a notification when a hype train ends on the specified channel.
//Must have channel:read:hype_train scope.
type ChannelHypeTrainEndEvent struct {
	BroadcasterUID       string                 `json:"broadcaster_user_id"`
	BroadcasterUserLogin string                 `json:"broadcaster_user_login"`
	BroadcasterUserName  string                 `json:"broadcaster_user_name"`
	Level                int                    `json:"level"`
	Total                int                    `json:"total"`
	TopContributions     []HypeTrainContributor `json:"top_contributions"`
	StartedAt            time.Time              `json:"started_at"`
	ExpiresAt            time.Time              `json:"expires_at"`
	CooldownEndsAt       time.Time              `json:"cooldown_ends_at"`
}

//StreamOnlineEvent represents a stream.online event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#streamonline`
//The stream.online subscription type sends a notification when the specified broadcaster starts a stream.
//No authorization required.
type StreamOnlineEvent struct {
	ID                   string    `json:"id"`
	BroadcasterUID       string    `json:"broadcaster_user_id"`
	BroadcasterUserLogin string    `json:"broadcaster_user_login"`
	BroadcasterUserName  string    `json:"broadcaster_user_name"`
	Type                 string    `json:"type"`
	StartedAt            time.Time `json:"started_at"`
}

//StreamOfflineEvent represents a stream.offline event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#streamoffline`
//The stream.offline subscription type sends a notification when the specified broadcaster stops a stream.
//No authorization required.
type StreamOfflineEvent struct {
	BroadcasterUID       string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

//UserAuthorizationRevokeEvent represents a user.authorization.revoke event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#userauthorizationrevoke`
//The user.authorization.revoke subscription type sends a notification when a user’s authorization has been revoked for your client id. Use this webhook to meet government requirements for handling user data, such as GDPR, LGPD, or CCPA.
//Provided client_id must match the client id in the application access token.
type UserAuthorizationRevokeEvent struct {
	ClientID  string `json:"client_id"`
	UserUID   string `json:"user_id"`
	UserLogin string `json:"user_login,omitempty"`
	UserName  string `json:"user_name,omitempty"`
}

//UserUpdateEvent represents a user.update event `https://dev.twitch.tv/docs/eventsub/eventsub-subscription-types#userupdate`
//The user.update subscription type sends a notification when user updates their account.
//No authorization required. If you have the user:read:email scope, the notification will include email field.
type UserUpdateEvent struct {
	UserUID     string `json:"user_id"`
	UserLogin   string `json:"user_login"`
	UserName    string `json:"user_name"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description"`
}
