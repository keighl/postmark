package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestTemplate(t *testing.T) {
	tMux.HandleFunc(pat.Get("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"Name": "Onboarding Email",
	        "TemplateId": 1234,
	        "Subject": "Hi there, {{Name}}",
	        "HtmlBody": "Hello dear Postmark user. {{Name}}",
	        "TextBody": "{{Name}} is a {{Occupation}}",
	        "AssociatedServerId": 1,
	        "Active": false
		}`))
	})

	res, err := client.Template("1234")
	if err != nil {
		t.Fatalf("Template: %s", err.Error())
	}

	if res.Name != "Onboarding Email" {
		t.Fatalf("Template: wrong name!")
	}
}

func TestTemplates(t *testing.T) {
	tMux.HandleFunc(pat.Get("/templates"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
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
	          }]
		}`))
	})

	res, err := client.Templates()
	if err != nil {
		t.Fatalf("Templates: %s", err.Error())
	}

	if len(res) == 0 {
		t.Fatalf("Templates: unmarshaled to empty")
	}
}

func TestCreateTemplate(t *testing.T) {
	tMux.HandleFunc(pat.Post("/templates"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"TemplateId": 1234,
			"Name": "Onboarding Email",
			"Active": true
		}`))
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
	tMux.HandleFunc(pat.Put("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"TemplateId": 1234,
		  	  "Name": "Onboarding Emailzzzzz",
		  	  "Active": true
		}`))
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
	tMux.HandleFunc(pat.Delete("/templates/:templateID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
		  "ErrorCode": 0,
		  "Message": "Template 1234 removed."
		}`))
	})

	err := client.DeleteTemplate("1234")
	if err != nil {
		t.Fatalf("DeleteTemplate: %s", err.Error())
	}
}
