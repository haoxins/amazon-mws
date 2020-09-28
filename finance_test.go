package mws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListFinancialEvents(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}

	assert.NotEqual(t, s.SellerID, "")

	// s.ListFinancialEvents(ListFinancialEventsParams{
	// 	PostedAfter:  time.Now().Add(-2 * time.Hour),
	// 	PostedBefore: time.Now().Add(-1 * time.Hour),
	// })
}
