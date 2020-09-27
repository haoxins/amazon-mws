package mws

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
)

// ListOrdersParams MWS ListOrders API params
type ListOrdersParams struct {
	LastUpdatedAfter time.Time
	LastUpdatedBefore time.Time
}

// ListOrders List orders
func (seller *Seller) ListOrders(params ListOrdersParams) {
	opts, err := seller.genListOrdersParams(params)
	tools.AssertError(err)

	body, err := seller.requestOrders(opts)
	tools.AssertError(err)

	// TODO
	fmt.Println(string(body))
}

func (seller *Seller) genListOrdersParams(params ListOrdersParams) (string, error) {
	v := url.Values{}

	mid := MarketplaceID[seller.Country]

	v.Add("Action", "ListOrders")
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("MarketplaceId.Id.1", mid)
	v.Add("LastUpdatedAfter", params.LastUpdatedAfter.Format(time.RFC3339))
	v.Add("LastUpdatedBefore", params.LastUpdatedBefore.Format(time.RFC3339))
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2013-09-01")

	s := v.Encode()

	return seller.AddSignature(OrdersPath, s), nil
}

func (seller *Seller) requestOrders(params string) ([]byte, error) {
	return seller.Request(OrdersPath, params)
}
