package postmark

import (
	"context"
	"fmt"
)

// Server represents a server registered in your Postmark account
type Server struct {
	// ID of server
	ID int64
	// Name of server
	Name string
	// ApiTokens associated with server.
	ApiTokens []string
	// ServerLink to your server overview page in Postmark.
	ServerLink string
	// Color of the server in the rack screen. Purple Blue Turquoise Green Red Yellow Grey
	Color string
	// SmtpApiActivated specifies whether or not SMTP is enabled on this server.
	SmtpApiActivated bool
	// RawEmailEnabled allows raw email to be sent with inbound.
	RawEmailEnabled bool
	// InboundAddress is the inbound email address
	InboundAddress string
	// InboundHookUrl to POST to every time an inbound event occurs.
	InboundHookUrl string
	// BounceHookUrl to POST to every time a bounce event occurs.
	BounceHookUrl string
	// OpenHookUrl to POST to every time an open event occurs.
	OpenHookUrl string
	// PostFirstOpenOnly - If set to true, only the first open by a particular recipient will initiate the open webhook. Any
	// subsequent opens of the same email by the same recipient will not initiate the webhook.
	PostFirstOpenOnly bool
	// TrackOpens indicates if all emails being sent through this server have open tracking enabled.
	TrackOpens bool
	// InboundDomain is the inbound domain for MX setup
	InboundDomain string
	// InboundHash is the inbound hash of your inbound email address.
	InboundHash string
	// InboundSpamThreshold is the maximum spam score for an inbound message before it's blocked.
	InboundSpamThreshold int64
}

///////////////////////////////////////
///////////////////////////////////////

// GetServer calls GetServerWithContext with empty context
func (client *Client) GetServer(serverID string) (Server, error) {
	return client.GetServerWithContext(context.Background(), serverID)

}

// GetServerWithContext fetches a specific server via serverID
func (client *Client) GetServerWithContext(ctx context.Context, serverID string) (Server, error) {
	res := Server{}
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("servers/%s", serverID),
		TokenType: account_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// EditServer calls EditServerWithContext with empty context
func (client *Client) EditServer(serverID string, server Server) (Server, error) {
	return client.EditServerWithContext(context.Background(), serverID, server)
}

// EditServerWithContext updates details for a specific server with serverID
func (client *Client) EditServerWithContext(ctx context.Context, serverID string, server Server) (Server, error) {
	res := Server{}
	err := client.doRequest(ctx, parameters{
		Method:    "PUT",
		Path:      fmt.Sprintf("servers/%s", serverID),
		TokenType: account_token,
	}, &res)
	return res, err
}
