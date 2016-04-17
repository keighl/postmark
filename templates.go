package postmark

import (
	"fmt"
)

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

func (client *Client) Template(templateID string) (Template, error) {
	res := Template{}
	path := fmt.Sprintf("templates/%s", templateID)
	err := client.doRequest("GET", path, nil, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

///////////////////////////////////////
///////////////////////////////////////

type templatesResponse struct {
	TotalCount int64
	Templates  []TemplateInfo
}

func (client *Client) Templates() ([]TemplateInfo, error) {
	res := templatesResponse{}
	// TODO include limit/offset
	err := client.doRequest("GET", "templates", nil, &res)
	if err != nil {
		return res.Templates, err
	}

	return res.Templates, nil
}

///////////////////////////////////////
///////////////////////////////////////

func (client *Client) CreateTemplate(template Template) (TemplateInfo, error) {
	res := TemplateInfo{}
	err := client.doRequest("POST", "templates", template, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

///////////////////////////////////////
///////////////////////////////////////

func (client *Client) EditTemplate(templateID string, template Template) (TemplateInfo, error) {
	res := TemplateInfo{}
	path := fmt.Sprintf("templates/%s", templateID)
	err := client.doRequest("PUT", path, template, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

///////////////////////////////////////
///////////////////////////////////////

func (client *Client) DeleteTemplate(templateID string) error {
	errRes := APIError{}
	path := fmt.Sprintf("templates/%s", templateID)
	err := client.doRequest("DELETE", path, nil, &errRes)
	if err != nil {
		return err
	}

	if errRes.ErrorCode == 0 {
		return nil
	}

	return errRes
}
