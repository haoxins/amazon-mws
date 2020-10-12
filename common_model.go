package mws

import "encoding/xml"

// ResponseMetadata ...
type ResponseMetadata struct {
	Text      string `xml:",chardata"`
	RequestID string `xml:"RequestId"`
}

// ErrorResponse ...
type ErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Error   struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"Type"`
		Code    string `xml:"Code"`
		Message string `xml:"Message"`
	} `xml:"Error"`
	RequestID string `xml:"RequestID"`
}
