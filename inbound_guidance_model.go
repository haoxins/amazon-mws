package mws

import "encoding/xml"

// GetInboundGuidanceForSKUResponse ...
type GetInboundGuidanceForSKUResponse struct {
	XMLName                        xml.Name                       `xml:"GetInboundGuidanceForSKUResponse"`
	Xmlns                          string                         `xml:"xmlns,attr"`
	GetInboundGuidanceForSKUResult GetInboundGuidanceForSKUResult `xml:"GetInboundGuidanceForSKUResult"`

	InvalidSKUList   InvalidSKUList   `xml:"InvalidSKUList"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// GetInboundGuidanceForSKUResult ...
type GetInboundGuidanceForSKUResult struct {
	SKUInboundGuidanceList SKUInboundGuidanceList `xml:"SKUInboundGuidanceList"`
	InvalidSKUList         InvalidSKUList         `xml:"InvalidSKUList"`
}

// SKUInboundGuidanceList ...
type SKUInboundGuidanceList struct {
	SKUInboundGuidances []SKUInboundGuidance `xml:"SKUInboundGuidance"`
}

// InvalidSKUList ...
type InvalidSKUList struct {
	XMLName     xml.Name     `xml:"InvalidSKUList"`
	InvalidSKUs []InvalidSKU `xml:"InvalidSKU"`
}

// InvalidSKU ...
type InvalidSKU struct {
	ErrorReason string `xml:"ErrorReason"`
	SellerSKU   string `xml:"SellerSKU"`
}

// SKUInboundGuidance ...
type SKUInboundGuidance struct {
	InboundGuidance    string             `xml:"InboundGuidance"`
	GuidanceReasonList GuidanceReasonList `xml:"GuidanceReasonList"`
	ASIN               string             `xml:"ASIN"`
	SellerSKU          string             `xml:"SellerSKU"`
}

// GuidanceReasonList ...
type GuidanceReasonList struct {
	GuidanceReason string `xml:"GuidanceReason"`
}
