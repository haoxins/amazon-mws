package mws

import (
	"encoding/xml"
	"errors"
	"net/url"

	"github.com/haoxins/tools/v2"
	"github.com/spf13/cast"
)

// GetInboundGuidanceForSKUParams GetInboundGuidanceForSKU API params
type GetInboundGuidanceForSKUParams struct {
	SellerSKUList []string
}

// GetInboundGuidanceForSKU ...
func (seller *Seller) GetInboundGuidanceForSKU(params GetInboundGuidanceForSKUParams) ([]SKUInboundGuidance, []InvalidSKU) {
	qs := seller.genGetInboundGuidanceForSKUParams(params)

	result, err := seller.requestInboundGuidanceForSKU(qs)
	tools.PanicError(err)

	guidances := result.SKUInboundGuidanceList.SKUInboundGuidances
	invalidSKUs := result.InvalidSKUList.InvalidSKUs
	return guidances, invalidSKUs
}

func (seller *Seller) genGetInboundGuidanceForSKUParams(params GetInboundGuidanceForSKUParams) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	v.Add("Action", "GetInboundGuidanceForSKU")
	v.Add("Version", "2010-10-01")

	mid := MarketplaceID[seller.Country]
	v.Add("MarketplaceId", mid)

	skus := params.SellerSKUList
	if len(skus) > 50 {
		// TODO
		skus = skus[0:50]
	}

	for k, sku := range skus {
		s := cast.ToString(k + 1)
		if k+1 < 10 {
			// TODO: Hack
			// When sign
			// 1 10 2 3
			// But MWS wants
			// 1 2 3 ... 10
			// So
			// 01 02 03 ... 10
			s = "0" + s
		}
		v.Add("SellerSKUList.Id."+s, fmtSku(sku))
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestInboundGuidanceForSKU(qs string) (GetInboundGuidanceForSKUResult, error) {
	body, err := seller.get(InboundShipmentPath, qs)
	tools.PanicError(err)

	data := GetInboundGuidanceForSKUResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.GetInboundGuidanceForSKUResult, errors.New(string(body))
	}

	return data.GetInboundGuidanceForSKUResult, nil
}
