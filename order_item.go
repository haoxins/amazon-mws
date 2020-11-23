package mws

import (
	"encoding/xml"
	"errors"
	"net/url"

	"github.com/haoxins/tools/v2"
)

// ListOrderItems List order items
func (seller *Seller) ListOrderItems(amazonOrderID string) []OrderItem {
	qs := seller.genListOrderItemsParams(amazonOrderID, "")

	result, err := seller.requestOrderItems(qs, false)
	tools.PanicError(err)

	var items []OrderItem

	items = append(items, result.OrderItemList.OrderItems...)

	for {
		if result.NextToken == "" {
			break
		}

		result = seller.ListOrderItemsByNextToken(result.NextToken)
		items = append(items, result.OrderItemList.OrderItems...)
	}

	return items
}

// ListOrderItemsByNextToken ...
func (seller *Seller) ListOrderItemsByNextToken(nextToken string) ListOrderItemsResult {
	qs := seller.genListOrderItemsParams("", nextToken)

	result, err := seller.requestOrderItems(qs, true)
	tools.PanicError(err)

	return result
}

func (seller *Seller) genListOrderItemsParams(amazonOrderID string, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	if nextToken != "" {
		v.Add("Action", "ListOrderItemsByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "ListOrderItems")
		v.Add("AmazonOrderId", amazonOrderID)
	}

	v.Add("Version", "2013-09-01")

	s := v.Encode()

	return s
}

func (seller *Seller) requestOrderItems(qs string, byNextToken bool) (ListOrderItemsResult, error) {
	body, err := seller.get(OrdersPath, qs)
	tools.PanicError(err)

	if byNextToken {
		data := ListOrderItemsByNextTokenResponse{}
		err := xml.Unmarshal(body, &data)

		if err != nil {
			return data.ListOrderItemsResult, errors.New(string(body))
		}

		return data.ListOrderItemsResult, nil
	}

	data := ListOrderItemsResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.ListOrderItemsResult, errors.New(string(body))
	}

	return data.ListOrderItemsResult, nil
}
