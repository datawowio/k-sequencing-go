package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetMessagesEndpoint(t *testing.T) {
	g := &GetMessages{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/messages", path)
}

func TestGetMessageEndpoint(t *testing.T) {
	g := &GetMessage{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/message", path)
}

func TestPostMessage(t *testing.T) {
	p := &PostMessage{}
	endpoint, method, path := p.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/messages", path)
}

func TestGetmessagePayload(t *testing.T) {
	g := &GetMessage{
		ID:       "5a44671ab3957c2ab5c33326",
		CustomID: "5022342340",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := g.Payload(endpoint, method, path)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}

func TestGetListmessagePayload(t *testing.T) {
	g := &GetMessages{
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

func TestPostMessagePayload(t *testing.T) {
	p := &PostMessage{
		Data:           "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
		Instruction:    "Some instruction",
		PostbackURL:    "http://someUrl.co",
		PostbackMethod: "GET",
		CustomID:       "custom_id",
	}
	endpoint, method, path := p.Endpoint()
	req, _ := p.Payload(endpoint, method, path)
	a.NotNil(t, req)
}
