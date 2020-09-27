package mws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"

	resty "github.com/go-resty/resty/v2"
	"github.com/pkg4go/tools"
)

// Seller The seller info
type Seller struct {
	Country   string
	SellerID  string
	AuthToken string
	AccessKey string
	SecretKey string
}

// AddSignature add signature
func (seller *Seller) AddSignature(path string, urlencode string) string {
	escapedParams := strings.Replace(urlencode, ",", "%2C", -1)
	escapedParams = strings.Replace(escapedParams, ":", "%3A", -1)

	params := strings.Split(escapedParams, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	host := endpointToHost(Endpoint[seller.Country])
	toSignString := fmt.Sprintf("GET\n%s\n%s\n%s", host, path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(seller.SecretKey))
	_, err := hasher.Write([]byte(toSignString))
	tools.AssertError(err)

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	return sortedParams + "&Signature=" + hash
}

// Request ...
func (seller *Seller) Request(path string, params string) ([]byte, error) {
	h := resty.New()
	res, err := h.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		Get(Endpoint[seller.Country] + path + "?" + params)

	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}
