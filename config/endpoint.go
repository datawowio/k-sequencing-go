package config

import (
	"os"
)

const (
	API      string = "https://k-sequencing.datawow.io"
	LocalAPI string = "http://localhost:3001"
)

func local() bool {
	_, l := os.LookupEnv("LOCAL_DEV")
	if l {
		return true
	}
	return false
}

func GetEndpoint() string {
	endpoint := API
	if local() {
		endpoint = LocalAPI
	}
	return endpoint
}
