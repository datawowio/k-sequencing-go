package kseq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/datawowio/k-sequencing-go/actions"
)

// Client is like a API Gateway for K-Sequencing. Client will let you call available
// K-Sequencing's APIs. It will be used with action structures from actions sub-package.
type Client struct {
	*http.Client
	ProjectKey string
}

// NewClient is used for creates and return a Client with given project key
func NewClient(projectKey string) (*Client, error) {
	if projectKey == "" {
		return nil, errors.New("invalid project key")
	}
	return &Client{&http.Client{}, projectKey}, nil
}

// Call performs supplied operations against K-Sequencing's API and unmarshal response into
// given action object.
//
// In successful case, result will contain 2 main objects, data and meta. (status code and
// message) Failed case, response will contain an error message.
func (c *Client) Call(result interface{}, act actions.Action) error {
	endpoint, method, path := act.Endpoint()
	req, err := act.Payload(endpoint, method, path)
	if err != nil {
		return err
	}

	if req.Header.Get("Authorization") == "" {
		c.setHeaders(req)
	}

	resp, e := c.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if e != nil {
		return e
	}

	buffer, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Println("Error while reading response body")
	}

	switch {
	case e != nil:
		return e
	case resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated:
		err := errors.New(resp.Status)
		return err
	}

	if result != nil {
		if e := json.Unmarshal(buffer, result); e != nil {
			return e
		}
	}

	return nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", c.ProjectKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}
