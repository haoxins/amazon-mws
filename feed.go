package mws

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
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

	result := seller.requestFeed(opts, false)

	var feeds []FeedSubmissionInfo

	feeds = append(feeds, result.FeedSubmissionInfo...)

	for {
		if !result.HasNext {
			break
		}

		result = seller.getFeedSubmissionListByNextToken(result.NextToken)
		feeds = append(feeds, result.FeedSubmissionInfo...)
	}

	return feeds
}

func (seller *Seller) getFeedSubmissionListByNextToken(nextToken string) GetFeedSubmissionListResult {
	opts := seller.genGetFeedSubmissionListParams(GetFeedSubmissionListParams{}, nextToken)

	result := seller.requestFeed(opts, true)

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
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
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

func (seller *Seller) requestFeed(qs string, byNextToken bool) GetFeedSubmissionListResult {
	// According to the document, this should be POST
	// But, only GET works
	body, err := seller.get(FeedsPath, qs)
	tools.AssertError(err)

	// TODO - Remove this
	fmt.Println(string(body))

	if byNextToken {
		data := GetFeedSubmissionListByNextTokenResponse{}
		err = xml.Unmarshal(body, &data)

		if err != nil {
			return GetFeedSubmissionListResult{}
		}

		return data.GetFeedSubmissionListResult
	}

	data := GetFeedSubmissionListResponse{}
	err = xml.Unmarshal(body, &data)

	if err != nil {
		return GetFeedSubmissionListResult{}
	}

	return data.GetFeedSubmissionListResult
}
