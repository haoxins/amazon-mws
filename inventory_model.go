package mws

import "encoding/xml"

// ListInventorySupplyResponse ...
type ListInventorySupplyResponse struct {
	XMLName                   xml.Name                  `xml:"ListInventorySupplyResponse"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListInventorySupplyResult ListInventorySupplyResult `xml:"ListInventorySupplyResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListInventorySupplyByNextTokenResponse ...
type ListInventorySupplyByNextTokenResponse struct {
	XMLName                   xml.Name                  `xml:"ListInventorySupplyByNextTokenResponse"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListInventorySupplyResult ListInventorySupplyResult `xml:"ListInventorySupplyByNextTokenResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListInventorySupplyResult ...
type ListInventorySupplyResult struct {
	MarketplaceID       string              `xml:"MarketplaceId"`
	InventorySupplyList InventorySupplyList `xml:"InventorySupplyList"`
	NextToken           string              `xml:"NextToken"`
}

// InventorySupplyList ...
type InventorySupplyList struct {
	Members []InventorySupplyMember `xml:"member"`
}

// InventorySupplyMember ...
type InventorySupplyMember struct {
	Condition             string               `xml:"Condition"`
	SupplyDetail          string               `xml:"SupplyDetail"`
	TotalSupplyQuantity   string               `xml:"TotalSupplyQuantity"`
	EarliestAvailability  EarliestAvailability `xml:"EarliestAvailability"`
	FNSKU                 string               `xml:"FNSKU"`
	InStockSupplyQuantity string               `xml:"InStockSupplyQuantity"`
	ASIN                  string               `xml:"ASIN"`
	SellerSKU             string               `xml:"SellerSKU"`
}

// EarliestAvailability ...
type EarliestAvailability struct {
	TimepointType string `xml:"TimepointType"`
	DateTime      string `xml:"DateTime"`
}
