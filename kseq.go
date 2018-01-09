/*
Package kseq provides GO binding for K Sequencing REST API.
Full REST API documentation is available at https://datawow.readme.io/v1.0/reference.

Usage

Create a client with kseq.NewClient, along with supply your project key. After that, use
client.Call with actions object from the https://www.github.com/datawowio/k-sequencing-go/actions
package to perform API calls. The first parameter to client.Call lets you supply a struct
to unmarshal the result.

Example

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
	log.Printf("%#v\n", closedQuestion)

We also provide Get "any type" Image endpoint API. You only supply project key and Image
ID (or Customer ID) for reference.

Example

	c, err := kseq.NewClient(ProjectKey)
	if err != nil {
		log.Fatal(err)
	}
	resp := make(map[string]interface{})

	getImage := &actions.GetImage{
		ID: "5a52fb556e11571f570c1530",
	}

	if err := c.Call(&resp, getImage); err != nil {
		log.Fatal(err)
	}

	data := resp["data"].(map[string]interface{})
	meta := resp["meta"].(map[string]interface{})
	image := data["image"].(map[string]interface{})
	log.Println("Image ID: " + image["id"])
	log.Println("Image Status: " + image["status"])
	log.Println("Response code: " + meta["code"])

*/
package kseq
