package mws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetReportList(t *testing.T) {
	s := &Seller{
		Country:   os.Getenv("country"),
		SellerID:  os.Getenv("seller_id"),
		AuthToken: os.Getenv("auth_token"),
		AccessKey: os.Getenv("access_key"),
		SecretKey: os.Getenv("secret_key"),
	}

	assert.NotEqual(t, s.SellerID, "")

	// uids := s.GetReportList(GetReportListParams{
	// 	ReportType: InventoryReport,
	// 	StartTime:  time.Now().Add(-24 * time.Hour),
	// 	EndTime:    time.Now().Add(-1 * time.Hour),
	// })

	// t.Logf("uids: %+v", uids)

	// csv := s.GetReportByID(uids[0])

	// t.Logf("Report %+v", ParseInventoryReport(csv))
}
