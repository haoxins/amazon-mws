package mws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/haoxins/tools/v2"
)

// Seller The seller info
type Seller struct {
	Country   string
	SellerID  string
	AuthToken string
	AccessKey string
	SecretKey string
}

func (seller *Seller) addBasicParams(v *url.Values) {
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
}

func (seller *Seller) addSignature(method string, path string, payload string) string {
	escapedParams := strings.Replace(payload, ",", "%2C", -1)
	escapedParams = strings.Replace(escapedParams, ":", "%3A", -1)

	params := strings.Split(escapedParams, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	host := endpointToHost(Endpoint[seller.Country])
	toSignString := fmt.Sprintf(method+"\n%s\n%s\n%s", host, path, sortedParams)

	hasher := hmac.New(sha256.New, []byte(seller.SecretKey))
	_, err := hasher.Write([]byte(toSignString))
	tools.PanicError(err)

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	return sortedParams + "&Signature=" + hash
}

func (seller *Seller) get(path string, payload string) ([]byte, error) {
	payload = seller.addSignature("GET", path, payload)

	h := resty.New().SetTimeout(5 * time.Minute)
	res, err := h.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		Get(Endpoint[seller.Country] + path + "?" + payload)

	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

func (seller *Seller) post(path string, payload string) ([]byte, error) {
	payload = seller.addSignature("POST", path, payload)

	h := resty.New().SetTimeout(5 * time.Minute)
	res, err := h.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		SetBody([]byte(payload)).
		Post(Endpoint[seller.Country] + path)

	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}
