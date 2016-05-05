package postmark

import (
	"net/http"
	"testing"

	"goji.io/pat"
)

func TestGetOutboundStats(t *testing.T) {
	responseJSON := `{
	  "Sent": 615,
	  "Bounced": 64,
	  "SMTPApiErrors": 25,
	  "BounceRate": 10.406,
	  "SpamComplaints": 10,
	  "SpamComplaintsRate": 1.626,
	  "Opens": 166,
	  "UniqueOpens": 26,
	  "Tracked": 111,
	  "WithClientRecorded": 14,
	  "WithPlatformRecorded": 10,
	  "WithReadTimeRecorded": 10
	}`

	tMux.HandleFunc(pat.Get("/stats/outbound"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetOutboundStats(map[string]interface{}{
		"fromdate": "2014-01-01",
		"todate":   "2014-02-01",
	})

	if err != nil {
		t.Fatalf("GetOutboundStats: %v", err.Error())
	}

	if res.Sent != 615 {
		t.Fatalf("GetOutboundStats: wrong Sent: %v", res.Sent)
	}
}

func TestGetSentCounts(t *testing.T) {
	responseJSON := `{
	  "Days": [
	    {
	      "Date": "2014-01-01",
	      "Sent": 140
	    },
	    {
	      "Date": "2014-01-02",
	      "Sent": 160
	    },
	    {
	      "Date": "2014-01-04",
	      "Sent": 50
	    },
	    {
	      "Date": "2014-01-05",
	      "Sent": 115
	    }
	  ],
	  "Sent": 615
	}`

	tMux.HandleFunc(pat.Get("/stats/outbound/sends"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetSentCounts(map[string]interface{}{
		"fromdate": "2014-01-01",
		"todate":   "2014-02-01",
	})

	if err != nil {
		t.Fatalf("GetSentCounts: %v", err.Error())
	}

	if res.Sent != 615 {
		t.Fatalf("GetSentCounts: wrong Sent: %v", res.Sent)
	}

	if res.Days[0].Sent != 140 {
		t.Fatalf("GetSentCounts: wrong day Sent count")
	}
}

func TestGetBounceCounts(t *testing.T) {
	responseJSON := `{
	  "Days": [
	    {
	      "Date": "2014-01-01",
	      "HardBounce": 12,
	      "SoftBounce": 36
	    },
	    {
	      "Date": "2014-01-03",
	      "Transient": 7
	    },
	    {
	      "Date": "2014-01-04",
	      "Transient": 4
	    },
	    {
	      "Date": "2014-01-05",
	      "SMTPApiError": 25,
	      "Transient": 5
	    }
	  ],
	  "HardBounce": 12,
	  "SMTPApiError": 25,
	  "SoftBounce": 36,
	  "Transient": 16
	}`

	tMux.HandleFunc(pat.Get("/stats/outbound/bounces"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetBounceCounts(map[string]interface{}{
		"fromdate": "2014-01-01",
		"todate":   "2014-02-01",
	})

	if err != nil {
		t.Fatalf("GetBounceCounts: %v", err.Error())
	}

	if res.HardBounce != 12 {
		t.Fatalf("GetBounceCounts: wrong HardBounce: %v", res.HardBounce)
	}

	if res.Days[0].HardBounce != 12 {
		t.Fatalf("GetBounceCounts: wrong day HardBounce count")
	}
}

func TestGetSpamCounts(t *testing.T) {
	responseJSON := `{
	  "Days": [
	    {
	      "Date": "2014-01-01",
	      "SpamComplaint": 2
	    },
	    {
	      "Date": "2014-01-02",
	      "SpamComplaint": 3
	    },
	    {
	      "Date": "2014-01-05",
	      "SpamComplaint": 5
	    }
	  ],
	  "SpamComplaint": 10
	}`

	tMux.HandleFunc(pat.Get("/stats/outbound/spam"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetSpamCounts(map[string]interface{}{
		"fromdate": "2014-01-01",
		"todate":   "2014-02-01",
	})

	if err != nil {
		t.Fatalf("GetSpamCounts: %v", err.Error())
	}

	if res.SpamComplaint != 10 {
		t.Fatalf("GetSpamCounts: wrong SpamComplaint: %v", res.SpamComplaint)
	}

	if res.Days[0].SpamComplaint != 2 {
		t.Fatalf("GetSpamCounts: wrong day SpamComplaint count")
	}
}

func TestGetTrackedCounts(t *testing.T) {
	responseJSON := `{
	  "Days": [
	    {
	      "Date": "2014-01-01",
	      "Tracked": 24
	    },
	    {
	      "Date": "2014-01-02",
	      "Tracked": 26
	    },
	    {
	      "Date": "2014-01-03",
	      "Tracked": 15
	    },
	    {
	      "Date": "2014-01-04",
	      "Tracked": 15
	    },
	    {
	      "Date": "2014-01-05",
	      "Tracked": 31
	    }
	  ],
	  "Tracked": 111
	}`

	tMux.HandleFunc(pat.Get("/stats/outbound/tracked"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(responseJSON))
	})

	res, err := client.GetTrackedCounts(map[string]interface{}{
		"fromdate": "2014-01-01",
		"todate":   "2014-02-01",
	})

	if err != nil {
		t.Fatalf("GetTrackedCounts: %v", err.Error())
	}

	if res.Tracked != 111 {
		t.Fatalf("GetTrackedCounts: wrong Tracked: %v", res.Tracked)
	}

	if res.Days[0].Tracked != 24 {
		t.Fatalf("GetTrackedCounts: wrong day Tracked count")
	}
}
