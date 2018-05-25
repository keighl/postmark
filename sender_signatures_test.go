package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestListSenderSignatures(t *testing.T) {
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

	res, err := client.ListSenderSignatures(50, 0)
	if err != nil {
		t.Fatalf("ListSenderSignatures: %s", err.Error())
	}

	if res.TotalCount != 2 {
		t.Fatalf("ListSenderSignatures: wrong name!")
	}
}
