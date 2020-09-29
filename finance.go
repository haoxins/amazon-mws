package mws

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
)

// ListFinancialEventsParams ListFinancialEvents API Params
type ListFinancialEventsParams struct {
	PostedAfter  time.Time
	PostedBefore time.Time
}

// ListFinancialEvents ...
func (seller *Seller) ListFinancialEvents(params ListFinancialEventsParams) error {
	opts := seller.genListFinancialEventsParams(params)

	body, err := seller.requestFinances(opts, false)
	tools.AssertError(err)

	fmt.Println(string(body))

	// data := ListFinancialEventsResponse{}
	// err = xml.Unmarshal(body, &data)
	// if err != nil {

	// 	return errors.New(string(body))
	// }

	// fmt.Printf("%+v", data)
	return nil
}

func (seller *Seller) genListFinancialEventsParams(params ListFinancialEventsParams) string {
	v := url.Values{}

	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("Action", "ListFinancialEvents")
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("SellerId", seller.SellerID)
	v.Add("MaxResultsPerPage", "100")
	v.Add("PostedAfter", params.PostedAfter.Format(time.RFC3339))
	v.Add("PostedBefore", params.PostedBefore.Format(time.RFC3339))
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("SignatureVersion", "2")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("Version", "2015-05-01")

	s := v.Encode()

	return s
}

func (seller *Seller) requestFinances(qs string, byNextToken bool) ([]byte, error) {
	// According to the document, this should be POST
	// But, only GET works
	return seller.get(FinancesPath, qs)
}
