package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/datawowio/k-sequencing-go/actions"
)

const (
	ProjectKey    = "A7sRxaQKxo2hRQzNwkk5Qqx4"
	LocalEndpoint = "http://localhost:3001"
)

type Response struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

type Image struct {
	ID            string `json:"id"`
	Answer        string `json:"answer"`
	CreditCharged int    `json:"credit_charged"`
	CustomID      string `json:"custom_id"`
	Source        string `json:"data"`
	PostbackUrl   string `json:"postback_url"`
	ProcessedAt   string `json:"processed_at"`
	ProjectID     int    `json:"project_id"`
	Status        string `json:"status"`
}

type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Image Image `json:"image"`
}

type Client struct {
	*http.Client
	ProjectKey string
}

func NewClient(projectKey string) *Client {
	return &Client{&http.Client{}, projectKey}
}

func (c *Client) Call(result interface{}, act actions.Action) error {
	method, path := act.Endpoint()
	req, e := http.NewRequest(method, string(LocalEndpoint)+path, nil)
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
	data, getImage := &Response{}, &actions.GetClosedQuestion{
		ID: "5a44671ab3957c2ab5c33326",
	}

	if e := c.Call(data, getImage); e != nil {
		log.Fatal(e)
	}
	log.Println(data.Data.Image.Answer)
	log.Println(data.Meta.Code)
}
