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
//  imgData, get := &kseq.GetPhotoTag{}, &actions.GetPhotoTag{
//      ID:       "5a546e916e11571f570c1533",
//      CustomID: "3423401123",
//  }
//
//  if err := client.Call(imgData, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Photo Tag: %#v\n", imgData)
//
type GetPhotoTag struct {
	ID       string
	CustomID string
}

// Example:
//
//  list, get := &kseq.GetPhotoTags{}, &actions.GetPhotoTags{
//      ID:   "5a546e916e11571f570c1533",
//      Page: 1,
//      Item: 20,
//  }
//
//  if err := client.Call(list, get); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Photo Tags: %#v\n", list)
//  fmt.Printf("First element: %#v\n", list.Data.Images[0])
//
type GetPhotoTags struct {
	ID   string
	Page string
	Item string
}

// Example:
//
//  imgData, post := &kseq.PostPhotoTag{}, &actions.PostPhotoTag{
//		Instruction: "Instruction"
//		Data:        TestImageDataURL,
//  }
//
//  if err := client.Call(imgData, post); err != nil {
//      log.Fatal(err)
//  }
//
//  fmt.Printf("Image Photo Tag: %#v\n", imgData)
//
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
	return config.KiyoImageAPIURL, "GET", "/api/images/photo_tag"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Get list of
// Image Photo Tag API.
func (*GetPhotoTags) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "GET", "/api/images/photo_tags"
}

// Endpoint returns K Sequencing's request url, verb and endpoint for calling Create Image
// Photo Tag API.
func (*PostPhotoTag) Endpoint() (string, string, string) {
	return config.KiyoImageAPIURL, "POST", "/api/images/photo_tags"
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
