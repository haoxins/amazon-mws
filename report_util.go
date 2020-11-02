package mws

import (
	"strings"

	"github.com/spf13/cast"
)

// ParseSettlementReport Parse Amazon settlement report
func ParseSettlementReport(csv string) []SettlementReportRow {
	text := strings.TrimSpace(strings.ReplaceAll(csv, "\t", ","))
	rows := []SettlementReportRow{}

	for _, v := range strings.Split(text, "\n") {
		if strings.TrimSpace(v) == "" {
			continue
		}
		if strings.Contains(v, "settlement-id,settlement-start-date,settlement-end-date") {
			// Header
			continue
		}

		cols := strings.Split(v, ",")
		if len(cols) < 24 {
			rest := []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""} // Never mind, 24
			cols = append(cols, rest...)
		}

		settlementID := cols[0]
		settlementStartDate := cols[1]
		settlementEndDate := cols[2]
		depositDate := cols[3]
		var totalAmount float64
		if cols[4] != "" {
			totalAmount = cast.ToFloat64(cols[4])
		}
		currency := cols[5]
		transactionType := cols[6]
		orderID := cols[7]
		merchantOrderID := cols[8]
		adjustmentID := cols[9]
		shipmentID := cols[10]
		marketplaceName := cols[11]
		amountType := cols[12]
		amountDescription := cols[13]
		var amount float64
		if cols[14] != "" {
			amount = cast.ToFloat64(cols[14])
		}
		fulfillmentID := cols[15]
		postedDate := cols[16]
		postedDateTime := cols[17]
		orderItemCode := cols[18]
		merchantOrderItemID := cols[19]
		merchantAdjustmentItemID := cols[20]
		sku := cols[21]
		quantityPurchased := cols[22]
		promotionID := cols[23]

		rows = append(rows, SettlementReportRow{
			SettlementID:             settlementID,
			SettlementStartDate:      settlementStartDate,
			SettlementEndDate:        settlementEndDate,
			DepositDate:              depositDate,
			TotalAmount:              totalAmount,
			Currency:                 currency,
			TransactionType:          transactionType,
			OrderID:                  orderID,
			MerchantOrderID:          merchantOrderID,
			AdjustmentID:             adjustmentID,
			ShipmentID:               shipmentID,
			MarketplaceName:          marketplaceName,
			AmountType:               amountType,
			AmountDescription:        amountDescription,
			Amount:                   amount,
			FulfillmentID:            fulfillmentID,
			PostedDate:               postedDate,
			PostedDateTime:           postedDateTime,
			OrderItemCode:            orderItemCode,
			MerchantOrderItemID:      merchantOrderItemID,
			MerchantAdjustmentItemID: merchantAdjustmentItemID,
			Sku:                      sku,
			QuantityPurchased:        quantityPurchased,
			PromotionID:              promotionID,
		})
	}

	return rows
}

// ParseInventoryReport Parse Amazon inventory report
func ParseInventoryReport(csv string) []InventoryReportRow {
	text := strings.TrimSpace(strings.ReplaceAll(csv, "\t", ","))
	rows := []InventoryReportRow{}

	for _, v := range strings.Split(text, "\n") {
		if strings.TrimSpace(v) == "" {
			continue
		}
		// TODO - parse more columns
		if strings.Contains(v, "sku,asin,price") {
			// Header
			continue
		}

		cols := strings.Split(v, ",")
		if len(cols) < 4 {
			rest := []string{"", "", "", ""}
			cols = append(cols, rest...)
		}

		sku := cols[0]
		asin := cols[1]
		var price float64
		if cols[2] != "" {
			price = cast.ToFloat64(cols[2])
		}
		var quantity float64
		if cols[3] != "" {
			quantity = cast.ToFloat64(cols[3])
		}

		rows = append(rows, InventoryReportRow{
			Sku:      sku,
			Asin:     asin,
			Price:    price,
			Quantity: quantity,
		})
	}

	return rows
}
