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
	qs := seller.genListOrdersParams(params)

	body, err := seller.requestOrders(qs)
	tools.PanicError(err)

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

	seller.addBasicParams(&v)

	mid := MarketplaceID[seller.Country]
	v.Add("MarketplaceId.Id.1", mid)

	v.Add("Action", "ListOrders")
	v.Add("LastUpdatedAfter", params.LastUpdatedAfter.Format(time.RFC3339))
	v.Add("LastUpdatedBefore", params.LastUpdatedBefore.Format(time.RFC3339))
	v.Add("Version", "2013-09-01")

	s := v.Encode()

	return s
}

func (seller *Seller) requestOrders(params string) ([]byte, error) {
	return seller.get(OrdersPath, params)
}
