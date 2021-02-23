package restclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

const usersEndpoint = apiBaseURL + "/users"

type TwitchUser struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageURL string    `json:"profile_image_url"`
	OfflineImageURL string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
}

type twitchGetUserResult struct {
	Data []TwitchUser `json:"data"`
}

//GetUsers returns user data for up to 100 user IDs or names (https://dev.twitch.tv/docs/api/reference#get-users)
func (c *Client) GetUsers(ids []string, logins []string) ([]TwitchUser, error) {
	logrus.Debugf("Requesting users with ids %v and logins %v", ids, logins)
	if len(ids)+len(logins) > 100 {
		return nil, fmt.Errorf("only a maximum of 100 users can be requested at a time")
	}
	//Build query URL
	url, err := url.Parse(usersEndpoint)
	if err != nil {
		logrus.Errorf("Failed to parse subscription endpoint with error %v", err)
		return nil, err
	}
	query := url.Query()
	for _, id := range ids {
		query.Set("id", id)
	}
	for _, login := range logins {
		query.Set("login", login)
	}
	url.RawQuery = query.Encode()

	//Send GET request
	req, err := http.NewRequest("GET", url.String(), http.NoBody)
	if err != nil {
		logrus.Warnf("Failed to make GetUsers request due to error %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-ID", c.clientID)
	logrus.Tracef("Sending request %#v to retrieve users", req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logrus.Warnf("Failed to make GetUsers request due to error %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	//Decode response
	var result twitchGetUserResult
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		dump, _ := httputil.DumpResponse(resp, true)
		logrus.Infof("Got non-OK response %q to subscription list request", dump)
		return nil, fmt.Errorf("got non-OK response %q to subscription list request", dump)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logrus.Warnf("Failed to decode response to subscriptions list request due to error %v", err)
		logrus.Tracef("response: %v", resp)
		return nil, err
	}
	logrus.Trace(result)
	return result.Data, nil
}
