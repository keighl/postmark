package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestSendEmail(t *testing.T) {
	tMux.HandleFunc(pat.Post("/email"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
			"To": "receiver@example.com",
			"SubmittedAt": "2014-02-17T07:25:01.4178645-05:00",
			"MessageID": "0a129aee-e1cd-480d-b08d-4f48548ff48d",
			"ErrorCode": 0,
			"Message": "OK"
		}`))
	})

	email := Email{
		From:     "sender@example.com",
		To:       "receiver@example.com",
		Cc:       "copied@example.com",
		Bcc:      "blank-copied@example.com",
		Subject:  "Test",
		Tag:      "Invitation",
		HtmlBody: "<b>Hello</b>",
		TextBody: "Hello",
		ReplyTo:  "reply@example.com",
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

	res, err := client.SendEmail(email)

	if err != nil {
		t.Fatalf("SendEmail: %s", err.Error())
	}

	if res.MessageID != "0a129aee-e1cd-480d-b08d-4f48548ff48d" {
		t.Fatalf("SendEmail: wrong id!")
	}
}
