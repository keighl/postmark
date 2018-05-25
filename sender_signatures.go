package postmark

import (
	"fmt"
	"net/url"
)

type SenderSignature struct {
	Domain              string
	EmailAddress        string
	ReplyToEmailAddress string
	Name                string
	Confirmed           bool
	ID                  int64
}

type SenderSignaturesList struct {
	TotalCount       int
	SenderSignatures []SenderSignature
}

func (client *Client) ListSenderSignatures(count, offset int64) (SenderSignaturesList, error) {
	res := SenderSignaturesList{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("senders?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}
