package actions

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

// GetMessage represents a Get Image Message's Payload that required to request the action.
type GetMessage struct {
	ID       string
	CustomID string
}

// GetMessages represents a Get list of Image Message's Payload that required to request
// the action.
type GetMessages struct {
	ID   string
	Page string
	Item string
}

// PostMessage represents a Create list of Image Message's Payload that required to
// request the action.
type PostMessage struct {
	Instruction    string
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
	StaffID        int64
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get Image Message API.
func (*GetMessage) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/message"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// Image Message API.
func (*GetMessages) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/messages"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create Image
// Message API.
func (*PostMessage) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "POST", "/api/images/messages"
}

// Payload creates request's payload for Get Image Message API. Returns http.Request object
// which contains required query parameters.
func (g *GetMessage) Payload(endpoint, method, path string) (*http.Request, error) {
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

// Payload creates request's payload for Get list of Image Message API. Returns http.Request
// object which contains required query parameters.
func (g *GetMessages) Payload(endpoint, method, path string) (*http.Request, error) {
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

// Payload creates request's payload for Create Image Message API. Returns http.Request
// object which contains required formData parameters.
func (p *PostMessage) Payload(endpoint, method, path string) (*http.Request, error) {
	values := url.Values{}
	if p.Data != "" {
		values.Set("data", p.Data)
	}
	if p.Instruction != "" {
		values.Set("instruction", p.Instruction)
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
