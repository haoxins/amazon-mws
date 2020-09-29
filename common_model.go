package mws

// ResponseMetadata ...
type ResponseMetadata struct {
	Text      string `xml:",chardata"`
	RequestID string `xml:"RequestId"`
}
