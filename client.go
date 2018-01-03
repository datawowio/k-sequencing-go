package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	// fmt.Println("resp:", resp.StatusCode, string(buffer))

	switch {
	case e != nil:
		return e
	case resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated:
		err := errors.New("Error occurred from API side")
		return err
	}

	if result != nil {
		if e := json.Unmarshal(buffer, result); e != nil {
			return e
		}
	}

	return nil
}

func main() {
	c := NewClient(ProjectKey)
	closedQuestion, getImage := &GetClosedQuestion{}, &actions.GetClosedQuestion{
		ID: "5a44671ab3957c2ab5c33326",
	}

	if e := c.Call(closedQuestion, getImage); e != nil {
		log.Fatal(e)
	}
	log.Println(closedQuestion.Data.Image.Answer)
	log.Println(closedQuestion.Meta.Code)

	pClosedQuestion, postImage := &PostClosedQuestion{}, &actions.PostClosedQuestion{
		Data: "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
	}
	if e := c.Call(pClosedQuestion, postImage); e != nil {
		log.Fatal(e)
	}
	log.Println(pClosedQuestion.Data)
	log.Println(pClosedQuestion.Data.ID)

	pChoice, postChoice := &PostChoice{}, &actions.PostChoice{
		Instruction: "Instruction",
		Categories:  []string{"Gnoon"},
		Data:        "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
	}
	if e := c.Call(pChoice, postChoice); e != nil {
		log.Fatal(e)
	}
	fmt.Println(pChoice.Data.Source)

	pm, postMsg := &PostMessage{}, &actions.PostMessage{
		Instruction: "Some instruction",
		Data:        "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png",
	}
	if e := c.Call(pm, postMsg); e != nil {
		log.Fatal(e)
	}
	log.Println(pm.Data)

	m, getMsg := &GetMessages{}, &actions.GetMessages{
		ID: "5a4c982402bc7141ad009dbb",
	}

	if e := c.Call(m, getMsg); e != nil {
		log.Fatal(e)
	}
	log.Println(m.Data.Images[0].ID)
	log.Println(m.Meta.Code)
}
