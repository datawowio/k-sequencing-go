package config

import (
	"os"
)

const (
	// API represents the endpoint value of K Sequencing.
	API string = "https://k-sequencing.datawow.io"
	// LocalAPI represents the local endpoint of K Sequencing
	LocalAPI string = "http://localhost:3001"
)

func local() bool {
	_, l := os.LookupEnv("LOCAL_DEV")
	if l {
		return true
	}
	return false
}

// GetEndpoint returns the string of endpoint. Returns localAPI if you config environment
// variable `LOCAL_DEV` to true. Otherwise will be production site.
func GetEndpoint() string {
	endpoint := API
	if local() {
		endpoint = LocalAPI
	}
	return endpoint
}
