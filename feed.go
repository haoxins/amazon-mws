package mws

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
)

// GetFeedSubmissionListParams MWS GetFeedSubmissionList API params
type GetFeedSubmissionListParams struct {
	SubmittedFromDate time.Time
	SubmittedToDate   time.Time
}

// GetFeedSubmissionList ...
func (seller *Seller) GetFeedSubmissionList(params GetFeedSubmissionListParams) {
	opts, err := seller.genGetFeedSubmissionListParams(params)
	tools.AssertError(err)

	body, err := seller.requestFeed(opts)
	tools.AssertError(err)

	// TODO
	fmt.Println(string(body))
}

func (seller *Seller) genGetFeedSubmissionListParams(params GetFeedSubmissionListParams) (string, error) {
	v := url.Values{}

	mid := MarketplaceID[seller.Country]

	v.Add("Action", "GetFeedSubmissionList")
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("Marketplace", mid)
	v.Add("SubmittedFromDate", params.SubmittedFromDate.Format(time.RFC3339))
	v.Add("SubmittedToDate", params.SubmittedToDate.Format(time.RFC3339))
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2009-01-01")

	s := v.Encode()

	return seller.AddSignature("POST", OrdersPath, s), nil
}

func (seller *Seller) requestFeed(params string) ([]byte, error) {
	return seller.post(FeedsPath, params)
}
