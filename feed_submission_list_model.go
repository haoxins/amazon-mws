package mws

import "encoding/xml"

// GetFeedSubmissionListResponse ...
type GetFeedSubmissionListResponse struct {
	XMLName                     xml.Name                    `xml:"GetFeedSubmissionListResponse"`
	Xmlns                       string                      `xml:"xmlns,attr"`
	GetFeedSubmissionListResult GetFeedSubmissionListResult `xml:"GetFeedSubmissionListResult"`
	ResponseMetadata            ResponseMetadata            `xml:"ResponseMetadata"`
}

// GetFeedSubmissionListByNextTokenResponse ...
type GetFeedSubmissionListByNextTokenResponse struct {
	XMLName                     xml.Name                    `xml:"GetFeedSubmissionListByNextTokenResponse"`
	Xmlns                       string                      `xml:"xmlns,attr"`
	GetFeedSubmissionListResult GetFeedSubmissionListResult `xml:"GetFeedSubmissionListByNextTokenResult"`
	ResponseMetadata            ResponseMetadata            `xml:"ResponseMetadata"`
}

// GetFeedSubmissionListResult ...
type GetFeedSubmissionListResult struct {
	NextToken           string               `xml:"NextToken"`
	HasNext             bool                 `xml:"HasNext"`
	FeedSubmissionInfos []FeedSubmissionInfo `xml:"FeedSubmissionInfo"`
}

// FeedSubmissionInfo ...
type FeedSubmissionInfo struct {
	FeedProcessingStatus    string `xml:"FeedProcessingStatus"`
	FeedType                string `xml:"FeedType"`
	FeedSubmissionID        string `xml:"FeedSubmissionId"`
	StartedProcessingDate   string `xml:"StartedProcessingDate"`
	SubmittedDate           string `xml:"SubmittedDate"`
	CompletedProcessingDate string `xml:"CompletedProcessingDate"`
}
