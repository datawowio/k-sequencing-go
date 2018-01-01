package actions

import "net/http"

// Action ...
type Action interface {
	Endpoint() (string, string, string)
	Payload(string, string, string) (*http.Request, error)
}
