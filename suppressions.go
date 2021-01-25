package postmark

import (
	"fmt"
	"net/url"
)

// SuppressionReasonType - The reason type of suppression
type SuppressionReasonType string

// OriginType - The reason type of origin
type OriginType string

const (
	HardBounceReason        SuppressionReasonType = "HardBounce"
	SpamComplaintReason     SuppressionReasonType = "SpamComplaint"
	ManualSuppressionReason SuppressionReasonType = "ManualSuppression"
	RecipientOrigin         OriginType            = "Recipient"
	CustomerOrigin          OriginType            = "Customer"
	AdminOrigin             OriginType            = "Admin"
)

// SuppressionEmail - Suppression email entry
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

// suppressionsResponse - A message received from the Postmark server
type suppressionsResponse struct {
	// Suppressions - The slice of suppression email address.
	Suppressions []SuppressionEmail
}

// EmailAddress - Email address to delete
type EmailAddress struct {
	EmailAddress string `json:",omitempty"`
}

// DeleteSuppressions - A payload of email address to delete
type DeleteSuppressions struct {
	// Suppressions - The slice of email address to delete.
	Suppressions []EmailAddress `json:",omitempty"`
}

// deleteSuppressionsResponse - A message received from the Postmark server
type deleteSuppressionsResponse struct {
	// Suppressions - The slice of deleted email status and reason.
	Suppressions []SuppressionDelete
}

// SuppressionDelete - Suppression email deleted
type SuppressionDelete struct {
	EmailAddress string
	Status       string
	Message      *string
}

// GetSuppressionEmails fetches a email addresses in the list of suppression dump on the server
// It returns a SuppressionEmails slice, and any error that occurred
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

	res := suppressionsResponse{}
	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      path,
		TokenType: server_token,
	}, &res)
	return res.Suppressions, err
}

// DeleteSuppressionEmails delete email addresses in the list of suppression dump on the server
// Noet: SpamComplaint and ManualSuppression with Origin: Customer cannot be deleted
// It return a SuppressionDelete slice, and any error that occurred
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
