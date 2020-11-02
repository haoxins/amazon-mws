package mws

import (
	"strings"
)

func endpointToHost(endpoint string) string {
	return strings.Replace(endpoint, "https://", "", 1)
}
