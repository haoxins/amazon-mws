package mws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/pkg4go/tools"
)

const (
	// ReportTypeSettlement ...
	ReportTypeSettlement = "_GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE_V2_"
	// ReportsPath ...
	ReportsPath = "/Reports/2009-01-01"
)

// AmazonSeller the amazon seller info
type AmazonSeller struct {
	SellerID  string
	AuthToken string
	Endpoint  string
	AccessKey string
	SecretKey string
}

// AddSignature add signature
func (seller *AmazonSeller) AddSignature(urlencode string) string {
	escapedParams := strings.Replace(urlencode, ",", "%2C", -1)
	escapedParams = strings.Replace(escapedParams, ":", "%3A", -1)

	params := strings.Split(escapedParams, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	host := endpointToHost(seller.MWSEndpoint)
	toSignString := fmt.Sprintf("GET\n%s\n%s\n%s", host, ReportsPath, sortedParams)

	hasher := hmac.New(sha256.New, []byte(seller.MwsSecretKey))
	_, err := hasher.Write([]byte(toSignString))
	tools.AssertError(err)

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	return sortedParams + "&Signature=" + hash
}
