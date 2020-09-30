package mws

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
)

// ListOrdersParams MWS ListOrders API params
type ListOrdersParams struct {
	LastUpdatedAfter  time.Time
	LastUpdatedBefore time.Time
}

// ListOrders List orders
func (seller *Seller) ListOrders(params ListOrdersParams) {
	opts := seller.genListOrdersParams(params)

	body, err := seller.requestOrders(opts)
	tools.AssertError(err)

	data := ListOrdersResponse{}
	err = xml.Unmarshal(body, &data)
	if err != nil {
		// TODO
		log.Println(string(body))
		return
	}

	fmt.Printf("%+v", data)
}

func (seller *Seller) genListOrdersParams(params ListOrdersParams) string {
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

	return s
}

func (seller *Seller) requestOrders(params string) ([]byte, error) {
	return seller.get(OrdersPath, params)
}
