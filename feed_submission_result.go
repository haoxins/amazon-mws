package mws

import (
	"fmt"
	"io/ioutil"
	"net/url"

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
