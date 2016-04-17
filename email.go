package postmark

import (
	"time"
)

type Email struct {
	// From: REQUIRED The sender email address. Must have a registered and confirmed Sender Signature.
	From string `json:omitempty`
	// To: REQUIRED Recipient email address. Multiple addresses are comma seperated. Max 50.
	To string `json:omitempty`
	// Cc recipient email address. Multiple addresses are comma seperated. Max 50.
	Cc string `json:omitempty`
	// Bcc recipient email address. Multiple addresses are comma seperated. Max 50.
	Bcc string `json:omitempty`
	// Subject: Email subject
	Subject string `json:omitempty`
	// Tag: Email tag that allows you to categorize outgoing emails and get detailed statistics.
	Tag string `json:omitempty`
	// HtmlBody: HTML email message. REQUIRED, If no TextBody specified
	HtmlBody string `json:omitempty`
	// TextBody: Plain text email message. REQUIRED, If no HtmlBody specified
	TextBody string `json:omitempty`
	// ReplyTo: Reply To override email address. Defaults to the Reply To set in the sender signature.
	ReplyTo string `json:omitempty`
	// Headers: List of custom headers to include.
	Headers []Header `json:omitempty`
	// TrackOpens: Activate open tracking for this email.
	TrackOpens bool `json:omitempty`
	// Attachments: List of attachments
	Attachments []Attachment `json:omitempty`
}

type Header struct {
	// Name: custom header name
	Name string
	// Value: custom header value
	Value string
}

type Attachment struct {
	// Name: attachment name
	Name string
	// Content: Base64 encoded attachment data
	Content string
	// ContentType: attachment MIME type
	ContentType string
}

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

func (client *Client) SendEmail(email Email) (EmailResponse, error) {
	res := EmailResponse{}
	err := client.doRequest("POST", "email", email, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
