package postmark

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// OutboundMessage - a message sent from the Postmark server
type OutboundMessage struct {
	// TextBody - Text body of the message.
	TextBody string
	// HtmlBody - Html body of the message.
	HtmlBody string
	// Body - Raw source of the message.
	Body string
	// Tag - Tags associated with this message.
	Tag string
	// MessageID - Unique ID of the message.
	MessageID string
	// To - List of objects that contain To recipients.
	To []Recipient
	// Cc - List of objects that contain Cc recipients.
	Cc []Recipient
	// Bcc - List of objects that contain Bcc recipients.
	Bcc []Recipient
	// Recipients - List of recipients (just emails)
	Recipients []string
	// ReceivedAt - Timestamp
	ReceivedAt time.Time
	// From - The sender email address.
	From string
	// Subject - Email subject
	Subject string
	// Attachments - List of objects that each represent an attachment.
	Attachments []string
	// Status - Status of message in your Postmark activity.
	Status string
	// MessageEvents - List of summaries (MessageEvent) of things that have happened to this message. They can be Delivered, Opened, or Bounced as shown in the type field.
	MessageEvents []MessageEvent
}

// Recipient represents an individual who received a message
type Recipient struct {
	// Name is the recipient's name
	Name string
	// Emails is the recipient's email address
	Email string
}

// MessageEvent represents things that have happened to a message.
type MessageEvent struct {
	// Recipient is who received the message (just email address)
	Recipient string
	// ReceivedAt is the event timestamp
	ReceivedAt time.Time
	// Type of event (Delivered, Opened, or Bounced)
	Type string
	// Details contain information regarding the event
	// http://developer.postmarkapp.com/developer-api-messages.html#outbound-message-details
	Details map[string]string
}

///////////////////////////////////////
///////////////////////////////////////

// GetOutboundMessage calls GetOutboundMessageWithContext with empty context
func (client *Client) GetOutboundMessage(messageID string) (OutboundMessage, error) {
	return client.GetOutboundMessageWithContext(context.Background(), messageID)
}

// GetOutboundMessageWithContext fetches a specific outbound message via serverID
func (client *Client) GetOutboundMessageWithContext(ctx context.Context, messageID string) (OutboundMessage, error) {
	res := OutboundMessage{}
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/outbound/%s/details", messageID),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// GetOutboundMessageDump calls GetOutboundMessageDumpWithContext with empty context
func (client *Client) GetOutboundMessageDump(messageID string) (string, error) {
	return client.GetOutboundMessageDumpWithContext(context.Background(), messageID)
}

// GetOutboundMessageDumpWithContext fetches the raw source of message. If no dump is available this will return an empty string.
func (client *Client) GetOutboundMessageDumpWithContext(ctx context.Context, messageID string) (string, error) {
	res := dumpResponse{}
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/outbound/%s/dump", messageID),
		TokenType: server_token,
	}, &res)
	return res.Body, err
}

///////////////////////////////////////
///////////////////////////////////////

type outboundMessagesResponse struct {
	TotalCount int64
	Messages   []OutboundMessage
}

// GetOutboundMessages calls GetOutboundMessagesWithContext with empty context
func (client *Client) GetOutboundMessages(count int64, offset int64, options map[string]interface{}) ([]OutboundMessage, int64, error) {
	return client.GetOutboundMessagesWithContext(context.Background(), count, offset, options)
}

// GetOutboundMessagesWithContext fetches a list of outbound message on the server
// It returns a OutboundMessage slice, the total message count, and any error that occurred
// Note: that a single open is bound to a single recipient, so if the same message was sent to two recipients and both of them opened it, that will be represented by two entries in this array.
// Available options: http://developer.postmarkapp.com/developer-api-messages.html#outbound-message-search
func (client *Client) GetOutboundMessagesWithContext(ctx context.Context, count int64, offset int64, options map[string]interface{}) ([]OutboundMessage, int64, error) {
	res := outboundMessagesResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/outbound?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res.Messages, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// Open represents a single email open.
type Open struct {
	// FirstOpen - Indicates if the open was first open of message with MessageID and by Recipient. Any subsequent opens of the same message by the same Recipient will show false in this field. Postmark only saves first opens to its store, while all opens are available via Open web hooks.
	FirstOpen bool
	// UserAgent - Full user-agent header passed by the client software to Postmark. Postmark will fill in the Platform Client and OS fields based on this.
	UserAgent string
	// MessageID - Unique ID of the message.
	MessageID string
	// Client - Shows the email client (or browser) used to open the email. Name company and family are described in the parameters specification for this endpoint.
	Client map[string]string
	// OS - Shows the operating system used to open the email.
	OS map[string]string
	// Platform - Shows what platform was used to open the email. WebMail Desktop Mobile Unknown
	Platform string
	// ReadSeconds - Shows the reading time in seconds
	ReadSeconds int64
	// Geo - Contains IP of the recipient’s machine where the email was opened and the information based on that IP - geo coordinates (Coords) and country, region, city and zip.
	Geo map[string]string
}

type outboundMessageOpensResponse struct {
	TotalCount int64
	Opens      []Open
}

// GetOutboundMessagesOpens calls GetOutboundMessagesOpensWithContext with empty context
func (client *Client) GetOutboundMessagesOpens(count int64, offset int64, options map[string]interface{}) ([]Open, int64, error) {
	return client.GetOutboundMessagesOpensWithContext(context.Background(), count, offset, options)
}

// GetOutboundMessagesOpens fetches a list of opens on the server
// It returns a Open slice, the total opens count, and any error that occurred
// To get opens for a specific message, use GetOutboundMessageOpens()
// Available options: http://developer.postmarkapp.com/developer-api-messages.html#message-opens
func (client *Client) GetOutboundMessagesOpensWithContext(ctx context.Context, count int64, offset int64, options map[string]interface{}) ([]Open, int64, error) {
	res := outboundMessageOpensResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/outbound/opens?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res.Opens, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// GetOutboundMessageOpens calls GetOutboundMessageOpensWithContext with empty context
func (client *Client) GetOutboundMessageOpens(messageID string, count int64, offset int64) ([]Open, int64, error) {
	return client.GetOutboundMessageOpensWithContext(context.Background(), messageID, count, offset)
}

// GetOutboundMessageOpensWithContext fetches a list of opens for a specific message
// It returns a Open slice, the total opens count, and any error that occurred
func (client *Client) GetOutboundMessageOpensWithContext(ctx context.Context, messageID string, count int64, offset int64) ([]Open, int64, error) {
	res := outboundMessageOpensResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/outbound/opens/%s?%s", messageID, values.Encode()),
		TokenType: server_token,
	}, &res)
	return res.Opens, res.TotalCount, err
}
