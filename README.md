# K Sequencing Go

[![GoDoc](https://godoc.org/github.com/datawowio/k-sequencing-go?status.svg)][1]
[![CircleCI Build Status](https://circleci.com/gh/datawowio/k-sequencing-go.svg?style=shield)][2]

Install with:

```go
go get github.com/datawowio/k-sequencing-go
```

# Usage

Create client with `kseq.NewClient` and use actions object from the
`github.com/datawowio/k-sequencing-go/actions` package to perform API operations.

```go
package main

import (
	"log"

	kseq "github.com/datawowio/k-sequencing-go"
	"github.com/datawowio/k-sequencing-go/actions"
)

const (
	ProjectKey = "qv1y2snY1evgYMXNUF3XGSNC"
)

func main() {
	c, err := kseq.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}

	closedQuestion, getClosedQuestion := &kseq.GetClosedQuestion{}, &actions.GetClosedQuestion{
		ID: "5a52fb556e11571f570c1530",
	}

	if err := c.Call(closedQuestion, getClosedQuestion); err != nil {
		log.Fatal(err)
	}
	log.Printf("Get image closed question: %s, status: %s",
		closedQuestion.Data.Image.ID,
		closedQuestion.Data.Image.Status)
}
```

See more documentation details on [godoc.org][1]

# Operational Actions

We provide list of K-Sequencing service action that client can call from list below.

## Image Choice

### Get Image Choice by ID

###### Payload

```go
type GetChoice struct {
    // Either `ID` or `CustomID` must be provided
    ID       string
    CustomID string
}
```

###### Example

```go
choiceData, get := &kseq.GetChoice{}, &actions.GetChoice{
    ID:       "5a546e916e11571f570c1533",
    CustomID: "3423401123",
}

if err := client.Call(choiceData, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Choice: %#v\n", choiceData)
```

### Get list of Image Choice

###### Payload

```go
type GetChoices struct {
    ID   string // required, `ID` could be Image's ID or CustomID
    Page string // optional
    Item string // optional
}
```

###### Example

```go
list, get := &kseq.GetChoices{}, &actions.GetChoices{
    ID: "5a546e916e11571f570c1533",
}

if err := client.Call(list, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Choices: %#v\n", list)
fmt.Printf("First element: %#v\n", list.Data.Images[0])
```

### Create Image Choice

###### Payload

```go
type PostChoice struct {
    Instruction    string   // required
    Categories     []string // required
    Data           string   // required
    AllowEmpty     bool     // optional
    PostbackURL    string   // optional
    PostbackMethod string   // optional
    Multiple       bool     // optional
    CustomID       string   // optional
    StaffID        int64    // optional
}
```

###### Example

```go
choiceData, post := &kseq.PostChoice{}, &actions.PostChoice{
    Instruction: "Image's instruction",
    Categories:  []string{"foo,bar"},
    Data:        TestImageDataURL,
}

if err := client.Call(choiceData, post); err != nil {
  log.Fatal(err)
}

fmt.Printf("Image Choice: %#v\n", choiceData)
```

## Image Closed Question

### Get Image Closed Question by ID

###### Payload

```go
type GetClosedQuestion struct {
    // Either `ID` or `CustomID` must be provided
    ID       string 
    CustomID string
}
```

###### Example

```go
imgData, get := &kseq.GetClosedQuestion{}, &actions.GetClosedQuestion{
    ID:       "5a546e916e11571f570c1533",
    CustomID: "3423401123",
}

if err := client.Call(imgData, get); err != nil {
  log.Fatal(err)
}

fmt.Printf("Image Closed Question: %#v\n", imgData)
```

### Get list of Image Closed Question

###### Payload

```go
type GetClosedQuestions struct {
    ID   string // required, `ID` could be either Image's ID or CustomID
    Page string // optional
    Item string // optional
}
```

###### Example
```go
list, get := &kseq.GetClosedQuestions{}, &actions.GetClosedQuestions{
    ID: "5a546e916e11571f570c1533",
    Page: 1,
    Item: 20,
}

if err := client.Call(list, get); err != nil {
  log.Fatal(err)
}

fmt.Printf("Image Closed Questions: %#v\n", list)
fmt.Printf("First element: %#v\n", list.Data.Images[0])
```

### Create Image Closed Question

###### Payload

```go
type PostClosedQuestion struct {
    Data           string // required
    PostbackURL    string // optional
    PostbackMethod string // optional
    CustomID       string // optional
    StaffID        int64  // optional
}
```

## Image Message

### Get Image Message by ID

###### Payload

```go
type GetMessage struct {
    // Either `ID` or `CustomID` must be provided
    ID       string
    CustomID string
}
```

###### Example
```go
imgData, get := &kseq.GetMessage{}, &actions.GetMessage{
    ID:       "5a546e916e11571f570c1533",
    CustomID: "3423401123",
}

if err := client.Call(imgData, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Message: %#v\n", imgData)
```

### Get list of Image Message

###### Payload

```go
type GetMessages struct {
    ID   string // required, `ID` could be either Image's ID or CustomID
    Page string
    Item string
}
```

###### Example
```go
list, get := &kseq.GetMessages{}, &actions.GetMessages{
    ID:   "5a546e916e11571f570c1533",
    Page: 1,
    Item: 20,
}

if err := client.Call(list, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Messages: %#v\n", list)
fmt.Printf("First element: %#v\n", list.Data.Images[0])
```

### Create Image Message

###### Payload
```go
type PostMessage struct {
    Instruction    string // required
    Data           string // required
    PostbackURL    string // optional
    PostbackMethod string // optional
    CustomID       string // optional
    StaffID        int64  // optional
}
```

###### Example
```go
imgData, post := &kseq.PostMessage{}, &actions.PostMessage{
    Instruction: "Instruction"
    Categories:  []string{"foo,bar"},
    Data:        TestImageDataURL,
}

if err := client.Call(imgData, post); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Message: %#v\n", imgData)
```

## Image Photo Tags

### Get Image Photo Tags by ID

###### Payload
```go
type GetMessage struct {
    // Either `ID` or `CustomID` must be provided
    ID       string 
    CustomID string
}
```

###### Example
```go
imgData, get := &kseq.GetPhotoTag{}, &actions.GetPhotoTag{
    ID:       "5a546e916e11571f570c1533",
    CustomID: "3423401123",
}

if err := client.Call(imgData, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Photo Tag: %#v\n", imgData)
```

### Get list of Image Photo Tags

###### Payload
```go
type GetMessages struct {
    ID   string // required, `ID` could be either Image's ID or CustomID
    Page string
    Item string
}
```

###### Example
```go
list, get := &kseq.GetPhotoTags{}, &actions.GetPhotoTags{
    ID:   "5a546e916e11571f570c1533",
    Page: 1,
    Item: 20,
}

if err := client.Call(list, get); err != nil {
  log.Fatal(err)
}

fmt.Printf("Image Photo Tags: %#v\n", list)
fmt.Printf("First element: %#v\n", list.Data.Images[0])
```

### Create Image Photo Tags

###### Payload
```go
type PostMessage struct {
	  Instruction    string // required
	  Data           string // required
	  PostbackURL    string // optional
	  PostbackMethod string // optional
	  CustomID       string // optional
	  StaffID        int64  // optional
}
```

###### Example
```go
imgData, post := &kseq.PostPhotoTag{}, &actions.PostPhotoTag{
    Instruction: "Instruction"
    Data:        TestImageDataURL,
}

if err := client.Call(imgData, post); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Photo Tag: %#v\n", imgData)
```

## Image Prediction

### Get Image Prediction by ID

###### Payload
```go
type GetPrediction struct {
	  ID string // required
}
```

###### Example
```go
imgData, get := &kseq.GetPrediction{}, &actions.GetPrediction{
    ID: "5a546e916e11571f570c1533",
}

if err := client.Call(imgData, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Prediction: %#v\n", imgData)
```

### Get list of Image Prediction

###### Payload
```go
type GetPredictions struct {
    ID   string // required, `ID` could be either Image's ID or CustomID
    Page string
    Item string
}
```

###### Example
```go
list, get := &kseq.GetPredictions{}, &actions.GetPredictions{
    ID:   "5a546e916e11571f570c1533",
    Page: 1,
    Item: 20,
}

if err := client.Call(list, get); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Predictions: %#v\n", list)
fmt.Printf("First element: %#v\n", list.Data.Images[0])
```

### Create Image Prediction

###### Payload
```go
type PostPrediction struct {
	Data           string // required
	PostbackURL    string // optional
	PostbackMethod string // optional
	CustomID       string // optional
}
```

###### Example
```go
imgData, post := &kseq.PostPrediction{}, &actions.PostPrediction{
    Data: TestImageDataURL,
}

if err := client.Call(imgData, post); err != nil {
    log.Fatal(err)
}

fmt.Printf("Image Prediction: %#v\n", imgData)
```

In case you don't want to specify type of struct when you call get list of image, you can 
call via `GetImage` which will filter the result from `ProjectKey` you supply to `Client`.

### Get list of Images from project

###### Payload
```go
type GetImage struct {
    ID string // required, `ID` could be Image's ID or CustomID you created.
}
```

###### Example
```go
resp := make(map[string]interface{})

getImage := &actions.GetImage{
    ID: "5a52fb556e11571f570c1530",
}

if err := client.Call(&resp, getImage); err != nil {
    log.Fatal(err)
}

data := resp["data"].(map[string]interface{})
meta := resp["meta"].(map[string]interface{})
image := data["image"].(map[string]interface{})
log.Println("Image ID: " + image["id"])
log.Println("Image Status: " + image["status"])
log.Println("Response code: " + meta["code"])
```

[0]: http://datawow.io
[1]: https://godoc.org/github.com/datawowio/k-sequencing-go
[2]: https://circleci.com/gh/datawowio/k-sequencing-go

# License

This project is licensed under [datawow.io][0]
