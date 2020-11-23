package mws

import (
	"encoding/xml"
	"errors"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
)

// ListInboundShipmentItemsParams ListInboundShipmentItems API params
type ListInboundShipmentItemsParams struct {
	ShipmentID        string
	LastUpdatedAfter  time.Time
	LastUpdatedBefore time.Time
}

// ListInboundShipmentItems ...
func (seller *Seller) ListInboundShipmentItems(params ListInboundShipmentItemsParams) []InboundShipmentItemMember {
	qs := seller.genListInboundShipmentItemsParams(params, "")

	var members []InboundShipmentItemMember

	result, err := seller.requestInboundShipmentItems(qs, false)
	tools.PanicError(err)

	members = append(members, result.ItemData.Members...)

	for {
		if result.NextToken == "" {
			break
		}

		result = seller.ListInboundShipmentItemsByNextToken(result.NextToken)
		members = append(members, result.ItemData.Members...)
	}

	return members
}

// ListInboundShipmentItemsByNextToken ...
func (seller *Seller) ListInboundShipmentItemsByNextToken(nextToken string) ListInboundShipmentItemsResult {
	qs := seller.genListInboundShipmentItemsParams(ListInboundShipmentItemsParams{}, nextToken)

	result, err := seller.requestInboundShipmentItems(qs, true)
	tools.PanicError(err)

	return result
}

func (seller *Seller) genListInboundShipmentItemsParams(params ListInboundShipmentItemsParams, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	v.Add("Version", "2010-10-01")

	if params.ShipmentID != "" {
		v.Add("ShipmentId", params.ShipmentID)

	} else {
		v.Add("LastUpdatedAfter", params.LastUpdatedAfter.Format(time.RFC3339))
		v.Add("LastUpdatedBefore", params.LastUpdatedBefore.Format(time.RFC3339))
	}

	if nextToken != "" {
		v.Add("Action", "ListInboundShipmentItemsByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "ListInboundShipmentItems")
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestInboundShipmentItems(qs string, byNextToken bool) (ListInboundShipmentItemsResult, error) {
	body, err := seller.get(InboundShipmentPath, qs)
	tools.PanicError(err)

	if byNextToken {
		data := ListInboundShipmentItemsByNextTokenResponse{}
		err = xml.Unmarshal(body, &data)

		if err != nil {
			return data.ListInboundShipmentItemsResult, errors.New(string(body))
		}

		return data.ListInboundShipmentItemsResult, nil
	}

	data := ListInboundShipmentItemsResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.ListInboundShipmentItemsResult, errors.New(string(body))
	}

	return data.ListInboundShipmentItemsResult, nil
}
