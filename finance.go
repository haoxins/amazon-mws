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
	qs := seller.genListFinancialEventsParams(params, "")
	var financialEventsList []FinancialEvents

	result, err := seller.requestFinances(qs, false)
	tools.PanicError(err)

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
	qs := seller.genListFinancialEventsParams(ListFinancialEventsParams{}, nextToken)

	result, err := seller.requestFinances(qs, true)
	tools.PanicError(err)

	return result
}

func (seller *Seller) genListFinancialEventsParams(params ListFinancialEventsParams, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	v.Add("MaxResultsPerPage", "100")
	v.Add("Version", "2015-05-01")

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
	tools.PanicError(err)

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
