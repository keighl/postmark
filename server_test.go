package postmark

import (
	"testing"
	"net/http"

	"goji.io/pat"
)

func TestGetCurrentServer(t *testing.T) {
	responseJSON := `{
		"ID": 1,
			"Name": "Staging Testing",
			"ApiTokens": [
		"server token"
		],
		"ServerLink": "https://postmarkapp.com/servers/1/overview",
			"Color": "red",
			"SmtpApiActivated": true,
			"RawEmailEnabled": false,
			"DeliveryHookUrl": "http://hooks.example.com/delivery",
			"InboundAddress": "yourhash@inbound.postmarkapp.com",
			"InboundHookUrl": "http://hooks.example.com/inbound",
			"BounceHookUrl": "http://hooks.example.com/bounce",
			"IncludeBounceContentInHook": true,
			"OpenHookUrl": "http://hooks.example.com/open",
			"PostFirstOpenOnly": false,
			"TrackOpens": false,
			"TrackLinks" : "None",
			"ClickHookUrl" : "http://hooks.example.com/click",
			"InboundDomain": "",
			"InboundHash": "yourhash",
			"InboundSpamThreshold": 0
	}`

	tMux.HandleFunc(pat.Get("/server"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetCurrentServer()
	if err != nil {
		t.Fatalf("GetCurrentServer: %s", err.Error())
	}

	if res.Name != "Staging Testing" {
		t.Fatalf("GetCurrentServer: wrong name!: %s", res.Name)
	}
}

func TestEditCurrentServer(t *testing.T) {
	responseJSON := `{
  "ID": 1,
  "Name": "Production Testing",
  "ApiTokens": [
    "Server Token"
  ],
  "ServerLink": "https://postmarkapp.com/servers/1/overview",
  "Color": "blue",
  "SmtpApiActivated": false,
  "RawEmailEnabled": false,
  "DeliveryHookUrl": "http://hooks.example.com/delivery",
  "InboundAddress": "yourhash@inbound.postmarkapp.com",
  "InboundHookUrl": "http://hooks.example.com/inbound",
  "BounceHookUrl": "http://hooks.example.com/bounce",
  "IncludeBounceContentInHook": true,
  "OpenHookUrl": "http://hooks.example.com/open",
  "PostFirstOpenOnly": false,
  "TrackOpens": false,
  "TrackLinks": "None",
  "ClickHookUrl": "http://hooks.example.com/click",
  "InboundDomain": "",
  "InboundHash": "yourhash",
  "InboundSpamThreshold": 10
}`
	tMux.HandleFunc(pat.Put("/server"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.EditCurrentServer(Server{
		Name: "Production Testing",
	})

	if err != nil {
		t.Fatalf("EditCurrentServer: %s", err.Error())
	}

	if res.Name != "Production Testing" {
		t.Fatalf("EditCurrentServer: wrong name!: %s", res.Name)
	}
}
