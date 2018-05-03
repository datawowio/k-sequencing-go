package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetClosedQuestionEndpoint(t *testing.T) {
	g := &GetClosedQuestion{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KSeqAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/closed_question", path)
}

func TestGetListClosedQuestionEndpoint(t *testing.T) {
	g := &GetClosedQuestions{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KSeqAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/closed_questions", path)
}

func TestPostCloseQuestionEndpoint(t *testing.T) {
	p := &PostClosedQuestion{}
	endpoint, method, path := p.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KSeqAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/closed_questions", path)
}

func TestGetCloseQuestionPayload(t *testing.T) {
	g := &GetClosedQuestion{
		ID:       "5a44671ab3957c2ab5c33326",
		CustomID: "5022342340",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := g.Payload(endpoint, method, path)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}

func TestGetListClosedQuestionPayload(t *testing.T) {
	g := &GetClosedQuestions{
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

func TestPostCloseQuestionPayload(t *testing.T) {
	p := &PostClosedQuestion{
		Data:           "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
		PostbackURL:    "http://someUrl.url",
		PostbackMethod: "SOME_METHOD",
		CustomID:       "custom_id",
	}
	endpoint, method, path := p.Endpoint()
	req, _ := p.Payload(endpoint, method, path)
	a.NotNil(t, req)
}
