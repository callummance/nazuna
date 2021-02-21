package restclient

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/endpoints"
)

const (
	ScopeReadAnalyticsExtensions  = "analytics:read:extensions"
	ScopeReadAnalyticsGames       = "analytics:read:games"
	ScopeReadBits                 = "bits:read"
	ScopeEditChannelCommercials   = "channel:edit:commercial"
	ScopeManageChannelBroadcast   = "channel:manage:broadcast"
	ScopeManageChannelExtensions  = "channel:manage:extensions"
	ScopeManageChannelRedemptions = "channel:manage:redemptions"
	ScopeManageChannelVideos      = "channel:manage:videos"
	ScopeReadChannelEditors       = "channel:read:editors"
	ScopeReadChannelHypeTrain     = "channel:read:hype_train"
	ScopeReadChannelRedemptions   = "channel:read:redemptions"
	ScopeReadChannelStreamKey     = "channel:read:stream_key"
	ScopeReadChannelSubscriptios  = "channel:read:subscriptions"
	ScopeEditClips                = "clips:edit"
	ScopeModerationRead           = "moderation:read"
	ScopeUserEdit                 = "user:edit"
	ScopeUserEditFollows          = "user:edit:follows"
	ScopeUserReadBlocked          = "user:read:blocked_users"
	ScopeUserManageBlocked        = "user:manage:blocked_users"
	ScopeUserReadBroadcast        = "user:read:broadcast"
	ScopeUserReadEmail            = "user:read:email"
)

func getClientCredentials(clientID, clientSecret string, scopes []string) *http.Client {
	ctx := context.Background()
	conf := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     endpoints.Twitch.TokenURL,
		Scopes:       scopes,
	}

	tok, _ := conf.Token(ctx)
	logrus.Debugf("Using token %#v", tok)

	client := conf.Client(ctx)
	return client
}
