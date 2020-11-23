package mws

import "encoding/xml"

// ListOrderItemsResponse ...
type ListOrderItemsResponse struct {
	XMLName              xml.Name             `xml:"ListOrderItemsResponse"`
	Xmlns                string               `xml:"xmlns,attr"`
	ListOrderItemsResult ListOrderItemsResult `xml:"ListOrderItemsResult"`
	ResponseMetadata     ResponseMetadata     `xml:"ResponseMetadata"`
}

// ListOrderItemsByNextTokenResponse ...
type ListOrderItemsByNextTokenResponse struct {
	XMLName              xml.Name             `xml:"ListOrderItemsByNextTokenResponse"`
	Xmlns                string               `xml:"xmlns,attr"`
	ListOrderItemsResult ListOrderItemsResult `xml:"ListOrderItemsByNextTokenResult"`
	ResponseMetadata     ResponseMetadata     `xml:"ResponseMetadata"`
}

// ListOrderItemsResult ...
type ListOrderItemsResult struct {
	AmazonOrderID string        `xml:"AmazonOrderId"`
	NextToken     string        `xml:"NextToken"`
	OrderItemList OrderItemList `xml:"OrderItems"`
}

// OrderItemList ...
type OrderItemList struct {
	OrderItems []OrderItem `xml:"OrderItem"`
}

// OrderItem ...
type OrderItem struct {
	QuantityOrdered   string `xml:"QuantityOrdered"`
	Title             string `xml:"Title"`
	PromotionDiscount struct {
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"PromotionDiscount"`
	IsGift         string `xml:"IsGift"`
	ASIN           string `xml:"ASIN"`
	SellerSKU      string `xml:"SellerSKU"`
	OrderItemID    string `xml:"OrderItemId"`
	IsTransparency string `xml:"IsTransparency"`
	ProductInfo    struct {
		NumberOfItems string `xml:"NumberOfItems"`
	} `xml:"ProductInfo"`
	QuantityShipped string `xml:"QuantityShipped"`
	ItemPrice       struct {
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"ItemPrice"`
	ItemTax struct {
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"ItemTax"`
	PromotionDiscountTax struct {
		Amount       string `xml:"Amount"`
		CurrencyCode string `xml:"CurrencyCode"`
	} `xml:"PromotionDiscountTax"`
}
