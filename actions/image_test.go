package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetImageEndpoint(t *testing.T) {
	g := &GetImage{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.KiyoImageAPIURL, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/projects/images", path)
}

func TestGetImagePayload(t *testing.T) {
	g := &GetImage{
		ID: "5a44671ab3957c2ab5c33326",
	}
	endpoint, method, path := g.Endpoint()
	req, err := g.Payload(endpoint, method, path)
	a.Nil(t, err)
	a.NotNil(t, req)
	a.Contains(t, req.URL.Path, g.ID)
}

func TestGetImagePayload_InvalidParams(t *testing.T) {
	g := &GetImage{}
	endpoint, method, path := g.Endpoint()
	req, err := g.Payload(endpoint, method, path)
	a.Nil(t, req)
	if a.Error(t, err) {
		a.Equal(t, "ID is required ", err.Error())
	}
}
