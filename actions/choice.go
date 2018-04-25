package actions

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

// Example:
//
//  choiceData, get := &kseq.GetChoice{}, &actions.GetChoice{
//      ID:       "5a546e916e11571f570c1533",
//      CustomID: "3423401123",
//  }
//
//  if err := client.Call(choiceData, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Choice: %#v\n", choiceData)
//
type GetChoice struct {
	ID       string
	CustomID string
}

// Example:
//
//  list, get := &kseq.GetChoices{}, &actions.GetChoices{
//      ID: "5a546e916e11571f570c1533",
//  }
//
//  if err := client.Call(list, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Choices: %#v\n", list)
//  fmt.Printf("First element: %#v\n", list.Data.Images[0])
//
type GetChoices struct {
	ID   string
	Page string
	Item string
}

// Example:
//
//  choiceData, post := &kseq.PostChoice{}, &actions.PostChoice{
//		Instruction: "Image's instruction",
//		Categories:  []string{"foo,bar"},
//		Data:        TestImageDataURL,
//  }
//
//  if err := client.Call(choiceData, post); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Choice: %#v\n", choiceData)
//
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

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get Image
// Choice API.
func (*GetChoice) Endpoint() (string, string, string) {
	return config.KSeqAPIURL, "GET", "/api/images/choice"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// Image Choice API.
func (*GetChoices) Endpoint() (string, string, string) {
	return config.KSeqAPIURL, "GET", "/api/images/choices"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create Image
// Choice API.
func (*PostChoice) Endpoint() (string, string, string) {
	return config.KSeqAPIURL, "POST", "/api/images/choices"
}

// Payload creates request's payload for Get Image Choice API. Returns http.Request object
// which contains required query parameters.
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

// Payload creates request's payload for Get list Image Choice API. Returns http.Request
// object which contains required query parameters.
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

// Payload creates request's payload for Create Image Choice API. Returns http.Request
// object which contains required formData parameters.
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
