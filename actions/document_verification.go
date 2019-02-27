package actions

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

const (
	documentVerificationAPIEndpoint = "/api/v1/images/document_verifications"
)

type DocumentVerificationGetParams struct {
	ID       string
	CustomID string
}

type DocumentVerificationParams struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
	Info           map[string]map[string]string
}

func (*DocumentVerificationGetParams) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "GET", documentVerificationAPIEndpoint
}

func (*DocumentVerificationParams) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "POST", documentVerificationAPIEndpoint
}

func (g *DocumentVerificationGetParams) Payload(endpoint, method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, string(endpoint)+path, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if g.ID != "" {
		q.Add("id", g.ID)
	} else if g.CustomID != "" {
		q.Add("id", g.CustomID)
	}

	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (p *DocumentVerificationParams) Payload(endpoint, method, path string) (*http.Request, error) {
	var jsonStr = []byte(fmt.Sprintf(`
		{
			"data": %q,
			"info": {%s},
			"postback_url": %q,
			"postback_method": %q,
			"custom_id": %q
		}
	`,
		p.Data,
		strings.Join(createKeyValuePairs(p.Info), ", "),
		p.PostbackURL,
		p.PostbackMethod,
		p.CustomID))

	req, err := http.NewRequest(method, string(endpoint)+path, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func createKeyValuePairs(m map[string]map[string]string) []string {
	var kv []string
	for key, value := range m {
		for _, v := range value {
			kv = append(
				kv,
				fmt.Sprintf(`%q: {"value": %q}`, key, v),
			)
		}
	}
	return kv
}
