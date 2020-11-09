package mws

import "encoding/xml"

// ListInboundShipmentItemsResponse ...
type ListInboundShipmentItemsResponse struct {
	XMLName                        xml.Name                       `xml:"ListInboundShipmentItemsResponse"`
	Xmlns                          string                         `xml:"xmlns,attr"`
	ListInboundShipmentItemsResult ListInboundShipmentItemsResult `xml:"ListInboundShipmentItemsResult"`
	ResponseMetadata               ResponseMetadata               `xml:"ResponseMetadata"`
}

// ListInboundShipmentItemsByNextTokenResponse ...
type ListInboundShipmentItemsByNextTokenResponse struct {
	XMLName                        xml.Name                       `xml:"ListInboundShipmentItemsByNextTokenResponse"`
	Xmlns                          string                         `xml:"xmlns,attr"`
	ListInboundShipmentItemsResult ListInboundShipmentItemsResult `xml:"ListInboundShipmentItemsByNextTokenResult"`
	ResponseMetadata               ResponseMetadata               `xml:"ResponseMetadata"`
}

// ListInboundShipmentItemsResult ...
type ListInboundShipmentItemsResult struct {
	ItemData  InboundShipmentItemData `xml:"ItemData"`
	NextToken string                  `xml:"NextToken"`
}

// InboundShipmentItemData ...
type InboundShipmentItemData struct {
	Members []InboundShipmentItemMember `xml:"member"`
}

// InboundShipmentItemMember ...
type InboundShipmentItemMember struct {
	QuantityShipped       string          `xml:"QuantityShipped"`
	ShipmentID            string          `xml:"ShipmentId"`
	PrepDetailsList       PrepDetailsList `xml:"PrepDetailsList"`
	FulfillmentNetworkSKU string          `xml:"FulfillmentNetworkSKU"`
	SellerSKU             string          `xml:"SellerSKU"`
	QuantityReceived      string          `xml:"QuantityReceived"`
	QuantityInCase        string          `xml:"QuantityInCase"`
}

// PrepDetailsList ...
type PrepDetailsList struct {
	PrepDetails struct {
		PrepOwner       string `xml:"PrepOwner"`
		PrepInstruction string `xml:"PrepInstruction"`
	} `xml:"PrepDetails"`
}
