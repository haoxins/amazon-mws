package mws

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
)

// ListFinancialEventsParams ListFinancialEvents API Params
type ListFinancialEventsParams struct {
	PostedAfter  time.Time
	PostedBefore time.Time
}

// ListFinancialEvents ...
func (seller *Seller) ListFinancialEvents(params ListFinancialEventsParams) {
	opts, err := seller.genListFinancialEventsParams(params)
	tools.AssertError(err)

	body, err := seller.requestFinances(opts)
	tools.AssertError(err)

	// TODO
	fmt.Println(string(body))
}

func (seller *Seller) genListFinancialEventsParams(params ListFinancialEventsParams) (string, error) {
	v := url.Values{}

	v.Add("Action", "ListFinancialEvents")
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("PostedAfter", params.PostedAfter.Format(time.RFC3339))
	v.Add("PostedBefore", params.PostedBefore.Format(time.RFC3339))
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2015-05-01")

	s := v.Encode()

	return seller.AddSignature("POST", OrdersPath, s), nil
}

func (seller *Seller) requestFinances(params string) ([]byte, error) {
	return seller.post(FinancesPath, params)
}
