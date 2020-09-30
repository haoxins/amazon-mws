package mws

import (
	"encoding/xml"
	"errors"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
)

// ListFinancialEventsParams ListFinancialEvents API Params
type ListFinancialEventsParams struct {
	PostedAfter  time.Time
	PostedBefore time.Time
}

// ListFinancialEvents ...
func (seller *Seller) ListFinancialEvents(params ListFinancialEventsParams) []FinancialEvents {
	opts := seller.genListFinancialEventsParams(params, "")
	var financialEventsList []FinancialEvents

	result, err := seller.requestFinances(opts, false)
	tools.AssertError(err)

	financialEventsList = append(financialEventsList, result.FinancialEvents)

	for {
		if result.NextToken == "" {
			break
		}

		result = seller.listFinancialEventsByNextToken(result.NextToken)
		financialEventsList = append(financialEventsList, result.FinancialEvents)
	}

	return financialEventsList
}

func (seller *Seller) listFinancialEventsByNextToken(nextToken string) ListFinancialEventsResult {
	opts := seller.genListFinancialEventsParams(ListFinancialEventsParams{}, nextToken)

	result, err := seller.requestFinances(opts, true)
	tools.AssertError(err)

	return result
}

func (seller *Seller) genListFinancialEventsParams(params ListFinancialEventsParams, nextToken string) string {
	v := url.Values{}

	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("MaxResultsPerPage", "100")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("Version", "2015-05-01")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("SignatureVersion", "2")

	if nextToken != "" {
		v.Add("Action", "ListFinancialEventsByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "ListFinancialEvents")
		v.Add("PostedAfter", params.PostedAfter.Format(time.RFC3339))
		v.Add("PostedBefore", params.PostedBefore.Format(time.RFC3339))
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestFinances(qs string, byNextToken bool) (ListFinancialEventsResult, error) {
	// According to the document, this should be POST
	// But, only GET works
	body, err := seller.get(FinancesPath, qs)
	tools.AssertError(err)

	if byNextToken {
		data := ListFinancialEventsByNextTokenResponse{}
		err = xml.Unmarshal(body, &data)

		if err != nil {
			return data.ListFinancialEventsResult, errors.New(string(body))
		}

		return data.ListFinancialEventsResult, nil
	}

	data := ListFinancialEventsResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.ListFinancialEventsResult, errors.New(string(body))
	}

	return data.ListFinancialEventsResult, nil
}
