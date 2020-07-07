package postmark

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

///////////////////////////////////////
///////////////////////////////////////

// BounceType represents a type of bounce, and how many bounces have occurred
// http://developer.postmarkapp.com/developer-api-bounce.html#bounce-types
type BounceType struct {
	// Type: bounce type identifier
	Type string
	// Name: full name of the bounce type
	Name string
	// Count: how many bounces have occurred
	Count int64
}

// DeliveryStats represents bounce stats
type DeliveryStats struct {
	// InactiveMails: Number of inactive emails
	InactiveMails int64
	// Bounces: List of bounce types with total counts.
	Bounces []BounceType
}

// GetDeliveryStats calls GetDeliveryStatsWithContext with empty context
func (client *Client) GetDeliveryStats() (DeliveryStats, error) {
	return client.GetDeliveryStatsWithContext(context.Background())
}

// GetDeliveryStatsWithContext returns delivery stats for the server
func (client *Client) GetDeliveryStatsWithContext(ctx context.Context) (DeliveryStats, error) {
	res := DeliveryStats{}
	path := "deliverystats"
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// Bounce represents a specific delivery failure
type Bounce struct {
	// ID: ID of bounce
	ID int64
	// Type: Bounce type
	Type string
	// TypeCode: Bounce type code
	TypeCode int64
	// Name: Bounce type name
	Name string
	// Tag: Tag name
	Tag string
	// MessageID: ID of message
	MessageID string
	// Description: Description of bounce
	Description string
	// Details: Details on the bounce
	Details string
	// Email: Email address that bounced
	Email string
	// BouncedAt: Timestamp of bounce
	BouncedAt time.Time
	// DumpAvailable: Specifies whether or not you can get a raw dump from this bounce. Postmark does not store bounce dumps older than 30 days.
	DumpAvailable bool
	// Inactive: Specifies if the bounce caused Postmark to deactivate this email.
	Inactive bool
	// CanActivate: Specifies whether or not you are able to reactivate this email.
	CanActivate bool
	// Subject: Email subject
	Subject string
}

type bouncesResponse struct {
	TotalCount int64
	Bounces    []Bounce
}

// GetBounces calls GetBouncesWithContext with empty context
func (client *Client) GetBounces(count int64, offset int64, options map[string]interface{}) ([]Bounce, int64, error) {
	return client.GetBouncesWithContext(context.Background(), count, offset, options)
}

// GetBouncesWithContext returns bounces for the server
// It returns a Bounce slice, the total bounce count, and any error that occurred
// Available options: http://developer.postmarkapp.com/developer-api-bounce.html#bounces
func (client *Client) GetBouncesWithContext(ctx context.Context, count int64, offset int64, options map[string]interface{}) ([]Bounce, int64, error) {
	res := bouncesResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	path := fmt.Sprintf("bounces?%s", values.Encode())

	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res.Bounces, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// GetBounce calls GetBounceWithContext with empty context
func (client *Client) GetBounce(bounceID int64) (Bounce, error) {
	return client.GetBounceWithContext(context.Background(), bounceID)
}

// GetBounceWithContext fetches a single bounce with bounceID
func (client *Client) GetBounceWithContext(ctx context.Context, bounceID int64) (Bounce, error) {
	res := Bounce{}
	path := fmt.Sprintf("bounces/%v", bounceID)
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

type dumpResponse struct {
	Body string
}

// GetBounceDump calls GetBounceDumpWithContext with empty context
func (client *Client) GetBounceDump(bounceID int64) (string, error) {
	return client.GetBounceDumpWithContext(context.Background(), bounceID)
}

// GetBounceDumpWithContext fetches a SMTP data dump for a single bounce
func (client *Client) GetBounceDumpWithContext(ctx context.Context, bounceID int64) (string, error) {
	res := dumpResponse{}
	path := fmt.Sprintf("bounces/%v/dump", bounceID)
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res.Body, err
}

///////////////////////////////////////
///////////////////////////////////////

type activateBounceResponse struct {
	Message string
	Bounce  Bounce
}

// ActivateBounce calls ActivateBounceWithContext with empty context
func (client *Client) ActivateBounce(bounceID int64) (Bounce, string, error) {
	return client.ActivateBounceWithContext(context.Background(), bounceID)
}

// ActivateBounceWithContext reactivates a bounce for resending. Returns the bounce, a
// message, and any error that occurs
// TODO: clarify this with Postmark
func (client *Client) ActivateBounceWithContext(ctx context.Context, bounceID int64) (Bounce, string, error) {
	res := activateBounceResponse{}
	path := fmt.Sprintf("bounces/%v/activate", bounceID)
	err := client.doRequest(ctx, parameters{
		Method:    "PUT",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res.Bounce, res.Message, err
}

///////////////////////////////////////
///////////////////////////////////////

type bouncedTagsResponse struct {
	Tags []string `json:"tags"`
}

// GetBouncedTags calls GetBouncedTagsWithContext with empty context
func (client *Client) GetBouncedTags() ([]string, error) {
	return client.GetBouncedTagsWithContext(context.Background())
}

// GetBouncedTagsWithContext retrieves a list of tags that have generated bounced emails
func (client *Client) GetBouncedTagsWithContext(ctx context.Context) ([]string, error) {
	var raw json.RawMessage
	path := "bounces/tags"
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &raw)

	if err != nil {
		return []string{}, err
	}

	// PM returns this payload in an impossible to unmarshal way
	// ["tag1","tag2","tag3"]. So let's rejigger it to make it possible.
	jsonString := fmt.Sprintf(`{"tags": %s}`, string(raw))
	res := bouncedTagsResponse{}
	err = json.Unmarshal([]byte(jsonString), &res)

	return res.Tags, err
}
