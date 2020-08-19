package mws

import (
	"net/url"
	"time"

	"github.com/pkg4go/tools"
)

// ListOrders List orders
func (seller *Seller) ListOrders(startTime time.Time, endTime time.Time) {
	params, err := seller.GenListOrdersParams(startTime, endTime)
	tools.AssertError(err)

	seller.RequestOrder(params)
}

// GenListOrdersParams Generate list orders params
func (seller *Seller) GenListOrdersParams(startTime time.Time, endTime time.Time) (string, error) {
	v := url.Values{}

	mid := MarketplaceID[seller.Country]

	v.Add("Action", "ListOrders")
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("MarketplaceId.Id.1", mid)
	v.Add("LastUpdatedAfter", startTime.Format(time.RFC3339))
	v.Add("LastUpdatedBefore", endTime.Format(time.RFC3339))
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2013-09-01")

	s := v.Encode()

	return seller.AddSignature(s), nil
}

// RequestOrder request order
func (seller *Seller) RequestOrder(params string) ([]byte, error) {
	return seller.Request(OrdersPath, params)
}
