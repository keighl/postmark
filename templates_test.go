package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetTemplate(t *testing.T) {
	responseJSON := `{
		"Name": "Onboarding Email",
		"TemplateId": 1234,
		"Subject": "Hi there, {{Name}}",
		"HtmlBody": "Hello dear Postmark user. {{Name}}",
		"TextBody": "{{Name}} is a {{Occupation}}",
		"AssociatedServerId": 1,
		"Active": false
	}`

	tMux.HandleFunc(pat.Get("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetTemplate("1234")
	if err != nil {
		t.Fatalf("Template: %s", err.Error())
	}

	if res.Name != "Onboarding Email" {
		t.Fatalf("Template: wrong name!")
	}
}

func TestGetTemplates(t *testing.T) {
	responseJSON := `{
		"TotalCount": 2,
		"Templates": [
		  {
			"Active": true,
			"TemplateId": 1234,
			"Name": "Account Activation Email"
		  },
		  {
			"Active": true,
			"TemplateId": 5678,
			"Name": "Password Recovery Email"
		  }
		]
	}`

	tMux.HandleFunc(pat.Get("/templates"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, count, err := client.GetTemplates(100, 10)
	if err != nil {
		t.Fatalf("GetTemplates: %s", err.Error())
	}

	if len(res) == 0 {
		t.Fatalf("GetTemplates: unmarshaled to empty")
	}

	if count != 2 {
		t.Fatalf("GetTemplates: unmarshaled to empty")
	}
}

func TestCreateTemplate(t *testing.T) {
	responseJSON := `{
		"TemplateId": 1234,
		"Name": "Onboarding Email",
		"Active": true
	}`

	tMux.HandleFunc(pat.Post("/templates"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.CreateTemplate(Template{
		Name:     "Onboarding Email",
		Subject:  "Hello from {{company.name}}!",
		TextBody: "Hello, {{name}}!",
		HtmlBody: "<html><body>Hello, {{name}}!</body></html>",
	})

	if err != nil {
		t.Fatalf("CreateTemplate: %s", err.Error())
	}

	if res.Name != "Onboarding Email" {
		t.Fatalf("CreateTemplate: wrong name!")
	}
}

func TestEditTemplate(t *testing.T) {
	responseJSON := `{
		"TemplateId": 1234,
		  "Name": "Onboarding Emailzzzzz",
		  "Active": true
	}`

	tMux.HandleFunc(pat.Put("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.EditTemplate("1234", Template{
		Name:     "Onboarding Emailzzzzz",
		Subject:  "Hello from {{company.name}}!",
		TextBody: "Hello, {{name}}!",
		HtmlBody: "<html><body>Hello, {{name}}!</body></html>",
	})
	if err != nil {
		t.Fatalf("EditTemplate: %s", err.Error())
	}

	if res.Name != "Onboarding Emailzzzzz" {
		t.Fatalf("EditTemplate: wrong name!")
	}
}

func TestDeleteTemplate(t *testing.T) {
	responseJSON := `{
	  "ErrorCode": 0,
	  "Message": "Template 1234 removed."
	}`

	tMux.HandleFunc(pat.Delete("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	// Success
	err := client.DeleteTemplate("1234")
	if err != nil {
		t.Fatalf("DeleteTemplate: %s", err.Error())
	}

	// Failure
	responseJSON = `{
	  "ErrorCode": 402,
	  "Message": "Invalid JSON"
	}`

	err = client.DeleteTemplate("1234")
	if err == nil {
		t.Fatalf("DeleteTemplate  should have failed")
	}
}

func TestValidateTemplate(t *testing.T) {
	responseJSON := `{
		{
			"AllContentIsValid": true,
			"HtmlBody": {
			  "ContentIsValid": true,
			  "ValidationErrors": [],
			  "RenderedContent": "address_Value name_Value "
			},
			"TextBody": {
			  "ContentIsValid": true,
			  "ValidationErrors": [{
				  "Message" : "The syntax for this template is invalid.",
				  "Line" : 1,
				  "CharacterPosition" : 1
			  }],
			  "RenderedContent": "phone_Value name_Value "
			},
			"Subject": {
			  "ContentIsValid": true,
			  "ValidationErrors": [],
			  "RenderedContent": "name_Value subjectHeadline_Value"
			},
			"SuggestedTemplateModel": {
			  "userName": "bobby joe",
			  "company": {
				"address": "address_Value",
				"phone": "phone_Value",
				"name": "name_Value"
			  },
			  "person": [
				{
				  "name": "name_Value"
				}
			  ],
			  "subjectHeadline": "subjectHeadline_Value"
			}
		  }
	}`

	tMux.HandleFunc(pat.Delete("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.ValidateTemplate(ValidateTemplateBody{
		Subject:  "{{#company}}{{name}}{{/company}} {{subjectHeadline}}",
		TextBody: "{{#company}}{{address}}{{/company}}{{#each person}} {{name}} {{/each}}",
		HTMLBody: "{{#company}}{{phone}}{{/company}}{{#each person}} {{name}} {{/each}}",
		TestRenderModel: map[string]interface{}{
			"userName": "bobby joe",
		},
		InlineCSSForHTMLTestRender: false,
	})

	if err != nil {
		t.Fatalf("ValidateTemplate: %s", err.Error())
	}

	if !res.AllContentIsValid {
		t.Fatalf("ValidateTemplate: AllContentIsValid should be true")
	}
}

var testTemplatedEmail = TemplatedEmail{
	TemplateId: 1234,
	TemplateModel: map[string]interface{}{
		"user_name": "John Smith",
		"company": map[string]interface{}{
			"name": "ACME",
		},
	},
	InlineCss: true,
	From:      "sender@example.com",
	To:        "receiver@example.com",
	Cc:        "copied@example.com",
	Bcc:       "blank-copied@example.com",
	Tag:       "Invitation",
	ReplyTo:   "reply@example.com",
	Headers: []Header{
		{
			Name:  "CUSTOM-HEADER",
			Value: "value",
		},
	},
	TrackOpens: true,
	Attachments: []Attachment{
		{
			Name:        "readme.txt",
			Content:     "dGVzdCBjb250ZW50",
			ContentType: "text/plain",
		},
		{
			Name:        "report.pdf",
			Content:     "dGVzdCBjb250ZW50",
			ContentType: "application/octet-stream",
		},
	},
}

func TestSendTemplatedEmail(t *testing.T) {
	responseJSON := `{
		"To": "receiver@example.com",
		"SubmittedAt": "2014-02-17T07:25:01.4178645-05:00",
		"MessageID": "0a129aee-e1cd-480d-b08d-4f48548ff48d",
		"ErrorCode": 0,
		"Message": "OK"
	}`

	tMux.HandleFunc(pat.Post("/email/withTemplate"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.SendTemplatedEmail(testTemplatedEmail)
	if err != nil {
		t.Fatalf("SendTemplatedEmail: %s", err.Error())
	}

	if res.MessageID != "0a129aee-e1cd-480d-b08d-4f48548ff48d" {
		t.Fatalf("SendTemplatedEmail: incorrect message ID")
	}
}
