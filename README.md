# K Sequencing Go

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

See documentation on [godoc.org][1]

# Development

Set this environment variable for running code against `k-sequencing` locally.
```go
export LOCAL_DEV=true
```

[0]: http://datawow.io
[1]: https://godoc.org/github.com/datawowio/k-sequencing-go

# License

This project is licensed under [datawow.io][0]
