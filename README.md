# Postmark

[![Build Status](https://travis-ci.org/keighl/postmark.png?branch=master)](https://travis-ci.org/keighl/postmark) [![Go Report Card](https://goreportcard.com/badge/github.com/keighl/postmark)](https://goreportcard.com/report/github.com/keighl/postmark)  [![codecov.io](https://codecov.io/github/keighl/postmark/coverage.svg?branch=master)](https://codecov.io/github/keighl/postmark?branch=master) [![GoDoc](https://godoc.org/github.com/keighl/postmark?status.svg)](https://godoc.org/github.com/keighl/postmark)

### Installation

    go get -u github.com/keighl/postmark

### API Coverage

* [x] Emails
    * [x] Send a single email
    * [x] Send a batch emails  
* [ ] Bounce
    * [x] Get delivery stats
    * [x] Get bounces
    * [x] Get a single bounce
    * [ ] Get bounce dump
    * [ ] Activate a bounce
    * [ ] Get bounced tags
    * [ ] Bounce types  
* [ ] Templates
    * [x] Get a template
    * [x] Create a template
    * [x] Edit a template
    * [x] List templates
    * [x] Delete a template
    * [ ] Validate a template
    * [x] Send email with template
* [ ] Servers
    * [ ] Get the server
    * [ ] Edit the server
* [ ] Messages
    * [ ] Outbound message search
    * [ ] Outbound message details
    * [ ] Outbound message dump
    * [ ] Inbound message search
    * [ ] Inbound message details
    * [ ] Bypass blocked inbound message
    * [ ] Retry a failed inbound message
    * [ ] Message opens
    * [ ] Opens for a single message
* Sender signatures
    * [ ] List sender signatures
    * [ ] Get a sender signature’s details
    * [ ] Create a signature
    * [ ] Edit a signature
    * [ ] Delete a signature
    * [ ] Resend a confirmation
    * [ ] Verify an SPF record
    * [ ] Request a new DKIM
* Stats
    * [ ] Get outbound overview
    * [ ] Get sent counts
    * [ ] Get bounce counts
    * [ ] Get spam complaints
    * [ ] Get tracked email counts
    * [ ] Get email open counts
    * [ ] Get email platform usage
    * [ ] Get email client usage
    * [ ] Get email read times
* Triggers
    * Tags triggers
        * [ ] Create a trigger for a tag
        * [ ] Get a single trigger
        * [ ] Edit a single trigger
        * [ ] Delete a single trigger
        * [ ] Search triggers
    * Inbound rules triggers
        * [ ] Create a trigger for inbound rule
        * [ ] Delete a single trigger
        * [ ] List triggers    
