// Package postmark ...
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

// Client provides a connection to the Postmark API
type Client struct {
	// HTTPClient is &http.Client{} by default
	HTTPClient *http.Client
	// Server Token: Used for requests that require server level privileges. This token can be found on the Credentials tab under your Postmark server.
	ServerToken string
	// AccountToken: Used for requests that require account level privileges. This token is only accessible by the account owner, and can be found on the Account tab of your Postmark account.
	AccountToken string
	// BaseURL is the root API endpoint
	BaseURL string
}

const (
	server_token  = "server"
	account_token = "account"
)

// Options is an object to hold variable parameters to perform request.
type parameters struct {
	// Method is HTTP method type.
	Method string
	// Path is postfix for URI.
	Path string
	// Payload for the request.
	Payload interface{}
	// TokenType defines which token to use
	TokenType string
}

// NewClient builds a new Client pointer using the provided tokens, a default HTTPClient, and a default API base URL
// Accepts `Server Token`, and `Account Token` as arguments
// http://developer.postmarkapp.com/developer-api-overview.html#authentication
func NewClient(serverToken string, accountToken string) *Client {
	return &Client{
		HTTPClient:   &http.Client{},
		ServerToken:  serverToken,
		AccountToken: accountToken,
		BaseURL:      postmarkURL,
	}
}

func (client *Client) doRequest(opts parameters, dst interface{}) error {
	url := fmt.Sprintf("%s/%s", client.BaseURL, opts.Path)

	req, err := http.NewRequest(opts.Method, url, nil)
	if err != nil {
		return err
	}

	if opts.Payload != nil {
		payloadData, err := json.Marshal(opts.Payload)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(payloadData))
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	switch opts.TokenType {
	case account_token:
		req.Header.Add("X-Postmark-Account-Token", client.AccountToken)

	default:
		req.Header.Add("X-Postmark-Server-Token", client.ServerToken)
	}

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

// APIError represents errors returned by Postmark
type APIError struct {
	// ErrorCode: see error codes here (http://developer.postmarkapp.com/developer-api-overview.html#error-codes)
	ErrorCode int64
	// Message contains error details
	Message string
}

// Error returns the error message details
func (res APIError) Error() string {
	return res.Message
}
