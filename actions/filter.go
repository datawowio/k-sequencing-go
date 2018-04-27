package actions

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

type FilterParams struct {
	ProjectType string
	FilterSet   []string
	UseDefault  bool
}

func (*FilterParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "POST", "/api/v1/text/profanity/filters"
}

func (p *FilterParams) Payload(endpoint, method, path string) (*http.Request, error) {
	values := url.Values{}
	values.Set("use_default", "false")
	if p.ProjectType != "" {
		values.Set("project_type", p.ProjectType)
	}
	if len(p.FilterSet) > 0 {
		values.Set("filter_set", strings.Join(p.FilterSet[:], ","))
	}
	if p.UseDefault {
		values.Set("use_default", "true")
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, string(endpoint)+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
