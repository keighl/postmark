package postmark

import (
	"context"
	"fmt"
	"net/url"
)

// SenderSignature contains the details of the signature of the senders
type SenderSignature struct {
	Domain              string
	EmailAddress        string
	ReplyToEmailAddress string
	Name                string
	Confirmed           bool
	ID                  int64
}

///////////////////////////////////////
///////////////////////////////////////

// SenderSignaturesList is just a list of SenderSignatures as they are in the response
type SenderSignaturesList struct {
	TotalCount       int
	SenderSignatures []SenderSignature
}

// GetSenderSignatures calls GetSenderSignaturesWithContext with empty context
func (client *Client) GetSenderSignatures(count, offset int64) (SenderSignaturesList, error) {
	return client.GetSenderSignaturesWithContext(context.Background(), count, offset)
}

// GetSenderSignaturesWithContext gets a list of sender signatures, limited by count and paged by offset
func (client *Client) GetSenderSignaturesWithContext(ctx context.Context, count, offset int64) (SenderSignaturesList, error) {
	res := SenderSignaturesList{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("senders?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}
