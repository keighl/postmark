package postmark

import (
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
	// Color of the server in the rack screen. Purple Blue Turqoise Green Red Yellow Grey
	Color string
	// SmtpApiActivated specifies whether or not SMTP is enabled on this server.
	SmtpApiActivated bool
	// RawEmailEnabled allows raw email to be sent with inbound.
	RawEmailEnabled bool
	// InboundAddress is the inbound email address
	InboundAddress string
	// InboundHookUrl to POST to everytime an inbound event occurs.
	InboundHookUrl string
	// BounceHookUrl to POST to everytime a bounce event occurs.
	BounceHookUrl string
	// OpenHookUrl to POST to everytime an open event occurs.
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

// GetServer fetches a specific server via serverID
func (client *Client) GetServer(serverID string) (Server, error) {
	res := Server{}
	path := fmt.Sprintf("servers/%s", serverID)
	err := client.doRequest("GET", path, nil, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// EditServer updates details for a specific server with serverID
func (client *Client) EditServer(serverID string, server Server) (Server, error) {
	res := Server{}
	path := fmt.Sprintf("servers/%s", serverID)
	err := client.doRequest("PUT", path, server, &res)
	return res, err
}
