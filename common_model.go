package mws

import "encoding/xml"

// ResponseMetadata ...
type ResponseMetadata struct {
	RequestID string `xml:"RequestId"`
}

// ErrorResponse ...
type ErrorResponse struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Xmlns   string   `xml:"xmlns,attr"`
	Error   struct {
		Type    string `xml:"Type"`
		Code    string `xml:"Code"`
		Message string `xml:"Message"`
	} `xml:"Error"`
	RequestID string `xml:"RequestID"`
}
