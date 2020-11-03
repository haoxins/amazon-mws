package mws

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_endpointToHost(t *testing.T) {
	assert.Equal(t, endpointToHost("https://mws.amazonservices.jp"), "mws.amazonservices.jp")
}

func Test_fmtSku(t *testing.T) {
	assert.Equal(t, "LFA-4D3845-USCR3A-ABCD", fmtSku("LFA-4D3845- USCR3A+-ABCD"))
}
