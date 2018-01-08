package actions

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

// GetClosedQuestion represents a Get Image Closed Question's Payload that required to
// request the action.
type GetClosedQuestion struct {
	ID       string
	CustomID string
}

// GetClosedQuestions represents a Get list of Image Closed Question's Payload that required
// to request the action.
type GetClosedQuestions struct {
	ID   string
	Page string
	Item string
}

// PostClosedQuestion represents a Create Image Closed Question's Payload that required to
// request the action.
type PostClosedQuestion struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
	StaffID        int64
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get Image
// Closed Question API.
func (*GetClosedQuestion) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/closed_question"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// Image Closed Question API.
func (*GetClosedQuestions) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/closed_questions"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create Image
// Closed Question API.
func (*PostClosedQuestion) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "POST", "/api/images/closed_questions"
}

// Payload creates request's payload for Get Image Closed Question API. Returns http.Request
// object which contains required query parameters.
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

// Payload creates request's payload for Get list Image Closed Question API. Returns
// http.Request which contains required query parameters.
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

// Payload creates request's payload for Create Image Close Question API. Returns
// http.Request which contains required query parameters.
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
