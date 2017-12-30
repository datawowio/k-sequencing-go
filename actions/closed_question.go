package actions

import "net/http"

// GetClosedQuestion ...
type GetClosedQuestion struct {
	ProjectKey string
	ID         string
	CustomID   string
}

// Endpoint ...
func (*GetClosedQuestion) Endpoint() (string, string) {
	return "GET", "/api/images/closed_question"
}

// Payload ...
func (g *GetClosedQuestion) Payload(req *http.Request) {
	q := req.URL.Query()
	if g.ID != "" {
		q.Add("id", g.ID)
	}
	if g.CustomID != "" {
		q.Add("custom_id", g.CustomID)
	}
	req.URL.RawQuery = q.Encode()
}
