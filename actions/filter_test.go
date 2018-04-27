package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestPostFilterEndpoint(t *testing.T) {
	p := &FilterParams{}
	url, method, endpoint := p.Endpoint()
	a.NotNil(t, url)
	a.Equal(t, config.KiyoTextAPIURL, url)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, endpoint)
	a.Equal(t, "/api/v1/text/profanity/filters", endpoint)
}

func TestFilterPostPayload(t *testing.T) {
	p := &FilterParams{
		ProjectType: "text_facebook",
		FilterSet:   []string{"foo", "bar"},
		UseDefault:  true,
	}
	url, method, endpoint := p.Endpoint()
	req, err := p.Payload(url, method, endpoint)
	a.Nil(t, err)
	a.NotNil(t, req)
}
