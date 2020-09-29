package mws

import "encoding/xml"

// ListFinancialEventsResponse ...
type ListFinancialEventsResponse struct {
	XMLName                   xml.Name                  `xml:"ListFinancialEventsResponse"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListFinancialEventsByNextTokenResponse ...
type ListFinancialEventsByNextTokenResponse struct {
	XMLName                   xml.Name                  `xml:"ListFinancialEventsByNextTokenResponse"`
	Text                      string                    `xml:",chardata"`
	Xmlns                     string                    `xml:"xmlns,attr"`
	ListFinancialEventsResult ListFinancialEventsResult `xml:"ListFinancialEventsByNextTokenResult"`
	ResponseMetadata          ResponseMetadata          `xml:"ResponseMetadata"`
}

// ListFinancialEventsResult ...
type ListFinancialEventsResult struct {
	Text            string          `xml:",chardata"`
	NextToken       string          `xml:"NextToken"`
	FinancialEvents FinancialEvents `xml:"FinancialEvents"`
}

// FinancialEvents ...
type FinancialEvents struct {
	Text                                   string                `xml:",chardata"`
	ProductAdsPaymentEventList             string                `xml:"ProductAdsPaymentEventList"`
	RentalTransactionEventList             string                `xml:"RentalTransactionEventList"`
	ServiceFeeEventList                    string                `xml:"ServiceFeeEventList"`
	ShipmentSettleEventList                string                `xml:"ShipmentSettleEventList"`
	ServiceProviderCreditEventList         string                `xml:"ServiceProviderCreditEventList"`
	ImagingServicesFeeEventList            string                `xml:"ImagingServicesFeeEventList"`
	SellerDealPaymentEventList             string                `xml:"SellerDealPaymentEventList"`
	SellerReviewEnrollmentPaymentEventList string                `xml:"SellerReviewEnrollmentPaymentEventList"`
	DebtRecoveryEventList                  string                `xml:"DebtRecoveryEventList"`
	ShipmentEventList                      []ShipmentEventList   `xml:"ShipmentEventList"`
	TaxWithholdingEventList                string                `xml:"TaxWithholdingEventList"`
	GuaranteeClaimEventList                string                `xml:"GuaranteeClaimEventList"`
	TDSReimbursementEventList              string                `xml:"TDSReimbursementEventList"`
	ChargebackEventList                    string                `xml:"ChargebackEventList"`
	NetworkComminglingTransactionEventList string                `xml:"NetworkComminglingTransactionEventList"`
	LoanServicingEventList                 string                `xml:"LoanServicingEventList"`
	RefundEventList                        []RefundEventList     `xml:"RefundEventList"`
	RemovalShipmentEventList               string                `xml:"RemovalShipmentEventList"`
	PerformanceBondRefundEventList         string                `xml:"PerformanceBondRefundEventList"`
	AffordabilityExpenseReversalEventList  string                `xml:"AffordabilityExpenseReversalEventList"`
	PayWithAmazonEventList                 string                `xml:"PayWithAmazonEventList"`
	AdhocDisbursementEventList             string                `xml:"AdhocDisbursementEventList"`
	CouponPaymentEventList                 string                `xml:"CouponPaymentEventList"`
	ChargeRefundEventList                  string                `xml:"ChargeRefundEventList"`
	RetrochargeEventList                   string                `xml:"RetrochargeEventList"`
	TrialShipmentEventList                 string                `xml:"TrialShipmentEventList"`
	SAFETReimbursementEventList            string                `xml:"SAFETReimbursementEventList"`
	RemovalShipmentAdjustmentEventList     string                `xml:"RemovalShipmentAdjustmentEventList"`
	FBALiquidationEventList                string                `xml:"FBALiquidationEventList"`
	AffordabilityExpenseEventList          string                `xml:"AffordabilityExpenseEventList"`
	AdjustmentEventList                    []AdjustmentEventList `xml:"AdjustmentEventList"`
}

// ShipmentEventList ...
type ShipmentEventList struct {
	Text          string `xml:",chardata"`
	ShipmentEvent []struct {
		Text             string `xml:",chardata"`
		ShipmentItemList struct {
			Text         string `xml:",chardata"`
			ShipmentItem []struct {
				Text                string `xml:",chardata"`
				ItemTaxWithheldList struct {
					Text                 string `xml:",chardata"`
					TaxWithheldComponent struct {
						Text               string `xml:",chardata"`
						TaxCollectionModel string `xml:"TaxCollectionModel"`
						TaxesWithheld      struct {
							Text            string `xml:",chardata"`
							ChargeComponent []struct {
								Text         string `xml:",chardata"`
								ChargeType   string `xml:"ChargeType"`
								ChargeAmount struct {
									Text           string `xml:",chardata"`
									CurrencyAmount string `xml:"CurrencyAmount"`
									CurrencyCode   string `xml:"CurrencyCode"`
								} `xml:"ChargeAmount"`
							} `xml:"ChargeComponent"`
						} `xml:"TaxesWithheld"`
					} `xml:"TaxWithheldComponent"`
				} `xml:"ItemTaxWithheldList"`
				ItemChargeList struct {
					Text            string `xml:",chardata"`
					ChargeComponent []struct {
						Text         string `xml:",chardata"`
						ChargeType   string `xml:"ChargeType"`
						ChargeAmount struct {
							Text           string `xml:",chardata"`
							CurrencyAmount string `xml:"CurrencyAmount"`
							CurrencyCode   string `xml:"CurrencyCode"`
						} `xml:"ChargeAmount"`
					} `xml:"ChargeComponent"`
				} `xml:"ItemChargeList"`
				ItemFeeList struct {
					Text         string `xml:",chardata"`
					FeeComponent []struct {
						Text      string `xml:",chardata"`
						FeeAmount struct {
							Text           string `xml:",chardata"`
							CurrencyAmount string `xml:"CurrencyAmount"`
							CurrencyCode   string `xml:"CurrencyCode"`
						} `xml:"FeeAmount"`
						FeeType string `xml:"FeeType"`
					} `xml:"FeeComponent"`
				} `xml:"ItemFeeList"`
				OrderItemID     string `xml:"OrderItemId"`
				QuantityShipped string `xml:"QuantityShipped"`
				SellerSKU       string `xml:"SellerSKU"`
				PromotionList   struct {
					Text      string `xml:",chardata"`
					Promotion []struct {
						Text            string `xml:",chardata"`
						PromotionType   string `xml:"PromotionType"`
						PromotionAmount struct {
							Text           string `xml:",chardata"`
							CurrencyAmount string `xml:"CurrencyAmount"`
							CurrencyCode   string `xml:"CurrencyCode"`
						} `xml:"PromotionAmount"`
						PromotionID string `xml:"PromotionId"`
					} `xml:"Promotion"`
				} `xml:"PromotionList"`
			} `xml:"ShipmentItem"`
		} `xml:"ShipmentItemList"`
		AmazonOrderID   string `xml:"AmazonOrderId"`
		PostedDate      string `xml:"PostedDate"`
		MarketplaceName string `xml:"MarketplaceName"`
		SellerOrderID   string `xml:"SellerOrderId"`
		ShipmentFeeList struct {
			Text         string `xml:",chardata"`
			FeeComponent struct {
				Text      string `xml:",chardata"`
				FeeAmount struct {
					Text           string `xml:",chardata"`
					CurrencyAmount string `xml:"CurrencyAmount"`
					CurrencyCode   string `xml:"CurrencyCode"`
				} `xml:"FeeAmount"`
				FeeType string `xml:"FeeType"`
			} `xml:"FeeComponent"`
		} `xml:"ShipmentFeeList"`
	} `xml:"ShipmentEvent"`
}

// RefundEventList ...
type RefundEventList struct {
	Text          string `xml:",chardata"`
	ShipmentEvent struct {
		Text                       string `xml:",chardata"`
		AmazonOrderID              string `xml:"AmazonOrderId"`
		PostedDate                 string `xml:"PostedDate"`
		ShipmentItemAdjustmentList struct {
			Text         string `xml:",chardata"`
			ShipmentItem struct {
				Text                string `xml:",chardata"`
				ItemTaxWithheldList struct {
					Text                 string `xml:",chardata"`
					TaxWithheldComponent struct {
						Text               string `xml:",chardata"`
						TaxCollectionModel string `xml:"TaxCollectionModel"`
						TaxesWithheld      struct {
							Text            string `xml:",chardata"`
							ChargeComponent struct {
								Text         string `xml:",chardata"`
								ChargeType   string `xml:"ChargeType"`
								ChargeAmount struct {
									Text           string `xml:",chardata"`
									CurrencyAmount string `xml:"CurrencyAmount"`
									CurrencyCode   string `xml:"CurrencyCode"`
								} `xml:"ChargeAmount"`
							} `xml:"ChargeComponent"`
						} `xml:"TaxesWithheld"`
					} `xml:"TaxWithheldComponent"`
				} `xml:"ItemTaxWithheldList"`
				ItemFeeAdjustmentList struct {
					Text         string `xml:",chardata"`
					FeeComponent []struct {
						Text      string `xml:",chardata"`
						FeeAmount struct {
							Text           string `xml:",chardata"`
							CurrencyAmount string `xml:"CurrencyAmount"`
							CurrencyCode   string `xml:"CurrencyCode"`
						} `xml:"FeeAmount"`
						FeeType string `xml:"FeeType"`
					} `xml:"FeeComponent"`
				} `xml:"ItemFeeAdjustmentList"`
				OrderAdjustmentItemID    string `xml:"OrderAdjustmentItemId"`
				QuantityShipped          string `xml:"QuantityShipped"`
				ItemChargeAdjustmentList struct {
					Text            string `xml:",chardata"`
					ChargeComponent []struct {
						Text         string `xml:",chardata"`
						ChargeType   string `xml:"ChargeType"`
						ChargeAmount struct {
							Text           string `xml:",chardata"`
							CurrencyAmount string `xml:"CurrencyAmount"`
							CurrencyCode   string `xml:"CurrencyCode"`
						} `xml:"ChargeAmount"`
					} `xml:"ChargeComponent"`
				} `xml:"ItemChargeAdjustmentList"`
				SellerSKU string `xml:"SellerSKU"`
			} `xml:"ShipmentItem"`
		} `xml:"ShipmentItemAdjustmentList"`
		MarketplaceName string `xml:"MarketplaceName"`
		SellerOrderID   string `xml:"SellerOrderId"`
	} `xml:"ShipmentEvent"`
}

// AdjustmentEventList ...
type AdjustmentEventList struct {
	Text            string `xml:",chardata"`
	AdjustmentEvent []struct {
		Text               string `xml:",chardata"`
		AdjustmentType     string `xml:"AdjustmentType"`
		AdjustmentItemList struct {
			Text           string `xml:",chardata"`
			AdjustmentItem struct {
				Text          string `xml:",chardata"`
				PerUnitAmount struct {
					Text           string `xml:",chardata"`
					CurrencyAmount string `xml:"CurrencyAmount"`
					CurrencyCode   string `xml:"CurrencyCode"`
				} `xml:"PerUnitAmount"`
				TotalAmount struct {
					Text           string `xml:",chardata"`
					CurrencyAmount string `xml:"CurrencyAmount"`
					CurrencyCode   string `xml:"CurrencyCode"`
				} `xml:"TotalAmount"`
				Quantity           string `xml:"Quantity"`
				SellerSKU          string `xml:"SellerSKU"`
				ProductDescription string `xml:"ProductDescription"`
			} `xml:"AdjustmentItem"`
		} `xml:"AdjustmentItemList"`
		AdjustmentAmount struct {
			Text           string `xml:",chardata"`
			CurrencyAmount string `xml:"CurrencyAmount"`
			CurrencyCode   string `xml:"CurrencyCode"`
		} `xml:"AdjustmentAmount"`
		PostedDate string `xml:"PostedDate"`
	} `xml:"AdjustmentEvent"`
}
