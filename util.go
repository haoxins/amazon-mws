package mws

import (
	"strings"
)

func endpointToHost(endpoint string) string {
	return strings.Replace(endpoint, "https://", "", 1)
}

func fmtSku(s string) string {
	return strings.Replace(strings.Replace(s, "+", "", -1), " ", "", -1)
}
