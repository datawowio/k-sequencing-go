package main

import (
	"testing"

	"github.com/datawowio/k-sequencing-go/actions"
	a "github.com/stretchr/testify/assert"
)

const (
	TestProjectKey = "A7sRxaQKxo2hRQzNwkk5Qqx4"
)

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
