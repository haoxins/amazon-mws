package mws

import (
	"encoding/xml"
	"errors"
	"net/url"
	"time"

	"github.com/haoxins/tools/v2"
)

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

	result, err := seller.requestFeed(opts, false)
	tools.AssertError(err)

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

	result, err := seller.requestFeed(opts, true)
	tools.AssertError(err)

	return result
}

func (seller *Seller) genGetFeedSubmissionListParams(params GetFeedSubmissionListParams, nextToken string) string {
	v := url.Values{}

	mid := MarketplaceID[seller.Country]

	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("Marketplace", mid)
	v.Add("MaxCount", "100")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("Version", "2009-01-01")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("SignatureVersion", "2")

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

func (seller *Seller) requestFeed(qs string, byNextToken bool) (GetFeedSubmissionListResult, error) {
	// According to the document, this should be POST
	// But, only GET works
	body, err := seller.get(FeedsPath, qs)
	tools.AssertError(err)

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
