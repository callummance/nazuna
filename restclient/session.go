package restclient

import (
	"net/http"
)

const apiBaseURL = "https://api.twitch.tv/helix"

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
