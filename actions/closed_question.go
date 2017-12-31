package actions

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

// GetClosedQuestion ...
type GetClosedQuestion struct {
	ID       string
	CustomID string
}

type PostClosedQuestion struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
	StaffID        int64
}

// Endpoint ...
func (*GetClosedQuestion) Endpoint() (string, string, string) {
	return config.LocalAPI, "GET", "/api/images/closed_question"
}

func (*PostClosedQuestion) Endpoint() (string, string, string) {
	return config.LocalAPI, "POST", "/api/images/closed_questions"
}

// Payload ...
func (g *GetClosedQuestion) Payload() *http.Request {
	endpoint, method, path := g.Endpoint()
	req, _ := http.NewRequest(method, string(endpoint)+path, nil)
	q := req.URL.Query()
	if g.ID != "" {
		q.Add("id", g.ID)
	}
	if g.CustomID != "" {
		q.Add("custom_id", g.CustomID)
	}
	req.URL.RawQuery = q.Encode()
	return req
}

func (p *PostClosedQuestion) Payload() *http.Request {
	values := url.Values{}
	if p.Data != "" {
		values.Set("data", p.Data)
	}
	if p.PostbackURL != "" {
		values.Set("postback_url", p.PostbackURL)
	}
	if p.PostbackMethod != "" {
		values.Set("postback_method", p.PostbackMethod)
	}
	if p.CustomID != "" {
		values.Set("custom_id", p.CustomID)
	}
	values.Add("staff_id", strconv.FormatInt(p.StaffID, 10))

	body := strings.NewReader(values.Encode())
	endpoint, method, path := p.Endpoint()
	req, _ := http.NewRequest(method, string(endpoint)+path, body)
	return req
}
