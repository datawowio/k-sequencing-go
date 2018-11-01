package actions

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/datawowio/k-sequencing-go/config"
)

// Example:
//
//  imgData, get := &kseq.GetImageCheck{}, &actions.GetImageCheck{
//      ID:       "5a546e916e11571f570c1533",
//      CustomID: "3423401123",
//  }
//
//  if err := client.Call(imgData, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Closed Question: %#v\n", imgData)
//
type GetImageCheck struct {
	ID string
}

// Example:
//
//  list, get := &kseq.GetImageChecks{}, &actions.GetImageChecks{
//      ID: "5a546e916e11571f570c1533",
//      Page: 1,
//      Item: 20,
//  }
//
//  if err := client.Call(list, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Closed Questions: %#v\n", list)
//  fmt.Printf("First element: %#v\n", list.Data.Images[0])
//
type GetImageChecks struct {
	ID   string
	Page string
	Item string
}

// Example:
//
//  imgData, post := &kseq.PostImageCheck{}, &actions.PostImageCheck{
//		Categories: []string{"foo,bar"},
//		Data:       TestImageDataURL,
//  }
//
//  if err := client.Call(imgData, post); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Closed Question: %#v\n", imgData)
//
type PostImageCheck struct {
	Data           string
	PostbackURL    string
	PostbackMethod string
	CustomID       string
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get Image
// Closed Question API.
func (*GetImageCheck) Endpoint() (string, string, string) {
	return config.KSeqAPIURL, "GET", "/api/images/check"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// Image Closed Question API.
func (*GetImageChecks) Endpoint() (string, string, string) {
	return config.KSeqAPIURL, "GET", "/api/images/check"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create Image
// Closed Question API.
func (*PostImageCheck) Endpoint() (string, string, string) {
	return config.KSeqAPIURL, "POST", "/api/images/check"
}

// Payload creates request's payload for Get Image Closed Question API. Returns http.Request
// object which contains required query parameters.
func (g *GetImageCheck) Payload(endpoint, method, path string) (*http.Request, error) {
	if g.ID == "" {
		return nil, errors.New("ID is required ")
	}

	req, err := http.NewRequest(method, string(endpoint)+path+"/"+g.ID, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Payload creates request's payload for Get list Image Closed Question API. Returns
// http.Request which contains required query parameters.
func (g *GetImageChecks) Payload(endpoint, method, path string) (*http.Request, error) {
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
func (p *PostImageCheck) Payload(endpoint, method, path string) (*http.Request, error) {
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
