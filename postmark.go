package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	postmarkURL = `https://api.postmarkapp.com`
)

type Client struct {
	// HTTPClient
	HTTPClient *http.Client

	// Server Token
	ServerToken string

	// AccountToken
	AccountToken string

	// BaseURL
	BaseURL string
}

func NewClient(serverToken string, accountToken string) *Client {
	return &Client{
		HTTPClient:   &http.Client{},
		ServerToken:  serverToken,
		AccountToken: accountToken,
		BaseURL:      postmarkURL,
	}
}

func (client *Client) doRequest(method string, path string, payload interface{}, dst interface{}) error {
	url := fmt.Sprintf("%s/%s", client.BaseURL, path)

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	if payload != nil {
		payloadData, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(payloadData))
	}

	req.Header.Add("X-Postmark-Server-Token", client.ServerToken)
	req.Header.Add("X-Postmark-Account-Token", client.AccountToken)

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, dst)
	return err
}

type APIError struct {
	ErrorCode int64
	Message   string
}

func (res APIError) Error() string {
	return res.Message
}
