package mws

import "encoding/xml"

// OrderStatus ...
type OrderStatus string

const (
	// OrderStatusPendingAvailability ...
	OrderStatusPendingAvailability OrderStatus = "PendingAvailability"
	// OrderStatusPending ...
	OrderStatusPending = "Pending"
	// OrderStatusUnshipped ...
	OrderStatusUnshipped = "Unshipped"
	// OrderStatusPartiallyShipped ...
	OrderStatusPartiallyShipped = "PartiallyShipped"
	// OrderStatusShipped ...
	OrderStatusShipped = "Shipped"
	// OrderStatusCanceled ...
	OrderStatusCanceled = "Canceled"
	// OrderStatusUnfulfillable ...
	OrderStatusUnfulfillable = "Unfulfillable"
)

// ListOrdersResponse ...
type ListOrdersResponse struct {
	XMLName          xml.Name         `xml:"ListOrdersResponse"`
	Xmlns            string           `xml:"xmlns,attr"`
	ListOrdersResult ListOrdersResult `xml:"ListOrdersResult"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// ListOrdersResult ...
type ListOrdersResult struct {
	NextToken         string    `xml:"NextToken"`
	Orders            OrderList `xml:"Orders"`
	LastUpdatedBefore string    `xml:"LastUpdatedBefore"`
}

// OrderList ...
type OrderList struct {
	Orders []OrderInfo `xml:"Order"`
}

// OrderInfo ...
type OrderInfo struct {
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
		PaymentMethodDetail string `xml:"PaymentMethodDetail"`
	} `xml:"PaymentMethodDetails"`
	IsGlobalExpressEnabled string `xml:"IsGlobalExpressEnabled"`
	IsSoldByAB             string `xml:"IsSoldByAB"`
	IsPremiumOrder         string `xml:"IsPremiumOrder"`
	OrderTotal             struct {
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"OrderTotal"`
	EarliestShipDate   string `xml:"EarliestShipDate"`
	MarketplaceID      string `xml:"MarketplaceId"`
	FulfillmentChannel string `xml:"FulfillmentChannel"`
	PaymentMethod      string `xml:"PaymentMethod"`
	ShippingAddress    struct {
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
