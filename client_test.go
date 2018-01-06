package kseq

import (
	"errors"
	"fmt"
	"io/ioutil"
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
	TestImageData  = "https://assets-cdn.github.com/images/modules/open_graph/github-mark.png"
)

func readFile(path string) []byte {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
	}
	return result
}

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

func TestClient_CallGetClosedQuestion(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	cq, getCQ := &GetClosedQuestion{}, &actions.GetClosedQuestion{
		ID: TestImageID,
	}

	endpoint, _, path := getCQ.Endpoint()
	mockResp := readFile("./testdata/closed_question.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(cq, getCQ)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, cq)
	a.Equal(t, getCQ.ID, cq.Data.Image.ID)
}

func TestClient_CallGetListClosedQuestion(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	cq, getCQ := &GetClosedQuestions{}, &actions.GetClosedQuestions{
		ID: TestImageID,
	}

	endpoint, _, path := getCQ.Endpoint()
	mockResp := readFile("./testdata/closed_questions.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(cq, getCQ)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, cq)
	a.Equal(t, getCQ.ID, cq.Data.Images[0].ID)
}

func TestClient_CallPostClosedQuestion(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	pq, postCQ := &PostClosedQuestion{}, &actions.PostClosedQuestion{
		Data: TestImageData,
	}

	endpoint, _, path := postCQ.Endpoint()
	mockResp := readFile("./testdata/post_closed_question.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(pq, postCQ)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, pq)
	a.Equal(t, postCQ.Data, pq.Data.Source)
}

func TestClient_CallGetChoice(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	ch, getC := &GetChoice{}, &actions.GetChoice{
		ID: TestImageID,
	}

	endpoint, _, path := getC.Endpoint()
	mockResp := readFile("./testdata/choice.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(ch, getC)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, ch)
	a.Equal(t, getC.ID, ch.Data.Image.ID)
}

func TestClient_CallGetListChoice(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	ch, getCh := &GetChoices{}, &actions.GetChoices{
		ID: TestImageID,
	}

	endpoint, _, path := getCh.Endpoint()
	mockResp := readFile("./testdata/choices.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(ch, getCh)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, ch)
	a.Equal(t, getCh.ID, ch.Data.Images[0].ID)
}

func TestClient_CallPostChoice(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	pch, postCh := &PostChoice{}, &actions.PostChoice{
		Instruction: "Image's instruction",
		Categories:  []string{"foo,bar"},
		Data:        TestImageData,
	}

	endpoint, _, path := postCh.Endpoint()
	mockResp := readFile("./testdata/post_choice.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(pch, postCh)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, pch)
	a.Equal(t, postCh.Data, pch.Data.Source)
	a.Equal(t, postCh.Instruction, pch.Data.Instruction)
	a.Equal(t, postCh.Categories, pch.Data.Categories)
}

func TestClient_CallGetMessage(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	msg, getMsg := &GetMessage{}, &actions.GetMessage{
		ID: TestImageID,
	}

	endpoint, _, path := getMsg.Endpoint()
	mockResp := readFile("./testdata/message.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(msg, getMsg)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, msg)
	a.Equal(t, getMsg.ID, msg.Data.Image.ID)
}

func TestClient_CallGetListMessage(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	msg, getMsg := &GetMessages{}, &actions.GetMessages{
		ID: TestImageID,
	}

	endpoint, _, path := getMsg.Endpoint()
	mockResp := readFile("./testdata/messages.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(msg, getMsg)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, msg)
	a.Equal(t, getMsg.ID, msg.Data.Images[0].ID)
}

func TestClient_CallPostMessage(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	msg, postMsg := &PostMessage{}, &actions.PostMessage{
		Instruction: "Instruction",
		Data:        TestImageData,
	}

	endpoint, _, path := postMsg.Endpoint()
	mockResp := readFile("./testdata/post_message.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(msg, postMsg)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, msg)
	a.Equal(t, postMsg.Instruction, msg.Data.Instruction)
	a.Equal(t, postMsg.Data, msg.Data.Source)
}

func TestClient_CallGetPhotoTag(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	p, getP := &GetPhotoTag{}, &actions.GetPhotoTag{
		ID: TestImageID,
	}

	endpoint, _, path := getP.Endpoint()
	mockResp := readFile("./testdata/photo_tag.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(p, getP)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, p)
	a.Equal(t, getP.ID, p.Data.Image.ID)
}

func TestClient_CallGetListPhotoTag(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	p, getP := &GetPhotoTags{}, &actions.GetPhotoTags{
		ID: TestImageID,
	}

	endpoint, _, path := getP.Endpoint()
	mockResp := readFile("./testdata/photo_tags.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(p, getP)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, p)
	a.Equal(t, getP.ID, p.Data.Images[0].ID)
}

func TestClient_CallPostPhotoTag(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	p, PostP := &PostPhotoTag{}, &actions.PostPhotoTag{
		Instruction: "Instruction",
		Data:        TestImageData,
	}

	endpoint, _, path := PostP.Endpoint()
	mockResp := readFile("./testdata/post_photo_tag.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(p, PostP)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, p)
	a.Equal(t, PostP.Data, p.Data.Source)
	a.Equal(t, PostP.Instruction, p.Data.Instruction)
}

func TestClient_CallGetPrediction(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	p, getP := &GetPrediction{}, &actions.GetPrediction{
		ID: TestImageID,
	}

	endpoint, _, path := getP.Endpoint()
	mockResp := readFile("./testdata/prediction.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(p, getP)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, p)
	a.Equal(t, getP.ID, p.Data.Image.ID)
}

func TestClient_CallGetListPrediction(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	p, getP := &GetPredictions{}, &actions.GetPredictions{
		ID: TestImageID,
	}

	endpoint, _, path := getP.Endpoint()
	mockResp := readFile("./testdata/predictions.json")
	gock.New(endpoint).
		Get(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(p, getP)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, p)
	a.Equal(t, getP.ID, p.Data.Images[0].ID)
}

func TestClient_CallPostPrediction(t *testing.T) {
	defer gock.Off()
	c, _ := NewClient(TestProjectKey)
	a.NotNil(t, c)

	p, postP := &PostPrediction{}, &actions.PostPrediction{
		Data: TestImageData,
	}

	endpoint, _, path := postP.Endpoint()
	mockResp := readFile("./testdata/post_prediction.json")
	gock.New(endpoint).
		Post(path).
		Reply(200).
		BodyString(string(mockResp))

	e := c.Call(p, postP)
	if !(a.NoError(t, e)) {
		return
	}

	a.NotNil(t, p)
	a.Equal(t, postP.Data, p.Data.Source)
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
