package kseq

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/datawowio/k-sequencing-go/actions"
	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

const (
	TestProjectKey = "A7sRxaQKxo2hRQzNwkk5Qqx4"
	TestImageID    = "5a44671ab3957c2ab5c33326"
)

type mockPayload struct {
	ID       string
	CustomID string
}

func (*mockPayload) Endpoint() (string, string, string) {
	return config.LocalAPI, "method", "path"
}

func (m *mockPayload) Payload(endpoint, method, path string) (*http.Request, error) {
	return nil, errors.New("Mock error for payload testing")
}

func TestNewClient(t *testing.T) {
	c, err := NewClient(TestProjectKey)
	a.Nil(t, err)
	a.NotNil(t, c)
}

func TestNewClient_ErrorInvalidKey(t *testing.T) {
	c, err := NewClient("")
	a.NotNil(t, err)
	a.Nil(t, c)
}

func TestClient_Call(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	closedQuestion, getImage := &GetClosedQuestion{}, &actions.GetClosedQuestion{
		ID: TestImageID,
	}

	endpoint, _, path := getImage.Endpoint()
	mockResp := fmt.Sprintf(`{ "data": { "image": { "id": %q } } }`, TestImageID)
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(mockResp)

	e := c.Call(closedQuestion, getImage)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, closedQuestion)
	a.Equal(t, getImage.ID, closedQuestion.Data.Image.ID)
}

func TestClient_InvalidCall(t *testing.T) {
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	closedQuestion, getImage := &GetClosedQuestion{}, &actions.GetClosedQuestion{}

	endpoint, _, path := getImage.Endpoint()
	gock.New(endpoint).
		Get(path).
		Reply(401)

	e := c.Call(closedQuestion, getImage)
	a.EqualError(t, e, e.Error())
}

func TestClient_InvalidPayload(t *testing.T) {
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m := &mockPayload{}
	e := c.Call(nil, m)
	a.NotNil(t, e)
	a.EqualError(t, e, "Mock error for payload testing")
}
