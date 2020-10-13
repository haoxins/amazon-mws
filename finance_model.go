package mws

import "encoding/xml"

// ListFinancialEventsResponse ...
type ListFinancialEventsResponse struct {
	XMLName                   xml.Name                  `xml:"ListFinancialEventsResponse"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListFinancialEventsByNextTokenResponse ...
type ListFinancialEventsByNextTokenResponse struct {
	XMLName                   xml.Name                  `xml:"ListFinancialEventsByNextTokenResponse"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsByNextTokenResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListFinancialEventsResult ...
type ListFinancialEventsResult struct {
	NextToken       string          `xml:"NextToken"`
	FinancialEvents FinancialEvents `xml:"FinancialEvents"`
}

// FinancialEvents ...
type FinancialEvents struct {
	ProductAdsPaymentEventList             ProductAdsPaymentEventList `xml:"ProductAdsPaymentEventList"`
	RentalTransactionEventList             string                     `xml:"RentalTransactionEventList"`
	ServiceFeeEventList                    ServiceFeeEventList        `xml:"ServiceFeeEventList"`
	ShipmentSettleEventList                string                     `xml:"ShipmentSettleEventList"`
	ServiceProviderCreditEventList         string                     `xml:"ServiceProviderCreditEventList"`
	ImagingServicesFeeEventList            string                     `xml:"ImagingServicesFeeEventList"`
	SellerDealPaymentEventList             string                     `xml:"SellerDealPaymentEventList"`
	SellerReviewEnrollmentPaymentEventList string                     `xml:"SellerReviewEnrollmentPaymentEventList"`
	DebtRecoveryEventList                  string                     `xml:"DebtRecoveryEventList"`
	ShipmentEventList                      ShipmentEventList          `xml:"ShipmentEventList"`
	TaxWithholdingEventList                string                     `xml:"TaxWithholdingEventList"`
	GuaranteeClaimEventList                string                     `xml:"GuaranteeClaimEventList"`
	TDSReimbursementEventList              string                     `xml:"TDSReimbursementEventList"`
	ChargebackEventList                    string                     `xml:"ChargebackEventList"`
	NetworkComminglingTransactionEventList string                     `xml:"NetworkComminglingTransactionEventList"`
	LoanServicingEventList                 string                     `xml:"LoanServicingEventList"`
	RefundEventList                        RefundEventList            `xml:"RefundEventList"`
	RemovalShipmentEventList               string                     `xml:"RemovalShipmentEventList"`
	PerformanceBondRefundEventList         string                     `xml:"PerformanceBondRefundEventList"`
	AffordabilityExpenseReversalEventList  string                     `xml:"AffordabilityExpenseReversalEventList"`
	PayWithAmazonEventList                 string                     `xml:"PayWithAmazonEventList"`
	AdhocDisbursementEventList             string                     `xml:"AdhocDisbursementEventList"`
	CouponPaymentEventList                 CouponPaymentEventList     `xml:"CouponPaymentEventList"`
	ChargeRefundEventList                  string                     `xml:"ChargeRefundEventList"`
	RetrochargeEventList                   string                     `xml:"RetrochargeEventList"`
	TrialShipmentEventList                 string                     `xml:"TrialShipmentEventList"`
	SAFETReimbursementEventList            string                     `xml:"SAFETReimbursementEventList"`
	RemovalShipmentAdjustmentEventList     string                     `xml:"RemovalShipmentAdjustmentEventList"`
	FBALiquidationEventList                string                     `xml:"FBALiquidationEventList"`
	AffordabilityExpenseEventList          string                     `xml:"AffordabilityExpenseEventList"`
	AdjustmentEventList                    AdjustmentEventList        `xml:"AdjustmentEventList"`
}

// ProductAdsPaymentEventList ...
type ProductAdsPaymentEventList struct {
	ProductAdsPaymentEvents []struct {
		TransactionType string `xml:"transactionType"`
		BaseValue       struct {
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"baseValue"`
		TaxValue struct {
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"taxValue"`
		InvoiceID        string `xml:"invoiceId"`
		TransactionValue struct {
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"transactionValue"`
		PostedDate string `xml:"postedDate"`
	} `xml:"ProductAdsPaymentEvent"`
}

// ServiceFeeEventList ...
type ServiceFeeEventList struct {
	ServiceFeeEvents []FinancialServiceFeeEvent `xml:"ServiceFeeEvent"`
}

// FinancialServiceFeeEvent ...
type FinancialServiceFeeEvent struct {
	FeeReason     string `xml:"FeeReason"`
	AmazonOrderID string `xml:"AmazonOrderId"`
	FeeList       struct {
		FeeComponents []FeeComponent `xml:"FeeComponent"`
	} `xml:"FeeList"`
}

// ShipmentEventList ...
type ShipmentEventList struct {
	ShipmentEvents []FinancialShipmentEvent `xml:"ShipmentEvent"`
}

// FinancialShipmentEvent ...
type FinancialShipmentEvent struct {
	ShipmentItemList struct {
		ShipmentItems []ShipmentItem `xml:"ShipmentItem"`
	} `xml:"ShipmentItemList"`
	AmazonOrderID   string `xml:"AmazonOrderId"`
	PostedDate      string `xml:"PostedDate"`
	MarketplaceName string `xml:"MarketplaceName"`
	SellerOrderID   string `xml:"SellerOrderId"`
	ShipmentFeeList struct {
		FeeComponents []FeeComponent `xml:"FeeComponent"`
	} `xml:"ShipmentFeeList"`
}

// RefundEventList ...
type RefundEventList struct {
	ShipmentEvents []FinancialRefundShipmentEvent `xml:"ShipmentEvent"`
}

// FinancialRefundShipmentEvent ...
type FinancialRefundShipmentEvent struct {
	AmazonOrderID              string `xml:"AmazonOrderId"`
	PostedDate                 string `xml:"PostedDate"`
	ShipmentItemAdjustmentList struct {
		ShipmentItems []ShipmentItemAdjustment `xml:"ShipmentItem"`
	} `xml:"ShipmentItemAdjustmentList"`
	MarketplaceName string `xml:"MarketplaceName"`
	SellerOrderID   string `xml:"SellerOrderId"`
}

// CouponPaymentEventList ...
type CouponPaymentEventList struct {
	CouponPaymentEvents []FinancialCouponPaymentEvent `xml:"CouponPaymentEvent"`
}

// FinancialCouponPaymentEvent ...
type FinancialCouponPaymentEvent struct {
	TotalAmount struct {
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"TotalAmount"`
	PaymentEventID          string          `xml:"PaymentEventId"`
	SellerCouponDescription string          `xml:"SellerCouponDescription"`
	FeeComponent            FeeComponent    `xml:"FeeComponent"`
	ChargeComponent         ChargeComponent `xml:"ChargeComponent"`
	CouponID                string          `xml:"CouponId"`
	ClipOrRedemptionCount   string          `xml:"ClipOrRedemptionCount"`
	PostedDate              string          `xml:"PostedDate"`
}

// AdjustmentEventList ...
type AdjustmentEventList struct {
	AdjustmentEvents []FinancialAdjustmentEvent `xml:"AdjustmentEvent"`
}

// FinancialAdjustmentEvent ...
type FinancialAdjustmentEvent struct {
	AdjustmentType     string `xml:"AdjustmentType"`
	AdjustmentItemList struct {
		AdjustmentItems []struct {
			PerUnitAmount struct {
				CurrencyAmount string `xml:"CurrencyAmount"`
				CurrencyCode   string `xml:"CurrencyCode"`
			} `xml:"PerUnitAmount"`
			TotalAmount struct {
				CurrencyAmount string `xml:"CurrencyAmount"`
				CurrencyCode   string `xml:"CurrencyCode"`
			} `xml:"TotalAmount"`
			Quantity           string `xml:"Quantity"`
			SellerSKU          string `xml:"SellerSKU"`
			ProductDescription string `xml:"ProductDescription"`
		} `xml:"AdjustmentItem"`
	} `xml:"AdjustmentItemList"`
	AdjustmentAmount struct {
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"AdjustmentAmount"`
	PostedDate string `xml:"PostedDate"`
}

// ChargeComponent ...
type ChargeComponent struct {
	ChargeType   string `xml:"ChargeType"`
	ChargeAmount struct {
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"ChargeAmount"`
}

// FeeComponent ...
type FeeComponent struct {
	FeeType   string `xml:"FeeType"`
	FeeAmount struct {
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"FeeAmount"`
}

// Promotion ...
type Promotion struct {
	PromotionType   string `xml:"PromotionType"`
	PromotionAmount struct {
		CurrencyAmount string `xml:"CurrencyAmount"`
		CurrencyCode   string `xml:"CurrencyCode"`
	} `xml:"PromotionAmount"`
	PromotionID string `xml:"PromotionId"`
}

// ItemTaxWithheldList ...
type ItemTaxWithheldList struct {
	TaxWithheldComponent struct {
		TaxCollectionModel string `xml:"TaxCollectionModel"`
		TaxesWithheld      struct {
			ChargeComponents []ChargeComponent `xml:"ChargeComponent"`
		} `xml:"TaxesWithheld"`
	} `xml:"TaxWithheldComponent"`
}

// ShipmentItem ...
type ShipmentItem struct {
	ItemTaxWithheldList ItemTaxWithheldList `xml:"ItemTaxWithheldList"`
	ItemChargeList      ItemChargeList      `xml:"ItemChargeList"`
	ItemFeeList         ItemFeeList         `xml:"ItemFeeList"`
	OrderItemID         string              `xml:"OrderItemId"`
	QuantityShipped     string              `xml:"QuantityShipped"`
	SellerSKU           string              `xml:"SellerSKU"`
	PromotionList       PromotionList       `xml:"PromotionList"`
}

// ItemFeeList ...
type ItemFeeList struct {
	FeeComponents []FeeComponent `xml:"FeeComponent"`
}

// ItemChargeList ...
type ItemChargeList struct {
	ChargeComponents []ChargeComponent `xml:"ChargeComponent"`
}

// ShipmentItemAdjustment ...
type ShipmentItemAdjustment struct {
	ItemTaxWithheldList      ItemTaxWithheldList      `xml:"ItemTaxWithheldList"`
	ItemFeeAdjustmentList    ItemFeeAdjustmentList    `xml:"ItemFeeAdjustmentList"`
	OrderAdjustmentItemID    string                   `xml:"OrderAdjustmentItemId"`
	QuantityShipped          string                   `xml:"QuantityShipped"`
	ItemChargeAdjustmentList ItemChargeAdjustmentList `xml:"ItemChargeAdjustmentList"`
	SellerSKU                string                   `xml:"SellerSKU"`
	PromotionAdjustmentList  PromotionList            `xml:"PromotionAdjustmentList"`
}

// ItemFeeAdjustmentList ...
type ItemFeeAdjustmentList struct {
	FeeComponents []FeeComponent `xml:"FeeComponent"`
}

// ItemChargeAdjustmentList ...
type ItemChargeAdjustmentList struct {
	ChargeComponents []ChargeComponent `xml:"ChargeComponent"`
}

// PromotionList ...
type PromotionList struct {
	Promotions []Promotion `xml:"Promotion"`
}
