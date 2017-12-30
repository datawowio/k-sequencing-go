package actions

import "net/http"

// Action ...
type Action interface {
	Endpoint() (string, string)
	Payload(*http.Request)
}
