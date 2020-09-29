package mws

import "encoding/xml"

// ListOrdersResponse ...
type ListOrdersResponse struct {
	XMLName          xml.Name         `xml:"ListOrdersResponse"`
	Text             string           `xml:",chardata"`
	Xmlns            string           `xml:"xmlns,attr"`
	ListOrdersResult ListOrdersResult `xml:"ListOrdersResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// ListOrdersResult ...
type ListOrdersResult struct {
	Text      string `xml:",chardata"`
	NextToken string `xml:"NextToken"`
	Orders    struct {
		Text  string      `xml:",chardata"`
		Order []OrderInfo `xml:"Order"`
	} `xml:"Orders"`
	LastUpdatedBefore string `xml:"LastUpdatedBefore"`
}

// OrderInfo ...
type OrderInfo struct {
	Text                   string `xml:",chardata"`
	LatestShipDate         string `xml:"LatestShipDate"`
	OrderType              string `xml:"OrderType"`
	PurchaseDate           string `xml:"PurchaseDate"`
	AmazonOrderID          string `xml:"AmazonOrderId"`
	BuyerEmail             string `xml:"BuyerEmail"`
	LastUpdateDate         string `xml:"LastUpdateDate"`
	IsReplacementOrder     string `xml:"IsReplacementOrder"`
	NumberOfItemsShipped   string `xml:"NumberOfItemsShipped"`
	ShipServiceLevel       string `xml:"ShipServiceLevel"`
	OrderStatus            string `xml:"OrderStatus"`
	SalesChannel           string `xml:"SalesChannel"`
	IsBusinessOrder        string `xml:"IsBusinessOrder"`
	NumberOfItemsUnshipped string `xml:"NumberOfItemsUnshipped"`
	PaymentMethodDetails   struct {
		Text                string `xml:",chardata"`
		PaymentMethodDetail string `xml:"PaymentMethodDetail"`
	} `xml:"PaymentMethodDetails"`
	IsGlobalExpressEnabled string `xml:"IsGlobalExpressEnabled"`
	IsSoldByAB             string `xml:"IsSoldByAB"`
	IsPremiumOrder         string `xml:"IsPremiumOrder"`
	OrderTotal             struct {
		Text         string `xml:",chardata"`
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"OrderTotal"`
	EarliestShipDate   string `xml:"EarliestShipDate"`
	MarketplaceID      string `xml:"MarketplaceId"`
	FulfillmentChannel string `xml:"FulfillmentChannel"`
	PaymentMethod      string `xml:"PaymentMethod"`
	ShippingAddress    struct {
		Text                         string `xml:",chardata"`
		City                         string `xml:"City"`
		PostalCode                   string `xml:"PostalCode"`
		IsAddressSharingConfidential string `xml:"isAddressSharingConfidential"`
		StateOrRegion                string `xml:"StateOrRegion"`
		CountryCode                  string `xml:"CountryCode"`
	} `xml:"ShippingAddress"`
	IsISPU                       string `xml:"IsISPU"`
	IsPrime                      string `xml:"IsPrime"`
	SellerOrderID                string `xml:"SellerOrderId"`
	ShipmentServiceLevelCategory string `xml:"ShipmentServiceLevelCategory"`
}
