package actions

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

const (
	AiConsensusPath = "/api/v1/jobs/ai/consensuses"
)

// Example:
//
//  imgData, get := &kseq.GetAiConsensus{}, &actions.GetAiConsensus{
//      ID: "5a546e916e11571f570c1533",
//  }
//
//  if err := client.Call(imgData, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("AI Consensus: %#v\n", imgData)
//
type GetAiConsensus struct {
	ID string
}

// Example:
//
//  list, get := &kseq.GetAiConsensuses{}, &actions.GetAiConsensuses{
//      ID: "5a546e916e11571f570c1533",
//      Page: 1,
//      Item: 20,
//  }
//
//  if err := client.Call(list, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("AI Consensus: %#v\n", list)
//  fmt.Printf("First element: %#v\n", list.Data.Images[0])
//
type GetAiConsensuses struct {
	ID   string
	Page string
	Item string
}

// Example:
//
//  imgData, post := &kseq.PostAiConsensus{}, &actions.PostAiConsensus{
//		Data: TestImageDataURL,
//  }
//
//  if err := client.Call(imgData, post); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("AI Consensus: %#v\n", imgData)
//
type PostAiConsensus struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get AI
// Consensus API.
func (g *GetAiConsensus) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "GET", AiConsensusPath
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// AI Consensus API.
func (g *GetAiConsensuses) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "GET", AiConsensusPath
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create AI
// Consensus API.
func (p *PostAiConsensus) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "POST", AiConsensusPath
}

// Payload creates request's payload for Get AI Consensus API. Returns http.Request
// object which contains required query parameters.
func (g *GetAiConsensus) Payload(endpoint, method, path string) (*http.Request, error) {
	if g.ID == "" {
		return nil, errors.New("ID is required ")
	}

	req, err := http.NewRequest(method, string(endpoint)+path+"/"+g.ID, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Payload creates request's payload for Get list AI Consensus API. Returns
// http.Request which contains required query parameters.
func (g *GetAiConsensuses) Payload(endpoint, method, path string) (*http.Request, error) {
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

// Payload creates request's payload for Create AI Consensus API. Returns
// http.Request which contains required query parameters.
func (p *PostAiConsensus) Payload(endpoint, method, path string) (*http.Request, error) {
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
