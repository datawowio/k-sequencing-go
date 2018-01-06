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

type GetClosedQuestions struct {
	ID   string
	Page string
	Item string
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
	return config.GetEndpoint(), "GET", "/api/images/closed_question"
}

func (*GetClosedQuestions) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/closed_questions"
}

func (*PostClosedQuestion) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "POST", "/api/images/closed_questions"
}

// Payload ...
func (g *GetClosedQuestion) Payload(endpoint, method, path string) (*http.Request, error) {
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

func (g *GetClosedQuestions) Payload(endpoint, method, path string) (*http.Request, error) {
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

func (p *PostClosedQuestion) Payload(endpoint, method, path string) (*http.Request, error) {
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
	req, err := http.NewRequest(method, string(endpoint)+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
