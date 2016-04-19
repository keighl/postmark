package postmark

import (
	"fmt"
	"net/url"
	"time"
)

///////////////////////////////////////
///////////////////////////////////////

// BounceType represents a type of bounce, and how many bounces have occured
// http://developer.postmarkapp.com/developer-api-bounce.html#bounce-types
type BounceType struct {
	// Type: bounce type identifier
	Type string
	// Name: full name of the bounce type
	Name string
	// Count: how many bounces have occured
	Count int64
}

// DeliveryStats represents bounce stats
type DeliveryStats struct {
	// InactiveMails: Number of inactive emails
	InactiveMails int64
	// Bounces: List of bounce types with total counts.
	Bounces []BounceType
}

// GetDeliveryStats returns delivery stats for the server
func (client *Client) GetDeliveryStats() (DeliveryStats, error) {
	res := DeliveryStats{}
	path := "deliverystats"
	err := client.doRequest("GET", path, nil, &res)
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
	// DumpAvailable: Specifies whether or not you can get a raw dump from this bounce. Postmark doesn’t store bounce dumps older than 30 days.
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

// GetBounces returns bounces for the server
// It returns a Bounce slice, the total bounce count, and any error that occured
// See options: http://developer.postmarkapp.com/developer-api-bounce.html#bounces
func (client *Client) GetBounces(count int64, offset int64, options map[string]interface{}) ([]Bounce, int64, error) {
	res := bouncesResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	path := fmt.Sprintf("bounces?%s", values.Encode())

	err := client.doRequest("GET", path, nil, &res)
	return res.Bounces, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// GetBounce fetches a bounce with bounceID
func (client *Client) GetBounce(bounceID int64) (Bounce, error) {
	res := Bounce{}
	path := fmt.Sprintf("bounces/%v", bounceID)
	err := client.doRequest("GET", path, nil, &res)
	return res, err
}
