package mws

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetInboundGuidanceForSKU(t *testing.T) {
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

	// t.Logf("uids length: %d", len(uids))

	// csv := s.GetReportByID(uids[0])
	// rows := ParseInventoryReport(csv)

	// t.Logf("rows length: %d", len(rows))

	// var skus []string
	// for _, r := range rows {
	// 	skus = append(skus, r.Sku)
	// }

	// skus = funk.UniqString(skus)

	// guidances, invalidSkus := s.GetInboundGuidanceForSKU(GetInboundGuidanceForSKUParams{
	// 	SellerSKUList: skus,
	// })

	// t.Logf("guidances length %d, invalid skus length %d", len(guidances), len(invalidSkus))
	// t.Logf("guidance %+v", guidances[0])
	// t.Logf("invalid skus %+v", invalidSkus)
}
