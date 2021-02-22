package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/callummance/nazuna/messages"
	"github.com/sirupsen/logrus"
)

func (c *Client) CreateSubscription(condition interface{}, transport messages.TransportOpts) (*messages.SubscriptionRequestStatus, error) {
	var reqBody messages.Subscription
	switch t := condition.(type) {
	case messages.ConditionChannelUpdate:
		reqBody.Type = messages.SubscriptionChannelUpdate
	case messages.ConditionChannelFollow:
		reqBody.Type = messages.SubscriptionChannelFollow
	case messages.ConditionChannelSubscribe:
		reqBody.Type = messages.SubscriptionChannelSubscribe
	case messages.ConditionChannelCheer:
		reqBody.Type = messages.SubscriptionChannelCheer
	case messages.ConditionChannelBan:
		reqBody.Type = messages.SubscriptionChannelBan
	case messages.ConditionChannelUnban:
		reqBody.Type = messages.SubscriptionChannelUnban
	case messages.ConditionChannelPointsCustomRewardAdd:
		reqBody.Type = messages.SubscriptionChannelPointsCustomRewardAdd
	case messages.ConditionChannelPointsCustomRewardUpdate:
		reqBody.Type = messages.SubscriptionChannelPointsCustomRewardUpdate
	case messages.ConditionChannelPointsCustomRewardRemove:
		reqBody.Type = messages.SubscriptionChannelPointsCustomRewardRemove
	case messages.ConditionChannelPointsCustomRewardRedemptionAdd:
		reqBody.Type = messages.SubscriptionChannelPointsCustomRewardRedemptionAdd
	case messages.ConditionChannelPointsCustomRewardRedemptionUpdate:
		reqBody.Type = messages.SubscriptionChannelPointsCustomRewardRedemptionUpdate
	case messages.ConditionChannelHypeTrainBegin:
		reqBody.Type = messages.SubscriptionChannelHypeTrainBegin
	case messages.ConditionChannelHypeTrainProgress:
		reqBody.Type = messages.SubscriptionChannelHypeTrainProgress
	case messages.ConditionChannelHypeTrainEnd:
		reqBody.Type = messages.SubscriptionChannelHypeTrainEnd
	case messages.ConditionStreamOnline:
		reqBody.Type = messages.SubscriptionStreamOnline
	case messages.ConditionStreamOffline:
		reqBody.Type = messages.SubscriptionStreamOffline
	case messages.ConditionUserAuthorizationRevoke:
		reqBody.Type = messages.SubscriptionUserAuthorizationRevoke
	case messages.ConditionUserUpdate:
		reqBody.Type = messages.SubscriptionUserUpdate
	default:
		logrus.Warnf("CreateSubscription call was provided with a condition of unrecognized type")
		return nil, fmt.Errorf("type %v supplied to CreateSubscription is not a valid condition", t)
	}
	reqBody.Version = "1"
	reqBody.Condition = condition
	reqBody.Transport = transport

	//Marshal request body
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		logrus.Warnf("Failed to marshal CreateSubscription request body due to error %v", err)
		return nil, err
	}
	logrus.Tracef("Submitting CreateSubscription request with body %s", bodyBytes)
	//Send POST request
	req, err := http.NewRequest("POST", subscriptionEndpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		logrus.Warnf("Failed to make CreateSubscription request due to error %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-ID", c.clientID)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		logrus.Warnf("Failed to make CreateSubscription request due to error %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	//Decode response
	var result messages.SubscriptionRequestStatus
	if resp.StatusCode == http.StatusConflict {
		//Quietly ignore duplicate subscriptions
		return nil, nil
	} else if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		dump, _ := httputil.DumpResponse(resp, true)
		logrus.Infof("Got non-OK response %s to subscription creation request", dump)
		return nil, fmt.Errorf("got non-OK response %s to subscription creation request", dump)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logrus.Warnf("Failed to decode response to CreateSubscription request due to error %v", err)
		logrus.Tracef("response: %v", resp)
		return nil, err
	}
	return &result, nil
}

type SubscriptionsParams struct {
	Status string `json:"status,omitempty"`
	Type   string `json:"type,omitempty"`
}

