package restclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
)

const streamsEndpoint = apiBaseURL + "/streams"

type TwitchStream struct {
	StreamID     string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserLogin    string    `json:"user_login"`
	UserName     string    `json:"user_name"`
	GameID       string    `json:"game_id"`
	GameName     string    `json:"game_name"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	ViewerCount  int       `json:"viewer_count"`
	StartedAt    time.Time `json:"start_at"`
	Language     string    `json:"language"`
	ThumbnailURL string    `json:"thumbnail_url"`
	TagIDs       []string  `json:"tag_ids"`
}

type GetStreamsOpts struct {
	After     string   `url:"after,omitempty"`
	Before    string   `url:"before,omitempty"`
	First     int      `url:"first,omitempty"`
	GameID    []string `url:"game_id,omitempty"`
	Language  []string `url:"language,omitempty"`
	UserID    []string `url:"user_id,omitempty"`
	UserLogin []string `url:"user_login,omitempty"`
}

func (c *Client) GetStreamsPage(opts GetStreamsOpts, pagination *pagination) (*streamsPage, error) {
	logrus.Debugf("Requesting page of streams with filters %#v from api.", opts)
	//Build query URL
	url, err := url.Parse(streamsEndpoint)
	if err != nil {
		logrus.Errorf("Failed to parse streams endpoint with error %v", err)
		return nil, err
	}
	vals, err := query.Values(opts)
	if err != nil {
		logrus.Errorf("Failed to encode stream options into querystring with error %v", err)
		return nil, err
	}
	if pagination != nil {
		pagination.insertToValues(&vals)
	}
	url.RawQuery = vals.Encode()

	//Send GET request
	req, err := http.NewRequest("GET", url.String(), http.NoBody)
	if err != nil {
		logrus.Warnf("Failed to make Streams request due to error %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-ID", c.clientID)
	logrus.Tracef("Sending request %#v to url %v retrieve streams", req, req.URL.String())
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logrus.Warnf("Failed to make Streams request due to error %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	//Decode response
	var result streamsPage
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		dump, _ := httputil.DumpResponse(resp, true)
		logrus.Infof("Got non-OK response %s to streams list request", dump)
		return nil, fmt.Errorf("got non-OK response %s to streams list request", dump)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logrus.Warnf("Failed to decode response to streams list request due to error %v", err)
		logrus.Tracef("response: %v", resp)
		return nil, err
	}
	return &result, nil
}

type StreamResult struct {
	Stream *TwitchStream
	Err    error
}

//GetStreamsIter returns a channel which will be populated with data on streams which match the provided options
func (c *Client) GetStreamsIter(opts GetStreamsOpts) chan StreamResult {
	ch := make(chan StreamResult)
	go func(c *Client) {
		var currentPage streamsPage
		initialPageFetched := false
		lastLoc := 0
		for {
			if lastLoc+1 < len(currentPage.Data) {
				//We still have streams from the previously fetched page
				lastLoc++
				ch <- StreamResult{
					Stream: &currentPage.Data[lastLoc],
					Err:    nil,
				}
			} else if currentPage.Pagination.Cursor == "" && initialPageFetched {
				//Run out of streams fetched and no pagination data, so we must be done
				close(ch)
				return
			} else {
				//Need to fetch more streams
				pagination := pagination{
					After: currentPage.Pagination.Cursor,
				}
				nextPage, err := c.GetStreamsPage(opts, &pagination)
				if err != nil {
					logrus.Warnf("Failed to fetch page of streams from API due to error %v", err)
					ch <- StreamResult{
						Stream: nil,
						Err:    err,
					}
				}
				//If new page is empty, we must also be done
				if len(nextPage.Data) == 0 {
					close(ch)
					return
				}
				initialPageFetched = true
				currentPage = *nextPage
				lastLoc = 0
				ch <- StreamResult{
					Stream: &currentPage.Data[0],
					Err:    nil,
				}
			}
		}
	}(c)
	return ch
}
