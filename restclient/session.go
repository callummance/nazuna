package restclient

import (
	"net/http"
)

const subscriptionEndpoint = "https://api.twitch.tv/helix/eventsub/subscriptions"

type Client struct {
	httpClient *http.Client
	clientID   string
}

func InitClient(clientID, clientSecret string, scopes []string) *Client {
	httpClient := getClientCredentials(clientID, clientSecret, scopes)
	return &Client{
		httpClient: httpClient,
		clientID:   clientID,
	}
}
