package actions

import (
	"errors"
	"net/http"

	"github.com/datawowio/k-sequencing-go/config"
)

// GetImage represents Get Image's payload that required for requesting the action
type GetImage struct {
	ID string
}

// Endpoint returns Get Image request url, verb and endpoint
func (*GetImage) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/projects/images"
}

// Payload creates request's payload for Get Image API. Returns http.Request object.
func (g *GetImage) Payload(endpoint, method, path string) (*http.Request, error) {
	if g.ID == "" {
		return nil, errors.New("ID is required ")
	}
	req, err := http.NewRequest(method, string(endpoint)+path+"/"+g.ID, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
