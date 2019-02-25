package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetPredictionTagEndpoint(t *testing.T) {
	g := &GetPrediction{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/prime/predictions", path)
}

func TestGetPredictionsPayload(t *testing.T) {
	g := &GetPredictions{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/prime/predictions", path)
}

func TestPostPrediction(t *testing.T) {
	p := &PostPrediction{}
	endpoint, method, path := p.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/prime/predictions", path)
}

func TestGetPredictionPayload(t *testing.T) {
	g := &GetPrediction{
		ID: "5a44671ab3957c2ab5c33326",
	}
	endpoint, method, path := g.Endpoint()
	req, err := g.Payload(endpoint, method, path)
	a.Nil(t, err)
	a.NotNil(t, req)
	a.Contains(t, req.URL.Path, g.ID)
}

func TestGetListPredictionsPayload(t *testing.T) {
	g := &GetPredictions{
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

func TestPostPredictionPayload(t *testing.T) {
	p := &PostPrediction{
		Data:           "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
		PostbackURL:    "http://someUrl.com",
		PostbackMethod: "POST",
		CustomID:       "custom_id",
	}
	endpoint, method, path := p.Endpoint()
	req, _ := p.Payload(endpoint, method, path)
	a.NotNil(t, req)
}
