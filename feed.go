package mws

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
)

// GetFeedSubmissionResult ...
func (seller *Seller) GetFeedSubmissionResult(feedSubmissionID string) {
	qs := seller.genGetFeedSubmissionResultParams(feedSubmissionID)

	body, err := seller.get(FeedsPath, qs)
	tools.PanicError(err)
	// WTF - TODO - xlsx, xml, txt
	fname := tools.Resolve(tools.Getwd(), fmt.Sprintf("feed_submission_result_%s.xlsx", feedSubmissionID))
	err = ioutil.WriteFile(fname, body, 0644)
	tools.PanicError(err)
}

func (seller *Seller) genGetFeedSubmissionResultParams(feedSubmissionID string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	v.Add("Action", "GetFeedSubmissionResult")
	v.Add("Version", "2009-01-01")
	v.Add("FeedSubmissionId", feedSubmissionID)

	s := v.Encode()

	return s
}

// GetFeedSubmissionListParams MWS GetFeedSubmissionList API params
type GetFeedSubmissionListParams struct {
	SubmittedFromDate time.Time
	SubmittedToDate   time.Time
}

// GetFeedSubmissionList ...
// https://docs.developer.amazonservices.com/en_UK/feeds/Feeds_GetFeedSubmissionList.html
// Maximum request quota Restore rate                 Hourly request quota
// 10 requests           One request every 45 seconds 80 requests per hour
func (seller *Seller) GetFeedSubmissionList(params GetFeedSubmissionListParams) []FeedSubmissionInfo {
	opts := seller.genGetFeedSubmissionListParams(params, "")
	var feeds []FeedSubmissionInfo

	result, err := seller.requestFeedSubmissionList(opts, false)
	tools.PanicError(err)

	feeds = append(feeds, result.FeedSubmissionInfos...)

	for {
		if !result.HasNext {
			break
		}

		result = seller.getFeedSubmissionListByNextToken(result.NextToken)
		feeds = append(feeds, result.FeedSubmissionInfos...)
	}

	return feeds
}

func (seller *Seller) getFeedSubmissionListByNextToken(nextToken string) GetFeedSubmissionListResult {
	opts := seller.genGetFeedSubmissionListParams(GetFeedSubmissionListParams{}, nextToken)

	result, err := seller.requestFeedSubmissionList(opts, true)
	tools.PanicError(err)

	return result
}

func (seller *Seller) genGetFeedSubmissionListParams(params GetFeedSubmissionListParams, nextToken string) string {
	v := url.Values{}

	seller.addBasicParams(&v)

	mid := MarketplaceID[seller.Country]
	v.Add("Marketplace", mid) // required ???

	v.Add("MaxCount", "100")
	v.Add("Version", "2009-01-01")

	if nextToken != "" {
		v.Add("Action", "GetFeedSubmissionListByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "GetFeedSubmissionList")
		v.Add("SubmittedFromDate", params.SubmittedFromDate.Format(time.RFC3339))
		v.Add("SubmittedToDate", params.SubmittedToDate.Format(time.RFC3339))
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestFeedSubmissionList(qs string, byNextToken bool) (GetFeedSubmissionListResult, error) {
	// According to the document, this should be POST
	// But, only GET works
	body, err := seller.get(FeedsPath, qs)
	tools.PanicError(err)

	if byNextToken {
		data := GetFeedSubmissionListByNextTokenResponse{}
		err = xml.Unmarshal(body, &data)

		if err != nil {
			return data.GetFeedSubmissionListResult, errors.New(string(body))
		}

		return data.GetFeedSubmissionListResult, nil
	}

	data := GetFeedSubmissionListResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return data.GetFeedSubmissionListResult, errors.New(string(body))
	}

	return data.GetFeedSubmissionListResult, nil
}
