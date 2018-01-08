package actions

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

// GetPhotoTag represents a Get Image Photo Tag's Payload that required to request the action.
type GetPhotoTag struct {
	ID       string
	CustomID string
}

// GetPhotoTags represents a Get list of Image Photo Tags's Payload that required to request the
// action.
type GetPhotoTags struct {
	ID   string
	Page string
	Item string
}

// PostPhotoTag represents a Create list of Image Photo Tag's Payload that required to request
// the action.
type PostPhotoTag struct {
	Instruction    string
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
	StaffID        int64
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get Image
// Photo Tag API.
func (*GetPhotoTag) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/photo_tag"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// Image Photo Tag API.
func (*GetPhotoTags) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/images/photo_tags"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create Image
// Photo Tag API.
func (*PostPhotoTag) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "POST", "/api/images/photo_tags"
}

// Payload creates request's payload for Get Image Photo Tag API. Returns http.Request object
// which contains required query parameters.
func (g *GetPhotoTag) Payload(endpoint, method, path string) (*http.Request, error) {
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

// Payload creates request's payload for Get list Image Photo Tag API. Returns http.Request
// object which contains required query parameters.
func (g *GetPhotoTags) Payload(endpoint, method, path string) (*http.Request, error) {
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

// Payload creates request's payload for Create Image Photo Tag API. Returns http.Request
// object which contains required formData parameters.
func (p *PostPhotoTag) Payload(endpoint, method, path string) (*http.Request, error) {
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
