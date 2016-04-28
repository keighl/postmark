package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetServer(t *testing.T) {
	tMux.HandleFunc(pat.Get("/servers/:serverID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
      "ID": 1,
      "Name": "Staging Testing",
      "ApiTokens": [
        "server token"
      ],
      "ServerLink": "https://postmarkapp.com/servers/1/overview",
      "Color": "red",
      "SmtpApiActivated": true,
      "RawEmailEnabled": false,
      "InboundAddress": "yourhash@inbound.postmarkapp.com",
      "InboundHookUrl": "http://hooks.example.com/inbound",
      "BounceHookUrl": "http://hooks.example.com/bounce",
      "OpenHookUrl": "http://hooks.example.com/open",
      "PostFirstOpenOnly": false,
      "TrackOpens": false,
      "InboundDomain": "",
      "InboundHash": "yourhash",
      "InboundSpamThreshold": 0
		}`))
	})

	res, err := client.GetServer("1")
	if err != nil {
		t.Fatalf("GetServer: %s", err.Error())
	}

	if res.Name != "Staging Testing" {
		t.Fatalf("GetServer: wrong name!: %s", res.Name)
	}
}

func TestEditServer(t *testing.T) {
	tMux.HandleFunc(pat.Put("/servers/:serverID"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`{
      "ID": 1,
      "Name": "Production Testing",
      "ApiTokens": [
        "Server Token"
      ],
      "ServerLink": "https://postmarkapp.com/servers/1/overview",
      "Color": "blue",
      "SmtpApiActivated": false,
      "RawEmailEnabled": false,
      "InboundAddress": "yourhash@inbound.postmarkapp.com",
      "InboundHookUrl": "http://hooks.example.com/inbound",
      "BounceHookUrl": "http://hooks.example.com/bounce",
      "OpenHookUrl": "http://hooks.example.com/open",
      "PostFirstOpenOnly": false,
      "TrackOpens": false,
      "InboundDomain": "",
      "InboundHash": "yourhash",
      "InboundSpamThreshold": 10
		}`))
	})

	res, err := client.EditServer("1234", Server{
		Name: "Production Testing",
	})

	if err != nil {
		t.Fatalf("EditServer: %s", err.Error())
	}

	if res.Name != "Production Testing" {
		t.Fatalf("EditServer: wrong name!: %s", res.Name)
	}
}
