package mws

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_GetFeedSubmissionList(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}

	assert.NotEqual(t, s.SellerID, "")

	feeds := s.GetFeedSubmissionList(GetFeedSubmissionListParams{
		SubmittedFromDate: time.Now().Add(-8 * 24 * time.Hour),
		SubmittedToDate:   time.Now().Add(-1 * time.Hour),
	})

	t.Logf("%+v", feeds)
}
