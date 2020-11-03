package mws

import (
	"encoding/xml"
	"errors"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
	"github.com/spf13/cast"
)

// SupplyDetailResponse The ResponseGroup param for ListInventorySupply
type SupplyDetailResponse string

const (
	// SupplyDetailResponseBasic ...
	SupplyDetailResponseBasic SupplyDetailResponse = "Basic"
	// SupplyDetailResponseDetail ...
	SupplyDetailResponseDetail = "SupplyDetail"
)

// ListInventorySupplyParams The ListInventorySupply params
type ListInventorySupplyParams struct {
	SellerSkus         []string
	QueryStartDateTime time.Time
	ResponseGroup      SupplyDetailResponse
}

// ListInventorySupply ...
func (seller *Seller) ListInventorySupply(params ListInventorySupplyParams) []InventorySupplyMember {
	opts := seller.genListInventorySupplyParams(params, "")

	result, err := seller.requestListInventorySupply(opts, false)
	tools.PanicError(err)

	var members []InventorySupplyMember
	members = append(members, result.InventorySupplyList.Members...)

	for {
		if result.NextToken == "" {
			break
		}

		result = seller.listInventorySupplyByNextToken(result.NextToken)
		members = append(members, result.InventorySupplyList.Members...)

	}

	return members
}

func (seller *Seller) listInventorySupplyByNextToken(nextToken string) ListInventorySupplyResult {
	opts := seller.genListInventorySupplyParams(ListInventorySupplyParams{}, nextToken)
	seller.requestListInventorySupply(opts, true)
	result, err := seller.requestListInventorySupply(opts, true)
	tools.PanicError(err)

	return result
}

func (seller *Seller) genListInventorySupplyParams(params ListInventorySupplyParams, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	if nextToken != "" {
		v.Add("Action", "ListInventorySupplyByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "ListInventorySupply")
	}

	if seller.Country == "US" {
		mid := MarketplaceID[seller.Country]
		v.Add("MarketplaceId", mid)
	}

	v.Add("Version", "2010-10-01")
	v.Add("ResponseGroup", string(params.ResponseGroup))

	if len(params.SellerSkus) > 0 {
		for k, sku := range params.SellerSkus {
			v.Add("SellerSkus.member."+cast.ToString(k+1), sku)
		}
	} else {
		v.Add("QueryStartDateTime", params.QueryStartDateTime.Format(time.RFC3339))
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestListInventorySupply(qs string, byNextToken bool) (ListInventorySupplyResult, error) {
	body, err := seller.get(InventoryPath, qs)
	tools.PanicError(err)

	if byNextToken {
		data := ListInventorySupplyByNextTokenResponse{}
		err = xml.Unmarshal(body, &data)

		if err != nil {
			return data.ListInventorySupplyResult, errors.New(string(body))
		}

		return data.ListInventorySupplyResult, nil
	}

	data := ListInventorySupplyResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.ListInventorySupplyResult, errors.New(string(body))
	}

	return data.ListInventorySupplyResult, nil
}
