package mws

import "encoding/xml"

// ListInboundShipmentsResponse ...
type ListInboundShipmentsResponse struct {
	XMLName                    xml.Name                   `xml:"ListInboundShipmentsResponse"`
	Xmlns                      string                     `xml:"xmlns,attr"`
	ListInboundShipmentsResult ListInboundShipmentsResult `xml:"ListInboundShipmentsResult"`
	ResponseMetadata           ResponseMetadata           `xml:"ResponseMetadata"`
}

// ListInboundShipmentsByNextTokenResponse ...
type ListInboundShipmentsByNextTokenResponse struct {
	XMLName                    xml.Name                   `xml:"ListInboundShipmentsByNextTokenResponse"`
	Xmlns                      string                     `xml:"xmlns,attr"`
	ListInboundShipmentsResult ListInboundShipmentsResult `xml:"ListInboundShipmentsByNextTokenResult"`
	ResponseMetadata           ResponseMetadata           `xml:"ResponseMetadata"`
}

// ListInboundShipmentsResult ...
type ListInboundShipmentsResult struct {
	NextToken    string       `xml:"NextToken"`
	ShipmentData ShipmentData `xml:"ShipmentData"`
}

// ShipmentData ...
type ShipmentData struct {
	Members []ShipmentMember `xml:"member"`
}

// ShipmentMember ...
type ShipmentMember struct {
	DestinationFulfillmentCenterID string          `xml:"DestinationFulfillmentCenterId"`
	LabelPrepType                  string          `xml:"LabelPrepType"`
	ShipFromAddress                ShipFromAddress `xml:"ShipFromAddress"`
	ShipmentID                     string          `xml:"ShipmentId"`
	AreCasesRequired               string          `xml:"AreCasesRequired"`
	ShipmentName                   string          `xml:"ShipmentName"`
	BoxContentsSource              string          `xml:"BoxContentsSource"`
	ShipmentStatus                 string          `xml:"ShipmentStatus"`
	ConfirmedNeedByDate            string          `xml:"ConfirmedNeedByDate"`
}

// ShipFromAddress ...
type ShipFromAddress struct {
	City                string `xml:"City"`
	CountryCode         string `xml:"CountryCode"`
	PostalCode          string `xml:"PostalCode"`
	Name                string `xml:"Name"`
	AddressLine1        string `xml:"AddressLine1"`
	StateOrProvinceCode string `xml:"StateOrProvinceCode"`
	DistrictOrCounty    string `xml:"DistrictOrCounty"`
	AddressLine2        string `xml:"AddressLine2"`
}
