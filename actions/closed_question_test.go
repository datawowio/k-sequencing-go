package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetEndpoint(t *testing.T) {
	g := &GetClosedQuestion{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.LocalAPI, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/closed_question", path)
}

func TestPostEndpoint(t *testing.T) {
	g := &PostClosedQuestion{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.LocalAPI, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/closed_questions", path)
}

func TestGetPayload(t *testing.T) {
	g := &GetClosedQuestion{
		ID:       "5a44671ab3957c2ab5c33326",
		CustomID: "5022342340",
	}
	req := g.Payload()
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}

func TestPostPayload(t *testing.T) {
	p := &PostClosedQuestion{
		Data: "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
	}
	req := p.Payload()
	a.NotNil(t, req)
}
