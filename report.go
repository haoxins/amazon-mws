package mws

import (
	"encoding/xml"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
	"github.com/thoas/go-funk"
)

// GetReportList Get seller report list by report type
func (seller *Seller) GetReportList(startTime time.Time, endTime time.Time, nextToken string) *GetReportListResult {
	params, err := seller.GenGetReportListParams(ReportTypeSettlement, startTime, endTime, nextToken)
	tools.AssertError(err)

	raw, err := seller.RequestReport(params)
	tools.AssertError(err)

	if nextToken == "" {
		data := GetReportListResponse{}
		err = xml.Unmarshal(raw, &data)
		tools.AssertError(err)
		return &data.GetReportListResult
	}

	data := GetReportListByNextTokenResponse{}
	err = xml.Unmarshal(raw, &data)
	tools.AssertError(err)
	return &data.GetReportListResult
}

// GetAllReportIds Get all report ids if has next
func (seller *Seller) GetAllReportIds(startTime time.Time, endTime time.Time) []string {
	nextToken := ""
	var ids []string

	for {
		result := seller.GetReportList(startTime, endTime, nextToken)
		for _, v := range result.ReportInfo {
			ids = append(ids, v.ReportID)
		}

		if !result.HasNext {
			break
		}

		nextToken = result.NextToken
	}

	uids := funk.UniqString(ids)

	return uids
}

// GetReportByID Get seller report by report id
func (seller *Seller) GetReportByID(reportID string) []SettlementReportRow {
	params, err := seller.GenGetReportParams(reportID)
	tools.AssertError(err)

	raw, err := seller.RequestReport(params)
	tools.AssertError(err)
	text := string(raw)

	return parseCSVReport(text)
}

// GenGetReportListParams gen get report list params
func (seller *Seller) GenGetReportListParams(reportType string, startTime time.Time, endTime time.Time, nextToken string) (string, error) {
	v := url.Values{}

	if nextToken != "" {
		v.Add("Action", "GetReportListByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "GetReportList")
	}

	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("ReportTypeList.Type.1", reportType)

	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("AvailableFromDate", startTime.Format(time.RFC3339))
	v.Add("AvailableToDate", endTime.Format(time.RFC3339))
	v.Add("MaxCount", "100")
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2009-01-01")

	s := v.Encode()

	return seller.AddSignature(s), nil
}

// GenGetReportParams gen get report params
func (seller *Seller) GenGetReportParams(reportID string) (string, error) {
	v := url.Values{}

	v.Add("Action", "GetReport")
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("ReportId", reportID)
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2009-01-01")

	urlencode := v.Encode()

	return seller.AddSignature(urlencode), nil
}

// RequestReport request report
func (seller *Seller) RequestReport(params string) ([]byte, error) {
	return seller.Request(ReportsPath, params)
}
