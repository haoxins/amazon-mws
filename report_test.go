package mws

import (
	"os"
	"testing"
)

func Test_GetReportList(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}
	t.Logf("%+v", s)
	// s.GetReportList(time.Now().Add(-8*24*time.Hour), time.Now().Add(-1*time.Hour), "")
}
