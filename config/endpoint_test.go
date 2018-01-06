package config

import (
	"os"
	"testing"

	a "github.com/stretchr/testify/assert"
)

func TestGetEndpoint(t *testing.T) {
	e := GetEndpoint()
	a.NotNil(t, e)
	a.Equal(t, API, e)
	os.Setenv("LOCAL_DEV", "true")
	e = GetEndpoint()
	a.NotNil(t, e)
	a.Equal(t, LocalAPI, e)
	os.Unsetenv("LOCAL_DEV")
}
