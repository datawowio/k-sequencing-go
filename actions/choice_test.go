package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestPostChoiceEndpoint(t *testing.T) {
	p := &PostChoice{}
	endpoint, method, path := p.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.API, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/choices", path)
}

func TestGetChoicesEndpoint(t *testing.T) {
	g := &GetChoices{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.API, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/choices", path)
}

func TestGetChoiceEndpoint(t *testing.T) {
	g := &GetChoice{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.API, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/choice", path)
}

func TestGetChoicePayload(t *testing.T) {
	g := &GetChoice{
		ID:       "5a44671ab3957c2ab5c33326",
		CustomID: "5022342340",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := g.Payload(endpoint, method, path)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}

func TestGetListChoicePayload(t *testing.T) {
	g := &GetChoices{
		ID:   "5a44671ab3957c2ab5c33326",
		Item: "1",
		Page: "5",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := g.Payload(endpoint, method, path)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.Item, queryValues.Get("per_page"))
	a.Equal(t, g.Page, queryValues.Get("page"))
}

func TestPostChoicePayload(t *testing.T) {
	p := &PostChoice{
		Instruction:    "Instruction",
		Categories:     []string{"Foo, Bar"},
		Data:           "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
		AllowEmpty:     true,
		PostbackURL:    "http://postback_url.com",
		PostbackMethod: "GET",
		Multiple:       true,
		CustomID:       "24234224",
		StaffID:        55,
	}
	endpoint, method, path := p.Endpoint()
	req, _ := p.Payload(endpoint, method, path)
	a.NotNil(t, req)
}
