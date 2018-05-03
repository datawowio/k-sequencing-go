package actions

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

type TextClosedQuestionGetParams struct {
	ID       string
	CustomID string
}

type TextClosedQuestionParams struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
}

func (*TextClosedQuestionGetParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "GET", "/api/v1/text/text_closed_questions"
}

func (*TextClosedQuestionParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "POST", "/api/v1/text/text_closed_questions"
}

func (g *TextClosedQuestionGetParams) Payload(endpoint, method, path string) (*http.Request, error) {
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

func (p *TextClosedQuestionParams) Payload(endpoint, method, path string) (*http.Request, error) {
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

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequest(method, string(endpoint)+path, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}
