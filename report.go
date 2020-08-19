package mws

import (
	"encoding/xml"
	"net/url"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/pkg4go/tools"
	"github.com/thoas/go-funk"
)

// GetReportList Get amazon seller report list by report type
func (seller *AmazonSeller) GetReportList(startTime time.Time, endTime time.Time, nextToken string) *GetReportListResult {
	params, err := seller.GenAmazonGetReportListParams(ReportTypeSettlement, startTime, endTime, nextToken)
	tools.AssertError(err)

	raw, err := seller.RequestAmazonReport(params)
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
func (seller *AmazonSeller) GetAllReportIds(startTime time.Time, endTime time.Time) []string {
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

// GetReportByID Get amazon seller report by report id
func (seller *AmazonSeller) GetReportByID(reportID string) []SettlementReportRow {
	params, err := seller.GenAmazonGetReportParams(reportID)
	tools.AssertError(err)

	raw, err := seller.RequestAmazonReport(params)
	tools.AssertError(err)
	text := string(raw)

	return parseCSVReport(text)
}

// GenAmazonGetReportListParams gen amazon get report list params
func (seller *AmazonSeller) GenAmazonGetReportListParams(reportType string, startTime time.Time, endTime time.Time, nextToken string) (string, error) {
	v := url.Values{}

	if nextToken != "" {
		v.Add("Action", "GetReportListByNextToken")
		v.Add("NextToken", nextToken)
	} else {
		v.Add("Action", "GetReportList")
	}

	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.MwsAccessKey)
	v.Add("ReportTypeList.Type.1", reportType)
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2009-01-01")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("AvailableFromDate", startTime.Format(time.RFC3339))
	v.Add("AvailableToDate", endTime.Format(time.RFC3339))
	v.Add("MaxCount", "100")
	s := v.Encode()

	return seller.AddSignature(s), nil
}

// GenAmazonGetReportParams gen amazon get report params
func (seller *AmazonSeller) GenAmazonGetReportParams(reportID string) (string, error) {
	v := url.Values{}

	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.MwsAccessKey)
	v.Add("Action", "GetReport")
	v.Add("ReportId", reportID)
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2009-01-01")
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))

	urlencode := v.Encode()

	return seller.AddSignature(urlencode), nil
}

// RequestAmazonReport request amazon report
func (seller *AmazonSeller) RequestAmazonReport(params string) ([]byte, error) {
	h := resty.New()
	res, err := h.R().
		SetHeader("Content-Type", "x-www-form-urlencoded").
		Get(seller.MwsEndpoint + ReportsPath + "?" + params)

	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}