func (p SubscriptionsParams) insertToValues(initialValues *url.Values) {
	if p.Status != "" {
		(*initialValues)["status"] = []string{p.Status}
	}

	if p.Type != "" {
		(*initialValues)["type"] = []string{p.Type}
	}
}

func (c *Client) getSubscriptionsPage(params *SubscriptionsParams, pagination *pagination) (*subscriptionsPage, error) {
	logrus.Debugf("Requesting page of subscriptions with filters %#v from api.", params)
	//Build query URL
	var query url.Values
	url, err := url.Parse(subscriptionEndpoint)
	if err != nil {
		logrus.Errorf("Failed to parse subscription endpoint with error %v", err)
		return nil, err
	}
	if params != nil {
		params.insertToValues(&query)
	}
	if pagination != nil {
		pagination.insertToValues(&query)
	}
	url.RawQuery = query.Encode()

	//Send GET request
	req, err := http.NewRequest("GET", url.String(), http.NoBody)
	if err != nil {
		logrus.Warnf("Failed to make Subscriptions request due to error %v", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-ID", c.clientID)
	logrus.Tracef("Sending request %#v to retrieve subscriptions", req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logrus.Warnf("Failed to make Subscriptions request due to error %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	//Decode response
	var result subscriptionsPage
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		dump, _ := httputil.DumpResponse(resp, true)
		logrus.Infof("Got non-OK response %s to subscription list request", dump)
		return nil, fmt.Errorf("got non-OK response %s to subscription list request", dump)
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		logrus.Warnf("Failed to decode response to subscriptions list request due to error %v", err)
		logrus.Tracef("response: %v", resp)
		return nil, err
	}
	return &result, nil
}

type SubscriptionResult struct {
	Subscription *messages.Subscription
	Err          error
}

//Subscriptions returns a channel which will be populated with all Eventsub subscriptions owned by the current app
func (c *Client) Subscriptions(filters *SubscriptionsParams) chan SubscriptionResult {
	ch := make(chan SubscriptionResult)
	go func(c *Client) {
		var currentPage subscriptionsPage
		initialPageFetched := false
		lastLoc := 0
		for {
			if lastLoc+1 < len(currentPage.Data) {
				//We still have subscriptions from the previously fetched page
				lastLoc++
				ch <- SubscriptionResult{
					Subscription: &currentPage.Data[lastLoc],
					Err:          nil,
				}
			} else if currentPage.Pagination.Cursor == "" && initialPageFetched {
				//Run out of subscriptions fetched and no pagination data, so we must be done
				close(ch)
				return
			} else {
				//Need to fetch more members
				pagination := pagination{
					After: currentPage.Pagination.Cursor,
				}
				nextPage, err := c.getSubscriptionsPage(filters, &pagination)
				if err != nil {
					logrus.Warnf("Failed to fetch page of subscriptions from API due to error %v", err)
					ch <- SubscriptionResult{
						Subscription: nil,
						Err:          err,
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
				ch <- SubscriptionResult{
					Subscription: &currentPage.Data[0],
					Err:          nil,
				}
			}
		}
	}(c)
	return ch
}

//DeleteSubscription attempts to delete a previously-created EventSub subscription from the twitch API
func (c *Client) DeleteSubscription(subscriptionID string) error {
	logrus.Debugf("Requested deletion of subscription with ID %v.", subscriptionID)
	//Build query URL
	url, err := url.Parse(subscriptionEndpoint)
	query := url.Query()
	if err != nil {
		logrus.Errorf("Failed to parse subscription endpoint with error %v", err)
		return err
	}
	query.Set("id", subscriptionID)
	url.RawQuery = query.Encode()

	//Send DELETE request
	req, err := http.NewRequest("DELETE", url.String(), http.NoBody)
	if err != nil {
		logrus.Warnf("Failed to make delete subscription request due to error %v", err)
		return err
	}
	req.Header.Add("Client-ID", c.clientID)
	logrus.Tracef("Sending request %#v to delete subscriptions", req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		logrus.Warnf("Failed to make delete subscription request due to error %v", err)
		return err
	}
	defer resp.Body.Close()

	//Decode response
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		dump, _ := httputil.DumpResponse(resp, true)
		logrus.Infof("Got non-OK response %s to subscription list request", dump)
		return fmt.Errorf("got non-OK response %s to subscription list request", dump)
	}
	return nil
}
