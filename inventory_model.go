package mws

import "encoding/xml"

// ListInventorySupplyResponse ...
type ListInventorySupplyResponse struct {
	XMLName                   xml.Name                  `xml:"ListInventorySupplyResponse"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListInventorySupplyResult ListInventorySupplyResult `xml:"ListInventorySupplyResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListInventorySupplyByNextTokenResponse ...
type ListInventorySupplyByNextTokenResponse struct {
	XMLName                   xml.Name                  `xml:"ListInventorySupplyByNextTokenResponse"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListInventorySupplyResult ListInventorySupplyResult `xml:"ListInventorySupplyByNextTokenResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListInventorySupplyResult ...
type ListInventorySupplyResult struct {
	Text                string              `xml:",chardata"`
	MarketplaceID       string              `xml:"MarketplaceId"`
	InventorySupplyList InventorySupplyList `xml:"InventorySupplyList"`
	NextToken           string              `xml:"NextToken"`
}

// InventorySupplyList ...
type InventorySupplyList struct {
	Text    string                  `xml:",chardata"`
	Members []InventorySupplyMember `xml:"member"`
}

// InventorySupplyMember ...
type InventorySupplyMember struct {
	Text                 string `xml:",chardata"`
	Condition            string `xml:"Condition"`
	SupplyDetail         string `xml:"SupplyDetail"`
	TotalSupplyQuantity  string `xml:"TotalSupplyQuantity"`
	EarliestAvailability struct {
		Text          string `xml:",chardata"`
		TimepointType string `xml:"TimepointType"`
		DateTime      string `xml:"DateTime"`
	} `xml:"EarliestAvailability"`
	FNSKU                 string `xml:"FNSKU"`
	InStockSupplyQuantity string `xml:"InStockSupplyQuantity"`
	ASIN                  string `xml:"ASIN"`
	SellerSKU             string `xml:"SellerSKU"`
}
