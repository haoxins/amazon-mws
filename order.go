package mws

import (
	"encoding/xml"
	"errors"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
	"github.com/spf13/cast"
)

// ListOrdersParams MWS ListOrders API params
type ListOrdersParams struct {
	LastUpdatedAfter  time.Time
	LastUpdatedBefore time.Time
	Statuses          []OrderStatus
}

// ListOrders List orders
func (seller *Seller) ListOrders(params ListOrdersParams) []OrderInfo {
	qs := seller.genListOrdersParams(params, "")

	result, err := seller.requestOrders(qs, false)
	tools.PanicError(err)

	var orders []OrderInfo

	orders = append(orders, result.OrderList.Orders...)

	for {
		if result.NextToken == "" {
			break
		}

		result = seller.ListOrdersByNextToken(result.NextToken)
		orders = append(orders, result.OrderList.Orders...)
	}

	return orders
}

// ListOrdersByNextToken ...
func (seller *Seller) ListOrdersByNextToken(nextToken string) ListOrdersResult {
	qs := seller.genListOrdersParams(ListOrdersParams{}, nextToken)

	result, err := seller.requestOrders(qs, true)
	tools.PanicError(err)

	return result
}

func (seller *Seller) genListOrdersParams(params ListOrdersParams, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	if nextToken != "" {
		v.Add("Action", "ListOrdersByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		mid := MarketplaceID[seller.Country]
		v.Add("MarketplaceId.Id.1", mid)

		for k, s := range params.Statuses {
			v.Add("OrderStatus.Status."+cast.ToString(k+1), string(s))
		}

		v.Add("Action", "ListOrders")
		v.Add("LastUpdatedAfter", params.LastUpdatedAfter.Format(time.RFC3339))
		v.Add("LastUpdatedBefore", params.LastUpdatedBefore.Format(time.RFC3339))
		v.Add("Version", "2013-09-01")
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestOrders(qs string, byNextToken bool) (ListOrdersResult, error) {
	body, err := seller.get(OrdersPath, qs)
	tools.PanicError(err)

	if byNextToken {
		data := ListOrdersByNextTokenResponse{}
		err := xml.Unmarshal(body, &data)

		if err != nil {
			return data.ListOrdersResult, errors.New(string(body))
		}

		return data.ListOrdersResult, nil
	}

	data := ListOrdersResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.ListOrdersResult, errors.New(string(body))
	}

	return data.ListOrdersResult, nil
}
