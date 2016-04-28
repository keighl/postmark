# Postmark

[![Build Status](https://travis-ci.org/keighl/postmark.png?branch=master)](https://travis-ci.org/keighl/postmark) [![Go Report Card](https://goreportcard.com/badge/github.com/keighl/postmark)](https://goreportcard.com/report/github.com/keighl/postmark)  [![codecov.io](https://codecov.io/github/keighl/postmark/coverage.svg?branch=master)](https://codecov.io/github/keighl/postmark?branch=master) [![GoDoc](https://godoc.org/github.com/keighl/postmark?status.svg)](https://godoc.org/github.com/keighl/postmark)

A Golang package for the using Postmark API.

### Installation

    go get -u github.com/keighl/postmark

### API Coverage

* [x] Emails
    * [x] `POST /email`
    * [x] `POST /email/batch`
    * [x] `POST /email/withTemplate`
* [x] Bounces
    * [x] `GET /deliverystats`
    * [x] `GET /bounces`
    * [x] `GET /bounces/:id`
    * [x] `GET /bounces/:id/dump`
    * [x] `PUT /bounces/:id/activate`
    * [x] `GET /bounces/tags`
* [ ] Templates
    * [x] `GET /templates`
    * [x] `POST /templates`
    * [x] `GET /templates/:id`
    * [x] `PUT /templates/:id`
    * [x] `DELETE /templates/:id`
    * [ ] `POST /templates/validate`
* [x] Servers
    * [x] `GET /servers/:id`
    * [x] `PUT /servers/:id`
* [x] Outbound Messages
    * [x] `GET /messages/outbound`
    * [x] `GET /messages/outbound/:id/details`
    * [x] `GET /messages/outbound/:id/dump`
    * [x] `GET /messages/outbound/opens`
    * [x] `GET /messages/outbound/opens/:id`
* [ ] Inbound Messages
    * [x] `GET /messages/inbound`
    * [x] `GET /messages/inbound/:id/details`
    * [ ] Bypass blocked inbound message
    * [ ] Retry a failed inbound message
* [ ] Sender signatures
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
