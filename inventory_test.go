package mws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListInventorySupply(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}

	assert.NotEqual(t, s.SellerID, "")

	// members := s.ListInventorySupply(ListInventorySupplyParams{
	// 	SellerSkus:         []string{},
	// 	QueryStartDateTime: time.Now().Add(-24 * time.Hour),
	// 	ResponseGroup:      SupplyDetailResponseDetail,
	// })

	// t.Logf("%+v", members)
}
