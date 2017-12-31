package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/datawowio/k-sequencing-go/actions"
)

const (
	ProjectKey = "A7sRxaQKxo2hRQzNwkk5Qqx4"
)

type Client struct {
	*http.Client
	ProjectKey string
}

func NewClient(projectKey string) *Client {
	return &Client{&http.Client{}, projectKey}
}

func (c *Client) Call(result interface{}, act actions.Action) error {
	endpoint, method, path := act.Endpoint()
	req, e := http.NewRequest(method, string(endpoint)+path, nil)
	if e != nil {
		return e
	}
	act.Payload(req)
	req.Header.Set("Authorization", c.ProjectKey)
	req.Header.Set("Content-Type", "application/json")

	resp, e := c.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	switch {
	case e != nil:
		return e
	case resp.StatusCode != 200:
		err := errors.New("Error occurred from API side")
		return err
	}

	if result != nil {
		if json.NewDecoder(resp.Body).Decode(result); e != nil {
			return e
		}
	}

	return nil
}

func main() {
	c := NewClient(ProjectKey)
	closedQuestion, getImage := &ClosedQuestion{}, &actions.GetClosedQuestion{
		ID: "5a44671ab3957c2ab5c33326",
	}

	if e := c.Call(closedQuestion, getImage); e != nil {
		log.Fatal(e)
	}
	log.Println(closedQuestion.Data.Image.Answer)
	log.Println(closedQuestion.Meta.Code)
}
