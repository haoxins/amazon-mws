package mws

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/pkg4go/tools"
	"github.com/thoas/go-funk"
)

const (
	// ReportTypeSettlement ...
	ReportTypeSettlement = "_GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE_V2_"
	// ReportsPath ...
	ReportsPath = "/Reports/2009-01-01"
)

// AmazonSeller the amazon seller info
type AmazonSeller struct {
	SellerID     string
	AuthToken    string
	MwsHost      string
	MwsEndpoint  string
	MwsAccessKey string
	MwsSecretKey string
}

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

		log.Println("Oh yeah, has next")
		nextToken = result.NextToken
	}

	uids := funk.UniqString(ids)
	log.Printf("All ids' length is %d, and uniq ids' length is %d", len(ids), len(uids))

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

// AddSignature add signature
func (seller *AmazonSeller) AddSignature(urlencode string) string {
	escapedParams := strings.Replace(urlencode, ",", "%2C", -1)
	escapedParams = strings.Replace(escapedParams, ":", "%3A", -1)

	params := strings.Split(escapedParams, "&")
	sort.Strings(params)
	sortedParams := strings.Join(params, "&")

	toSignString := fmt.Sprintf("GET\n%s\n%s\n%s", seller.MwsHost, ReportsPath, sortedParams)

	hasher := hmac.New(sha256.New, []byte(seller.MwsSecretKey))
	_, err := hasher.Write([]byte(toSignString))
	tools.AssertError(err)

	hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	hash = url.QueryEscape(hash)

	return sortedParams + "&Signature=" + hash
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
