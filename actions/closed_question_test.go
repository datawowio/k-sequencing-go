package actions

import (
	"net/http"
	"testing"

	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

func TestEndpoint(t *testing.T) {
	g := &GetClosedQuestion{}
	endpoint, method, path := g.Endpoint()
	a.NotNil(t, endpoint)
	a.Equal(t, config.LocalAPI, endpoint)
	a.NotNil(t, method)
	a.Equal(t, "GET", method)
	a.NotNil(t, path)
	a.Equal(t, "/api/images/closed_question", path)
}

func TestPayload(t *testing.T) {
	g := &GetClosedQuestion{
		ID:       "5a44671ab3957c2ab5c33326",
		CustomID: "5022342340",
	}
	endpoint, method, path := g.Endpoint()
	req, _ := http.NewRequest(method, string(endpoint)+path, nil)
	g.Payload(req)
	queryValues := req.URL.Query()
	a.Equal(t, g.ID, queryValues.Get("id"))
	a.Equal(t, g.CustomID, queryValues.Get("custom_id"))
}
