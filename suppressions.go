package postmark

import (
	"fmt"
	"net/url"
)

type SuppressionReasonType string
type OriginType string

const (
	HardBounce        SuppressionReasonType = "HardBounce"
	SpamComplaint     SuppressionReasonType = "SpamComplaint"
	ManualSuppression SuppressionReasonType = "ManualSuppression"
	Re                OriginType            = "Recipient"
	Customer          OriginType            = "Customer"
	Admin             OriginType            = "Admin"
)

// SuppressionEmail - a message received from the Postmark server
type SuppressionEmail struct {
	// EmailAddress - The email address of suppression.
	EmailAddress string
	// SuppressionReason - The reason of suppression.
	SuppressionReason SuppressionReasonType
	// Origin - Possible options: Recipient, Customer, Admin.
	Origin OriginType
	// CreatedAt - Timestamp.
	CreatedAt string
}

// suppressionsResp - a message received from the Postmark server
type suppressionsResp struct {
	// Suppressions - The email address of suppression.
	Suppressions []SuppressionEmail
}

// EmailAddress - a message received from the Postmark server
type EmailAddress struct {
	// EmailAddress - The email address of suppression.
	EmailAddress string `json:",omitempty"`
}

// DeleteSuppressions - a message received from the Postmark server
type DeleteSuppressions struct {
	// Suppressions - The slice of email address to delete.
	Suppressions []EmailAddress `json:",omitempty"`
}

// deleteSuppressionsResponse - a message received from the Postmark server
type deleteSuppressionsResponse struct {
	// Suppressions - The slice of email address status and message.
	Suppressions []SuppressionDelete
}

// SuppressionDelete - a message received from the Postmark server
type SuppressionDelete struct {
	EmailAddress string
	Status       string
	Message      *string
}

// GetSuppressionEmails fetches a email addresses in the list of suppression dump on the server
// It returns a SuppressionEmails slice, the total message count, and any error that occurred
// https://postmarkapp.com/developer/api/suppressions-api#suppression-dump
func (client *Client) GetSuppressionEmails(streamID string, options map[string]interface{}) ([]SuppressionEmail, error) {
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	path := fmt.Sprintf("message-streams/%s/suppressions/dump", streamID)
	if len(options) != 0 {
		path = fmt.Sprintf("%s?%s", path, values.Encode())
	}

	res := suppressionsResp{}
	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res.Suppressions, err
}

// DeleteSuppressionEmails delete email addresses in the list of suppression dump on the server
// Noet: SpamComplaint and ManualSuppression with Origin: Customer cannot be deleted
// It return email address failed delete and failing reason
// https://postmarkapp.com/developer/api/suppressions-api#delete-a-suppression
func (client *Client) DeleteSuppressionEmails(streamID string, delete []string) ([]SuppressionDelete, error) {
	emailAddresses := make([]EmailAddress, 0, 8)
	for _, address := range delete {
		emailAddresses = append(emailAddresses, EmailAddress{
			EmailAddress: address,
		})
	}

	res := deleteSuppressionsResponse{}
	err := client.doRequest(parameters{
		Method: "POST",
		Path:   fmt.Sprintf("message-streams/%s/suppressions/delete", streamID),
		Payload: DeleteSuppressions{
			Suppressions: emailAddresses,
		},
		TokenType: server_token,
	}, &res)
	return res.Suppressions, err
}
