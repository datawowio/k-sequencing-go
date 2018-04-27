package actions

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

type ConversationGetParams struct {
	ID       string
	CustomID string
}

type ConversationParams struct {
	Conversation          []string
	CustomConversationIDs []string
	Answers               []string
	PostbackURL           string
	PostbackMethod        string
	CustomID              string
}

func (*ConversationGetParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "GET", "/api/v1/text/text_conversations"
}

func (*ConversationParams) Endpoint() (string, string, string) {
	return config.KiyoTextAPIURL, "POST", "/api/v1/text/text_conversations"
}

func (g *ConversationGetParams) Payload(endpoint, method, path string) (*http.Request, error) {
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

func (p *ConversationParams) Payload(endpoint, method, path string) (*http.Request, error) {
	values := url.Values{}
	if len(p.Conversation) > 0 {
		values.Set("conversation", strings.Join(p.Conversation[:], ","))
	}
	if len(p.CustomConversationIDs) > 0 {
		values.Set("custom_conversation_ids", strings.Join(p.CustomConversationIDs[:], ","))
	}
	if len(p.Answers) > 0 {
		values.Set("answers", strings.Join(p.Answers[:], ","))
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
