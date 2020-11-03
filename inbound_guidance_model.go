package mws

import "encoding/xml"

// GetInboundGuidanceForSKUResponse ...
type GetInboundGuidanceForSKUResponse struct {
	XMLName                        xml.Name                       `xml:"GetInboundGuidanceForSKUResponse"`
	Text                           string                         `xml:",chardata"`
	Xmlns                          string                         `xml:"xmlns,attr"`
	GetInboundGuidanceForSKUResult GetInboundGuidanceForSKUResult `xml:"GetInboundGuidanceForSKUResult"`

	InvalidSKUList   InvalidSKUList   `xml:"InvalidSKUList"`
	ResponseMetadata ResponseMetadata `xml:"ResponseMetadata"`
}

// GetInboundGuidanceForSKUResult ...
type GetInboundGuidanceForSKUResult struct {
	Text                   string                 `xml:",chardata"`
	SKUInboundGuidanceList SKUInboundGuidanceList `xml:"SKUInboundGuidanceList"`
	InvalidSKUList         InvalidSKUList         `xml:"InvalidSKUList"`
}

// SKUInboundGuidanceList ...
type SKUInboundGuidanceList struct {
	Text                string               `xml:",chardata"`
	SKUInboundGuidances []SKUInboundGuidance `xml:"SKUInboundGuidance"`
}

// InvalidSKUList ...
type InvalidSKUList struct {
	XMLName     xml.Name     `xml:"InvalidSKUList"`
	Text        string       `xml:",chardata"`
	InvalidSKUs []InvalidSKU `xml:"InvalidSKU"`
}

// InvalidSKU ...
type InvalidSKU struct {
	Text        string `xml:",chardata"`
	ErrorReason string `xml:"ErrorReason"`
	SellerSKU   string `xml:"SellerSKU"`
}

// SKUInboundGuidance ...
type SKUInboundGuidance struct {
	Text               string `xml:",chardata"`
	InboundGuidance    string `xml:"InboundGuidance"`
	GuidanceReasonList struct {
		Text           string `xml:",chardata"`
		GuidanceReason string `xml:"GuidanceReason"`
	} `xml:"GuidanceReasonList"`
	ASIN      string `xml:"ASIN"`
	SellerSKU string `xml:"SellerSKU"`
}
