package postmark

import (
	"fmt"
	"net/url"
)

///////////////////////////////////////
///////////////////////////////////////

// OutboundStats - a brief overview of statistics for all of your outbound email.
type OutboundStats struct {
	// Sent - Number of sent emails
	Sent int64
	// Bounced - Number of bounced emails
	Bounced int64
	// SMTPApiErrors - Number of SMTP errors
	SMTPApiErrors int64
	// BounceRate - Bounce rate percentage calculated by total sent.
	BounceRate float64
	// SpamComplaints - Number of spam complaints received
	SpamComplaints int64
	// SpamComplaintsRate - Spam complaints percentage calculated by total sent.
	SpamComplaintsRate float64
	// Opens - Number of opens
	Opens int64
	// UniqueOpens - Number of unique opens
	UniqueOpens int64
	// Tracked - Number of tracked emails sent
	Tracked int64
	// WithClientRecorded - Number of emails where the client was successfully tracked.
	WithClientRecorded int64
	// WithPlatformRecorded - Number of emails where platform was successfully tracked.
	WithPlatformRecorded int64
	// WithReadTimeRecorded - Number of emails where read time was successfully tracked.
	WithReadTimeRecorded int64
}

// GetOutboundStats - Gets a brief overview of statistics for all of your outbound email.
// Available options: http://developer.postmarkapp.com/developer-api-stats.html#overview
func (client *Client) GetOutboundStats(options map[string]interface{}) (OutboundStats, error) {
	res := OutboundStats{}

	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// SendDay - send stats for a specific day
type SendDay struct {
	// Date - prettttay self explanatory
	Date string
	// Sent - number of emails sent
	Sent int64
}

// SendCounts - send stats for a period
type SendCounts struct {
	// Days - List of objects that each represent sent counts by date.
	Days []SendDay
	// Sent - Indicates the number of total sent emails returned.
	Sent int64
}

// GetSentCounts - Gets a total count of emails you’ve sent out.
// Available options: http://developer.postmarkapp.com/developer-api-stats.html#sent-counts
func (client *Client) GetSentCounts(options map[string]interface{}) (SendCounts, error) {
	res := SendCounts{}
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound/sends?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// BounceDay - bounce stats for a specific day
type BounceDay struct {
	// Date - prettttay self explanatory
	Date string
	// HardBounce - number of hard bounces
	HardBounce int64
	// SoftBounce - number of soft bounces
	SoftBounce int64
	// SMTPApiError - number of SMTP errors
	SMTPApiError int64
	// Transient - number of transient bounces.
	Transient int64
}

// BounceCounts - bounce stats for a period
type BounceCounts struct {
	// Days - List of objects that each represent sent counts by date.
	Days []BounceDay
	// HardBounce - total number of hard bounces
	HardBounce int64
	// SoftBounce - total number of soft bounces
	SoftBounce int64
	// SMTPApiError - total number of SMTP errors
	SMTPApiError int64
	// Transient - total number of transient bounces.
	Transient int64
}

// GetBounceCounts - Gets total counts of emails you’ve sent out that have been returned as bounced.
// Available options: http://developer.postmarkapp.com/developer-api-stats.html#bounce-counts
func (client *Client) GetBounceCounts(options map[string]interface{}) (BounceCounts, error) {
	res := BounceCounts{}
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound/bounces?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// SpamDay - spam complaints for a specific day
type SpamDay struct {
	// Date - prettttay self explanatory
	Date string
	// SpamComplaint - number of spam complaints received
	SpamComplaint int64
}

// SpamCounts - spam complaints for a period
type SpamCounts struct {
	// Days - List of objects that each represent spam complaint counts by date.
	Days []SpamDay
	// SpamComplaint - Indicates total number of spam complaints.
	SpamComplaint int64
}

// GetSpamCounts - Gets a total count of recipients who have marked your email as spam.
// Days that did not produce statistics won’t appear in the JSON response.
// Available options: http://developer.postmarkapp.com/developer-api-stats.html#spam-complaints
func (client *Client) GetSpamCounts(options map[string]interface{}) (SpamCounts, error) {
	res := SpamCounts{}
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound/spam?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// TrackedDay - tracked emails sent on a specific day
type TrackedDay struct {
	// Date - prettttay self explanatory
	Date string
	// Tracked - number of emails tracked sent
	Tracked int64
}

// TrackedCounts - tracked emails sent for a period
type TrackedCounts struct {
	// Days - List of objects that each represent tracked email counts by date.
	Days []TrackedDay
	// Tracked - Indicates total number of tracked emails sent.
	Tracked int64
}

// GetTrackedCounts - Gets a total count of emails you’ve sent with open tracking enabled.
// Available options: http://developer.postmarkapp.com/developer-api-stats.html#email-tracked-count
func (client *Client) GetTrackedCounts(options map[string]interface{}) (TrackedCounts, error) {
	res := TrackedCounts{}
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound/tracked?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// OpenedDay - opened outbound emails sent on a specific day
type OpenedDay struct {
	// Date - prettttay self explanatory
	Date string
	// Opens - Indicates total number of opened emails. This total includes recipients who opened your email multiple times.
	Opens int64
	// Unique - Indicates total number of uniquely opened emails.
	Unique int64
}

// OpenCounts - opened outbound emails for a period
type OpenCounts struct {
	// Days - List of objects that each represent opens by date.
	Days []OpenedDay
	// Opens - Indicates total number of opened emails. This total includes recipients who opened your email multiple times.
	Opens int64
	// Unique int64 - Indicates total number of uniquely opened emails.
	Unique int64
}

// GetOpenCounts - Gets total counts of recipients who opened your emails. This is only recorded when open tracking is enabled for that email.
// Available options: http://developer.postmarkapp.com/developer-api-stats.html#email-opens-count
func (client *Client) GetOpenCounts(options map[string]interface{}) (OpenCounts, error) {
	res := OpenCounts{}
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound/opens?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// PlatformCounts contains day-to-day usages, along with totals of email usages by platform
type PlatformCounts struct {
	// Days - List of objects that each represent email platform usages by date
	Days []PlatformDay
	// Desktop - The total number of email platform usages by Desktop
	Desktop int64

	// Mobile - The total number of email platform usages by Mobile
	Mobile int64

	// Unknown - The total number of email platform usages by others
	Unknown int64

	// WebMail - The total number of email platform usages by WebMail
	WebMail int64
}

// PlatformDay contains the totals of email usages by platform for a specific date
type PlatformDay struct {
	// Date - the date in question
	Date string

	// Desktop - The total number of email platform usages by Desktop for this date
	Desktop int64

	// Mobile - The total number of email platform usages by Mobile for this date
	Mobile int64

	// Unknown - The total number of email platform usages by others for this date
	Unknown int64

	// WebMail - The total number of email platform usages by WebMail for this date
	WebMail int64
}

// GetPlatformCounts gets the email platform usage
func (client *Client) GetPlatformCounts(options map[string]interface{}) (PlatformCounts, error) {
	res := PlatformCounts{}
	values := &url.Values{}
	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("stats/outbound/platform?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res, err
}
