package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetAiConsensusEndpoint(t *testing.T) {
	aiConsensus := &GetAiConsensus{}
	endpoint, method, path := aiConsensus.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/v1/jobs/ai/consensuses", path)
}

func TestGetListAiConsensusEndpoint(t *testing.T) {
	aiConsensus := &GetAiConsensuses{}
	endpoint, method, path := aiConsensus.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/v1/jobs/ai/consensuses", path)
}

func TestPostAiConsensusEndpoint(t *testing.T) {
	aiConsensus := &PostAiConsensus{}
	endpoint, method, path := aiConsensus.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/v1/jobs/ai/consensuses", path)
}

func TestGetAiConsensusPayload(t *testing.T) {
	g := &GetAiConsensus{
		ID: "5a44671ab3957c2ab5c33326",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := g.Payload(endpoint, method, path)
	a.Equal(t, req.URL.Path, path+"/5a44671ab3957c2ab5c33326")
}

func TestGetListAiConsensusPayload(t *testing.T) {
	g := &GetAiConsensuses{
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

func TestPostAiConsensusPayload(t *testing.T) {
	p := &PostAiConsensus{
		Data:           "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
		PostbackURL:    "http://someUrl.url",
		PostbackMethod: "POST",
		CustomID:       "custom_id",
	}
	endpoint, method, path := p.Endpoint()
	req, _ := p.Payload(endpoint, method, path)
	a.NotNil(t, req)
}
