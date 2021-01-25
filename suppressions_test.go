package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetSuppressionEmails(t *testing.T) {
	responseJSON := `{
		"Suppressions":[
		  {
			"EmailAddress":"address@wildbit.com",
			"SuppressionReason":"ManualSuppression",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-10T08:58:33-05:00"
		  },
		  {
			"EmailAddress":"bounce.address@wildbit.com",
			"SuppressionReason":"HardBounce",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-11T08:58:33-05:00"
		  },
		  {
			"EmailAddress":"spam.complaint.address@wildbit.com",
			"SuppressionReason":"SpamComplaint",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-12T08:58:33-05:00"
		  }
		]
	  }`

	tMux.HandleFunc(pat.Get("/message-streams/:StreamID/suppressions/dump"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetSuppressionEmails("outbound", nil)

	if err != nil {
		t.Fatalf("GetSuppressionEmails: %s", err.Error())
	}

	if len(res) != 3 {
		t.Fatalf("GetSuppressionEmails: wrong number of suppression (%d)", len(res))
	}

	if res[0].EmailAddress != "address@wildbit.com" {
		t.Fatalf("GetSuppressionEmails: wrong suppression email address: %s", res[0].EmailAddress)
	}

	responseJSON = `{
		"Suppressions":[
		  {
			"EmailAddress":"address@wildbit.com",
			"SuppressionReason":"ManualSuppression",
			"Origin": "Recipient",
			"CreatedAt":"2019-12-10T08:58:33-05:00"
		  }
		]
	  }`

	tMux.HandleFunc(pat.Get("/message-streams/:StreamID/suppressions/dump"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err = client.GetSuppressionEmails("outbound", map[string]interface{}{
		"emailaddress":      "address@wildbit.com",
		"fromdate":          "2019-12-10",
		"todate":            "2019-12-11",
		"suppressionreason": HardBounceReason,
		"origin":            RecipientOrigin,
	})

	if len(res) != 1 {
		t.Fatalf("GetSuppressionEmails: wrong number of suppression (%d)", len(res))
	}

	if res[0].EmailAddress != "address@wildbit.com" {
		t.Fatalf("GetSuppressionEmails: wrong suppression email address: %s", res[0].EmailAddress)
	}

}

func TestDeleteSuppressionEmails(t *testing.T) {
	responseJSON := `{
		"Suppressions":[
		  {
			"EmailAddress":"good.address@wildbit.com",
			"Status":"Deleted",
			"Message": null
		  },
		  {
			"EmailAddress":"not.suppressed@wildbit.com",
			"Status":"Deleted",
			"Message": null
		  },
		  {
			"EmailAddress":"spammy.address@wildbit.com",
			"Status":"Failed",
			"Message": "You do not have the required authority to change this suppression."
		  }
		]
	  }`

	tMux.HandleFunc(pat.Post("/message-streams/:StreamID/suppressions/delete"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.DeleteSuppressionEmails("outbound", []string{"good.address@wildbit.com", "not.suppressed@wildbit.com", "spammy.address@wildbit.com"})

	if err != nil {
		t.Fatalf("DeleteSuppressionEmails: %s", err.Error())
	}

	if len(res) != 3 {
		t.Fatalf("DeleteSuppressionEmails: wrong number of suppression (%d)", len(res))
	}
}
