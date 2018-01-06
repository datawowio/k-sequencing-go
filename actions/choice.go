package actions

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

type GetChoice struct {
	ID       string
	CustomID string
}

type GetChoices struct {
	ID   string
	Page string
	Item string
}

type PostChoice struct {
	Instruction    string
	Categories     []string
	Data           string
	AllowEmpty     bool
	PostbackURL    string
	PostbackMethod string
	Multiple       bool
	CustomID       string
	StaffID        int64
}

func (*GetChoice) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/choice"
}

func (*GetChoices) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/choices"
}

func (*PostChoice) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "POST", "/api/images/choices"
}

func (g *GetChoice) Payload(endpoint, method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, string(endpoint)+path, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if g.ID != "" {
		q.Add("id", g.ID)
	}
	if g.CustomID != "" {
		q.Add("custom_id", g.CustomID)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (g *GetChoices) Payload(endpoint, method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, string(endpoint)+path, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if g.ID != "" {
		q.Add("id", g.ID)
	}
	if g.Page != "" {
		q.Add("page", g.Page)
	}
	if g.Item != "" {
		q.Add("per_page", g.Item)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (p *PostChoice) Payload(endpoint, method, path string) (*http.Request, error) {
	values := url.Values{}
	if p.Instruction != "" {
		values.Set("instruction", p.Instruction)
	}
	if len(p.Categories) > 0 {
		values.Set("categories", strings.Join(p.Categories[:], ","))
	}
	if p.Data != "" {
		values.Set("data", p.Data)
	}
	if p.AllowEmpty {
		values.Set("allow_empty", "true")
	}
	if p.PostbackURL != "" {
		values.Set("postback_url", p.PostbackURL)
	}
	if p.PostbackMethod != "" {
		values.Set("postback_method", p.PostbackMethod)
	}
	if p.Multiple {
		values.Set("multiple", "true")
	}
	if p.CustomID != "" {
		values.Set("custom_id", p.CustomID)
	}
	values.Add("staff_id", strconv.FormatInt(p.StaffID, 10))

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, string(endpoint)+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
