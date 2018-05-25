package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetSenderSignatures(t *testing.T) {
	responseJSON := `{
	"TotalCount": 2,
	"SenderSignatures": [
	  {
		"Domain": "wildbit.com",
		"EmailAddress": "jp@wildbit.com",
		"ReplyToEmailAddress": "info@wildbit.com",
		"Name": "JP Toto",
		"Confirmed": true,
		"ID": 36735
	  },
	  {
		"Domain": "example.com",
		"EmailAddress": "jp@example.com",
		"ReplyToEmailAddress": "",
		"Name": "JP Toto",
		"Confirmed": true,
		"ID": 81605
	  }
	]
  }`

	tMux.HandleFunc(pat.Get("/senders"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetSenderSignatures(50, 0)
	if err != nil {
		t.Fatalf("GetSenderSignatures: %s", err.Error())
	}

	if res.TotalCount != 2 {
		t.Fatalf("GetSenderSignatures: wrong TotalCount!")
	}
}
