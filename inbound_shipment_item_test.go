package mws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListInboundShipmentItems(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}

	assert.NotEqual(t, s.SellerID, "")

	// members := s.ListInboundShipmentItems(ListInboundShipmentItemsParams{
	// 	ShipmentID:        "",
	// 	LastUpdatedAfter:  time.Now().Add(-24 * time.Hour),
	// 	LastUpdatedBefore: time.Now().Add(-1 * time.Hour),
	// })

	// pp.Printf("Members %+v", members)
}
