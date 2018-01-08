package kseq

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/datawowio/k-sequencing-go/actions"
)

type Client struct {
	*http.Client
	ProjectKey string
}

func NewClient(projectKey string) (*Client, error) {
	if projectKey == "" {
		return nil, errors.New("invalid project key")
	}
	return &Client{&http.Client{}, projectKey}, nil
}

func (c *Client) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", c.ProjectKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
}

func (c *Client) Call(result interface{}, act actions.Action) error {
	endpoint, method, path := act.Endpoint()
	req, err := act.Payload(endpoint, method, path)
	if err != nil {
		return err
	}
	c.setHeaders(req)

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