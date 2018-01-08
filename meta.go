package kseq

// Meta represents the response status code and message after called APIs.
type Meta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
