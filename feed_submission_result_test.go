package mws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFeedSubmissionResult(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}

	assert.NotEqual(t, s.SellerID, "")

	// TODO - get id from GetFeedSubmissionList
	// id := os.Getenv("feed_submission_id")
	// s.GetFeedSubmissionResult(id)
}
