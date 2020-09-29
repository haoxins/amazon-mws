package mws

import "encoding/xml"

// ListInboundShipmentsResponse ...
type ListInboundShipmentsResponse struct {
	XMLName                    xml.Name                   `xml:"ListInboundShipmentsResponse"`
	Text                       string                     `xml:",chardata"`
	Xmlns                      string                     `xml:"xmlns,attr"`
	ListInboundShipmentsResult ListInboundShipmentsResult `xml:"ListInboundShipmentsResult"`
	ResponseMetadata           ResponseMetadata           `xml:"ResponseMetadata"`
}

// ListInboundShipmentsByNextTokenResponse ...
type ListInboundShipmentsByNextTokenResponse struct {
	XMLName                    xml.Name                   `xml:"ListInboundShipmentsByNextTokenResponse"`
	Text                       string                     `xml:",chardata"`
	Xmlns                      string                     `xml:"xmlns,attr"`
	ListInboundShipmentsResult ListInboundShipmentsResult `xml:"ListInboundShipmentsByNextTokenResult"`
	ResponseMetadata           ResponseMetadata           `xml:"ResponseMetadata"`
}

// ListInboundShipmentsResult ...
type ListInboundShipmentsResult struct {
	Text         string       `xml:",chardata"`
	NextToken    string       `xml:"NextToken"`
	ShipmentData ShipmentData `xml:"ShipmentData"`
}

// ShipmentData ...
type ShipmentData struct {
	Text   string `xml:",chardata"`
	Member []struct {
		Text                           string `xml:",chardata"`
		DestinationFulfillmentCenterID string `xml:"DestinationFulfillmentCenterId"`
		LabelPrepType                  string `xml:"LabelPrepType"`
		ShipFromAddress                struct {
			Text                string `xml:",chardata"`
			City                string `xml:"City"`
			CountryCode         string `xml:"CountryCode"`
			PostalCode          string `xml:"PostalCode"`
			Name                string `xml:"Name"`
			AddressLine1        string `xml:"AddressLine1"`
			StateOrProvinceCode string `xml:"StateOrProvinceCode"`
			DistrictOrCounty    string `xml:"DistrictOrCounty"`
			AddressLine2        string `xml:"AddressLine2"`
		} `xml:"ShipFromAddress"`
		ShipmentID          string `xml:"ShipmentId"`
		AreCasesRequired    string `xml:"AreCasesRequired"`
		ShipmentName        string `xml:"ShipmentName"`
		BoxContentsSource   string `xml:"BoxContentsSource"`
		ShipmentStatus      string `xml:"ShipmentStatus"`
		ConfirmedNeedByDate string `xml:"ConfirmedNeedByDate"`
	} `xml:"member"`
}
