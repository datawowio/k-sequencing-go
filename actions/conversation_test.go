package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetConversationEndpoint(t *testing.T) {
	g := &ConversationGetParams{}
	url, method, endpoint := g.Endpoint()
	a.NotNil(t, url)
	a.Equal(t, config.KiyoTextAPIURL, url)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, endpoint)
	a.Equal(t, "/api/v1/text/text_conversations", endpoint)
}

func TestPostConversationEndpoint(t *testing.T) {
	p := &ConversationParams{}
	url, method, endpoint := p.Endpoint()
	a.NotNil(t, url)
	a.Equal(t, config.KiyoTextAPIURL, url)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, endpoint)
	a.Equal(t, "/api/v1/text/text_conversations", endpoint)
}

func TestConversationGetPayload(t *testing.T) {
	g := &ConversationGetParams{ID: "5a44671ab3957c2ab5c33326"}
	url, method, endpoint := g.Endpoint()
	req, _ := g.Payload(url, method, endpoint)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}

func TestConversationPostPayload(t *testing.T) {
	p := &ConversationParams{
		Conversation:          []string{"おはよ", "襲おうかなとか言えよ糞ガキめ (´◠ω◠`)"},
		CustomConversationIDs: []string{"123", "321"},
		Answers:               []string{},
		PostbackURL:           "http://localhost:3000",
		PostbackMethod:        "GET",
		CustomID:              "321",
	}
	url, method, endpoint := p.Endpoint()
	req, err := p.Payload(url, method, endpoint)
	a.Nil(t, err)
	a.NotNil(t, req)
}
