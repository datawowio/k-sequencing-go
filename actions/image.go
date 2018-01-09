package actions

import (
	"errors"
	"net/http"

	"github.com/datawowio/k-sequencing-go/config"
)

// Example:
//
// 	resp := make(map[string]interface{})
//
// 	getImage := &actions.GetImage{
//		ID: "5a52fb556e11571f570c1530",
// 	}
//
// 	if err := client.Call(&resp, getImage); err != nil {
//		log.Fatal(err)
// 	}
//
// 	data := resp["data"].(map[string]interface{})
// 	meta := resp["meta"].(map[string]interface{})
// 	image := data["image"].(map[string]interface{})
// 	log.Println("Image ID: " + image["id"])
// 	log.Println("Image Status: " + image["status"])
// 	log.Println("Response code: " + meta["code"])
//
type GetImage struct {
	ID string
}

// Endpoint returns Get Image request url, verb and endpoint
func (*GetImage) Endpoint() (string, string, string) {
	return config.GetEndpoint(), "GET", "/api/projects/images"
}

// Payload creates request's payload for Get Image API. Returns http.Request object.
func (g *GetImage) Payload(endpoint, method, path string) (*http.Request, error) {
	if g.ID == "" {
		return nil, errors.New("ID is required ")
	}
	req, err := http.NewRequest(method, string(endpoint)+path+"/"+g.ID, nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
