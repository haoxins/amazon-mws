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
	opts := seller.genListFinancialEventsParams(params)

	body, err := seller.requestFinances(opts)
	tools.AssertError(err)

	// TODO
	fmt.Println(string(body))
}

func (seller *Seller) genListFinancialEventsParams(params ListFinancialEventsParams) string {
	v := url.Values{}

	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("Action", "ListFinancialEvents")
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("SellerId", seller.SellerID)
	v.Add("PostedAfter", params.PostedAfter.Format(time.RFC3339))
	v.Add("PostedBefore", params.PostedBefore.Format(time.RFC3339))
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("SignatureVersion", "2")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("Version", "2015-05-01")

	s := v.Encode()

	return s
}

func (seller *Seller) requestFinances(params string) ([]byte, error) {
	// According to the document, this should be POST
	// But, only GET works
	return seller.get(FinancesPath, params)
}
