package postmark

import (
	"fmt"
	"time"
)

// Email is exactly what it sounds like
type Email struct {
	// From: REQUIRED The sender email address. Must have a registered and confirmed Sender Signature.
	From string `json:",omitempty"`
	// To: REQUIRED Recipient email address. Multiple addresses are comma separated. Max 50.
	To string `json:",omitempty"`
	// Cc recipient email address. Multiple addresses are comma separated. Max 50.
	Cc string `json:",omitempty"`
	// Bcc recipient email address. Multiple addresses are comma separated. Max 50.
	Bcc string `json:",omitempty"`
	// Subject: Email subject
	Subject string `json:",omitempty"`
	// Tag: Email tag that allows you to categorize outgoing emails and get detailed statistics.
	Tag string `json:",omitempty"`
	// HtmlBody: HTML email message. REQUIRED, If no TextBody specified
	HtmlBody string `json:",omitempty"`
	// TextBody: Plain text email message. REQUIRED, If no HtmlBody specified
	TextBody string `json:",omitempty"`
	// ReplyTo: Reply To override email address. Defaults to the Reply To set in the sender signature.
	ReplyTo string `json:",omitempty"`
	// Headers: List of custom headers to include.
	Headers []Header `json:",omitempty"`
	// TrackOpens: Activate open tracking for this email.
	TrackOpens bool `json:",omitempty"`
	// Attachments: List of attachments
	Attachments []Attachment `json:",omitempty"`
}

// Header - an email header
type Header struct {
	// Name: header name
	Name string
	// Value: header value
	Value string
}

// Attachment is an optional encoded file to send along with an email
type Attachment struct {
	// Name: attachment name
	Name string
	// Content: Base64 encoded attachment data
	Content string
	// ContentType: attachment MIME type
	ContentType string
}

// EmailResponse holds info in response to a send/send-batch request
// Even if API request comes back successful, check the ErrorCode to see if there might be a delivery problem
type EmailResponse struct {
	// To: Recipient email address
	To string
	// SubmittedAt: Timestamp
	SubmittedAt time.Time
	// MessageID: ID of message
	MessageID string
	// ErrorCode: API Error Codes
	ErrorCode int64
	// Message: Response message
	Message string
}

// SendEmail sends, well, an email.
func (client *Client) SendEmail(email Email) (EmailResponse, error) {
	res := EmailResponse{}
	err := client.doRequest(parameters{
		Method:    "POST",
		Path:      "email",
		Payload:   email,
		TokenType: server_token,
	}, &res)

	if res.ErrorCode != 0 {
		return res, fmt.Errorf(`%v %s`, res.ErrorCode, res.Message)
	}

	return res, err
}

// SendEmailBatch sends multiple emails together
// Note, individual emails in the batch can error, so it would be wise to
// range over the responses and sniff for errors
func (client *Client) SendEmailBatch(emails []Email) ([]EmailResponse, error) {
	res := []EmailResponse{}
	err := client.doRequest(parameters{
		Method:    "POST",
		Path:      "email/batch",
		Payload:   emails,
		TokenType: server_token,
	}, &res)
	return res, err
}
