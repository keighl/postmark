package postmark

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// InboundMessage - a message received from the Postmark server
type InboundMessage struct {
	// From - The sender email address.
	From string
	// FromName - The sender name.
	FromName string
	// FromFull - Sender email address and name.
	FromFull Recipient
	// To - Inbound address the message was sent to.
	To string
	// ToFull - Slice of all TO recipients
	ToFull []Recipient
	// CcFull - Slice of all CC recipients
	CcFull []Recipient
	// Cc - Cc recipient email address.
	Cc string
	// ReplyTo - Reply to override email address.
	ReplyTo string
	// OriginalRecipient - Receiver (RCPT TO) address this webhook is for.
	OriginalRecipient string
	// Subject - Email subject
	Subject string
	// Date - Timestamp
	Date string
	// MailboxHash - Custom hash that the email was sent to.
	MailboxHash string
	// TextBody - Plain text email message.
	TextBody string
	// HtmlBody - HTML email message.
	HtmlBody string
	// Tag - Tag name
	Tag string
	// Headers - List of objects that each represent a header name and value.
	Headers []Header
	// Attachments - List of objects that each represent an attachment.
	Attachments []Attachment
	// MessageID - Unique ID of the message.
	MessageID string
	// BlockedReason - Reason message was blocked.
	BlockedReason string
	// Status - Status of message in your Postmark activity.
	Status string
}

// Time returns a parsed time.Time struct
// Inbound messages return as RFC1123Z strangely
func (x InboundMessage) Time() (time.Time, error) {
	return time.Parse(time.RFC1123Z, x.Date)
}

///////////////////////////////////////
///////////////////////////////////////

// GetInboundMessage calls GetInboundMessageWithContext with empty context
func (client *Client) GetInboundMessage(messageID string) (InboundMessage, error) {
	return client.GetInboundMessageWithContext(context.Background(), messageID)
}

// GetInboundMessageWithContext fetches a specific inbound message via serverID
func (client *Client) GetInboundMessageWithContext(ctx context.Context, messageID string) (InboundMessage, error) {
	res := InboundMessage{}
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/inbound/%s/details", messageID),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

type inboundMessagesResponse struct {
	TotalCount int64
	Messages   []InboundMessage
}

// GetInboundMessages calls GetInboundMessagesWithContext with empty context
func (client *Client) GetInboundMessages(count int64, offset int64, options map[string]interface{}) ([]InboundMessage, int64, error) {
	return client.GetInboundMessagesWithContext(context.Background(), count, offset, options)
}

// GetInboundMessagesWithContext fetches a list of inbound message on the server
// It returns a InboundMessage slice, the total message count, and any error that occurred
// http://developer.postmarkapp.com/developer-api-messages.html#inbound-message-search
func (client *Client) GetInboundMessagesWithContext(ctx context.Context, count int64, offset int64, options map[string]interface{}) ([]InboundMessage, int64, error) {
	res := inboundMessagesResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/inbound?%s", values.Encode()),
		TokenType: server_token,
	}, &res)

	return res.Messages, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// BypassInboundMessage calls BypassInboundMessageWithContext with empty context
func (client *Client) BypassInboundMessage(messageID string) error {
	return client.BypassInboundMessageWithContext(context.Background(), messageID)
}

// BypassInboundMessageWithContext bypasses rules for a blocked inbound message
func (client *Client) BypassInboundMessageWithContext(ctx context.Context, messageID string) error {
	res := APIError{}
	err := client.doRequest(ctx, parameters{
		Method:    "PUT",
		Path:      fmt.Sprintf("messages/inbound/%s/bypass", messageID),
		TokenType: server_token,
	}, &res)

	if res.ErrorCode != 0 {
		return res
	}

	return err
}

///////////////////////////////////////
///////////////////////////////////////

// RetryInboundMessage calls RetryInboundMessageWithContext with empty context
func (client *Client) RetryInboundMessage(messageID string) error {
	return client.RetryInboundMessageWithContext(context.Background(), messageID)
}

// RetryInboundMessageWithContext retries a failed inbound message for processing
func (client *Client) RetryInboundMessageWithContext(ctx context.Context, messageID string) error {
	res := APIError{}
	err := client.doRequest(ctx, parameters{
		Method:    "PUT",
		Path:      fmt.Sprintf("messages/inbound/%s/retry", messageID),
		TokenType: server_token,
	}, &res)

	if res.ErrorCode != 0 {
		return res
	}

	return err
}
