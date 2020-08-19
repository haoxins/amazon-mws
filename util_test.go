package mws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_endpointToHost(t *testing.T) {
	assert.Equal(t, endpointToHost("https://mws.amazonservices.jp"), "mws.amazonservices.jp")
}
