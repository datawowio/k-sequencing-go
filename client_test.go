package kseq

import (
	"errors"
	"net/http"
	"testing"

	"github.com/datawowio/k-sequencing-go/actions"
	"github.com/datawowio/k-sequencing-go/config"
	a "github.com/stretchr/testify/assert"
)

const (
	TestProjectKey = "A7sRxaQKxo2hRQzNwkk5Qqx4"
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
	c := NewClient(TestProjectKey)
	a.NotNil(t, c)
}

func TestClient_Call(t *testing.T) {
	c := NewClient(TestProjectKey)
	a.NotNil(t, c)

	closedQuestion, getImage := &GetClosedQuestion{}, &actions.GetClosedQuestion{
		ID: "5a44671ab3957c2ab5c33326",
	}

	e := c.Call(closedQuestion, getImage)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, closedQuestion)
	a.Equal(t, getImage.ID, closedQuestion.Data.Image.ID)
}

func TestClient_InvalidCall(t *testing.T) {
	c := NewClient(TestProjectKey)
	a.NotNil(t, c)

	closedQuestion, getImage := &GetClosedQuestion{}, &actions.GetClosedQuestion{}

	e := c.Call(closedQuestion, getImage)
	a.EqualError(t, e, "Error occurred from API side")
}

func TestClient_InvalidPayload(t *testing.T) {
	c := NewClient(TestProjectKey)
	a.NotNil(t, c)

	m := &mockPayload{}
	e := c.Call(nil, m)
	a.NotNil(t, e)
	a.EqualError(t, e, "Mock error for payload testing")
}
