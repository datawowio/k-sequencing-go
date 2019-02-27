package actions

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestGetDocumentVerificationEndpoint(t *testing.T) {
	g := &DocumentVerificationGetParams{}
	url, method, endpoint := g.Endpoint()
	a.NotNil(t, url)
	a.Equal(t, config.KiyoImageAPIURL, url)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, endpoint)
	a.Equal(t, "/api/v1/images/document_verifications", endpoint)
}

func TestPostDocumentVerificationEndpoint(t *testing.T) {
	p := &DocumentVerificationParams{}
	url, method, endpoint := p.Endpoint()
	a.NotNil(t, url)
	a.Equal(t, config.KiyoImageAPIURL, url)
	a.NotNil(t, method)
	a.Equal(t, "POST", method)
	a.NotNil(t, endpoint)
	a.Equal(t, "/api/v1/images/document_verifications", endpoint)
}

func TestDocumentVerificationGetPayload(t *testing.T) {
	g := &DocumentVerificationGetParams{ID: "5a44671ab3957c2ab5c33326"}
	url, method, endpoint := g.Endpoint()
	req, _ := g.Payload(url, method, endpoint)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}

func TestDocumentVerificationPostPayload(t *testing.T) {
	p := &DocumentVerificationParams{
		Data: "https://sample-image.png",
		Info: map[string]map[string]string{
			"type": map[string]string{
				"value": "driver_license"},
			"dob": map[string]string{
				"value": "1991/11/28"},
		},
		PostbackURL:    "http://localhost:3000",
		PostbackMethod: "GET",
		CustomID:       "321",
	}
	url, method, endpoint := p.Endpoint()
	req, err := p.Payload(url, method, endpoint)
	a.Nil(t, err)
	a.NotNil(t, req)
}
