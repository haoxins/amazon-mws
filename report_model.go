package mws

import "encoding/xml"

// GetReportListResponse Get report list api response
type GetReportListResponse struct {
	XMLName             xml.Name            `xml:"GetReportListResponse"`
	Xmlns               string              `xml:"xmlns,attr"`
	GetReportListResult GetReportListResult `xml:"GetReportListResult"`
	ResponseMetadata    ResponseMetadata    `xml:"ResponseMetadata"`
}

// GetReportListByNextTokenResponse The Get report list by next token response
type GetReportListByNextTokenResponse struct {
	XMLName             xml.Name            `xml:"GetReportListByNextTokenResponse"`
	Xmlns               string              `xml:"xmlns,attr"`
	GetReportListResult GetReportListResult `xml:"GetReportListResult"`
	ResponseMetadata    ResponseMetadata    `xml:"ResponseMetadata"`
}

// GetReportListResult The Get report list result
type GetReportListResult struct {
	NextToken   string                `xml:"NextToken"`
	HasNext     bool                  `xml:"HasNext"`
	ReportInfos []SingleReportInfoRow `xml:"ReportInfo"`
}

// SingleReportInfoRow The report info row
type SingleReportInfoRow struct {
	ReportType      string `xml:"ReportType"`
	Acknowledged    string `xml:"Acknowledged"`
	ReportID        string `xml:"ReportId"`
	ReportRequestID string `xml:"ReportRequestId"`
	AvailableDate   string `xml:"AvailableDate"`
}

// SettlementReportRow The settlement report row
type SettlementReportRow struct {
	SettlementID             string
	SettlementStartDate      string
	SettlementEndDate        string
	DepositDate              string
	TotalAmount              float64
	Currency                 string
	TransactionType          string
	OrderID                  string
	MerchantOrderID          string
	AdjustmentID             string
	ShipmentID               string
	MarketplaceName          string
	AmountType               string
	AmountDescription        string
	Amount                   float64
	FulfillmentID            string
	PostedDate               string
	PostedDateTime           string
	OrderItemCode            string
	MerchantOrderItemID      string
	MerchantAdjustmentItemID string
	Sku                      string
	QuantityPurchased        string
	PromotionID              string
}

// InventoryReportRow The inventory report row
type InventoryReportRow struct {
	Sku      string
	Asin     string
	Price    float64
	Quantity float64
}
