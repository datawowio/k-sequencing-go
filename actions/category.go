package actions

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

type CategoryGetParams struct {
	ID       string
	CustomID string
}

type CategoryParams struct {
	Conversation   []Conversation
	Title          string
	AllowEmpty     bool
	PostbackURL    string
	PostbackMethod string
	CustomID       string
}

type Conversation struct {
	Name    string
	Message string
}

func (*CategoryGetParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "GET", "/api/v1/text/text_categories"
}

func (*CategoryParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "POST", "/api/v1/text/text_categories"
}

func (g *CategoryGetParams) Payload(endpoint, method, path string) (*http.Request, error) {
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

func (p *CategoryParams) Payload(endpoint, method, path string) (*http.Request, error) {
	var conversation []string

	for _, v := range p.Conversation {
		conversation = append(
			conversation,
			fmt.Sprintf(`{"name": %q, "message": %q}`, v.Name, v.Message),
		)
	}

	var jsonStr = []byte(fmt.Sprintf(`
		{
			"title": %q,
			"conversation": [%s],
			"allow_empty": %q,
			"postback_url": %q,
			"postback_method": %q,
			"custom_id": %q
		}
	`,
		p.Title,
		strings.Join(conversation, ", "),
		strconv.FormatBool(p.AllowEmpty),
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
