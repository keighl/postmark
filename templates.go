package postmark

import (
	"fmt"
	"net/url"
)

// Template represents an email template on the server
type Template struct {
	// TemplateId: ID of template
	TemplateId int64
	// Name: Name of template
	Name string
	// Subject: The content to use for the Subject when this template is used to send email.
	Subject string
	// HtmlBody: The content to use for the HtmlBody when this template is used to send email.
	HtmlBody string
	// TextBody: The content to use for the TextBody when this template is used to send email.
	TextBody string
	// AssociatedServerId: The ID of the Server with which this template is associated.
	AssociatedServerId int64
	// Active: Indicates that this template may be used for sending email.
	Active bool
}

// TemplateInfo is a limited set of template info returned via Index/Editing endpoints
type TemplateInfo struct {
	// TemplateId: ID of template
	TemplateId int64
	// Name: Name of template
	Name string
	// Active: Indicates that this template may be used for sending email.
	Active bool
}

///////////////////////////////////////
///////////////////////////////////////

// GetTemplate fetches a specific template via TemplateID
func (client *Client) GetTemplate(templateID string) (Template, error) {
	res := Template{}
	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("templates/%s", templateID),
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

type templatesResponse struct {
	TotalCount int64
	Templates  []TemplateInfo
}

// GetTemplates fetches a list of templates on the server
// It returns a TemplateInfo slice, the total template count, and any error that occurred
// Note: TemplateInfo only returns a subset of template attributes, use GetTemplate(id) to
// retrieve all template info.
func (client *Client) GetTemplates(count int64, offset int64) ([]TemplateInfo, int64, error) {
	res := templatesResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("templates?%s", values.Encode()),
		TokenType: server_token,
	}, &res)
	return res.Templates, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// CreateTemplate saves a new template to the server
func (client *Client) CreateTemplate(template Template) (TemplateInfo, error) {
	res := TemplateInfo{}
	err := client.doRequest(parameters{
		Method:    "POST",
		Path:      "templates",
		Payload:   template,
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// EditTemplate updates details for a specific template with templateID
func (client *Client) EditTemplate(templateID string, template Template) (TemplateInfo, error) {
	res := TemplateInfo{}
	err := client.doRequest(parameters{
		Method:    "PUT",
		Path:      fmt.Sprintf("templates/%s", templateID),
		Payload:   template,
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// DeleteTemplate removes a template (with templateID) from the server
func (client *Client) DeleteTemplate(templateID string) error {
	res := APIError{}
	err := client.doRequest(parameters{
		Method:    "DELETE",
		Path:      fmt.Sprintf("templates/%s", templateID),
		TokenType: server_token,
	}, &res)

	if res.ErrorCode != 0 {
		return res
	}

	return err
}

///////////////////////////////////////
///////////////////////////////////////

// ValidateTemplateBody contains the template/render model combination to be validated
type ValidateTemplateBody struct {
	Subject                    string
	TextBody                   string
	HTMLBody                   string `json:"HtmlBody"`
	TestRenderModel            map[string]interface{}
	InlineCSSForHTMLTestRender bool `json:"InlineCssForHtmlTestRender"`
}

// ValidateTemplateResponse contains information as to how the validation went
type ValidateTemplateResponse struct {
	AllContentIsValid      bool
	HTMLBody               Validation `json:"HtmlBody"`
	TextBody               Validation
	Subject                Validation
	SuggestedTemplateModel map[string]interface{}
}

// Validation contains the results of a field's validation
type Validation struct {
	ContentIsValid   bool
	ValidationErrors []ValidationError
	RenderedContent  string
}

// ValidationError contains information about the errors which occurred during validation for a given field
type ValidationError struct {
	Message           string
	Line              int
	CharacterPosition int
}

// ValidateTemplate validates the provided template/render model combination
func (client *Client) ValidateTemplate(validateTemplateBody ValidateTemplateBody) (ValidateTemplateResponse, error) {
	res := ValidateTemplateResponse{}
	err := client.doRequest(parameters{
		Method:    "POST",
		Path:      "templates/validate",
		Payload:   validateTemplateBody,
		TokenType: server_token,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// TemplatedEmail is used to send an email via a template
type TemplatedEmail struct {
	// TemplateId: REQUIRED - The template to use when sending this message.
	TemplateId int64
	// TemplateModel: The model to be applied to the specified template to generate HtmlBody, TextBody, and Subject.
	TemplateModel map[string]interface{} `json:",omitempty"`
	// InlineCss: By default, if the specified template contains an HTMLBody, we will apply the style blocks as inline attributes to the rendered HTML content. You may opt-out of this behavior by passing false for this request field.
	InlineCss bool `json:",omitempty"`
	// From: The sender email address. Must have a registered and confirmed Sender Signature.
	From string `json:",omitempty"`
	// To: REQUIRED Recipient email address. Multiple addresses are comma separated. Max 50.
	To string `json:",omitempty"`
	// Cc recipient email address. Multiple addresses are comma separated. Max 50.
	Cc string `json:",omitempty"`
	// Bcc recipient email address. Multiple addresses are comma separated. Max 50.
	Bcc string `json:",omitempty"`
	// Tag: Email tag that allows you to categorize outgoing emails and get detailed statistics.
	Tag string `json:",omitempty"`
	// Reply To override email address. Defaults to the Reply To set in the sender signature.
	ReplyTo string `json:",omitempty"`
	// Headers: List of custom headers to include.
	Headers []Header `json:",omitempty"`
	// TrackOpens: Activate open tracking for this email.
	TrackOpens bool `json:",omitempty"`
	// Attachments: List of attachments
	Attachments []Attachment `json:",omitempty"`
}

// SendTemplatedEmail sends an email using a template (TemplateId)
func (client *Client) SendTemplatedEmail(email TemplatedEmail) (EmailResponse, error) {
	res := EmailResponse{}
	err := client.doRequest(parameters{
		Method:    "POST",
		Path:      "email/withTemplate",
		Payload:   email,
		TokenType: server_token,
	}, &res)
	return res, err
}

// SendTemplatedEmail sends batch email using a template (TemplateId)
func (client *Client) SendTemplatedEmailBatch(emails []TemplatedEmail) ([]EmailResponse, error) {
	res := []EmailResponse{}
	var formatEmails map[string]interface{} = map[string]interface{}{
		"Messages": emails,
	}
	err := client.doRequest(parameters{
		Method:    "POST",
		Path:      "email/batchWithTemplates",
		Payload:   formatEmails,
		TokenType: server_token,
	}, &res)
	return res, err
}
