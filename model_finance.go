package mws

import "encoding/xml"

type ListFinancialEventsResponse struct {
	XMLName                   xml.Name `xml:"ListFinancialEventsResponse"`
	Text                      string   `xml:",chardata"`
	Xmlns                     string   `xml:"xmlns,attr"`
	ListFinancialEventsResult struct {
		Text            string `xml:",chardata"`
		FinancialEvents struct {
			Text                       string `xml:",chardata"`
			ProductAdsPaymentEventList []struct {
				Text                   string `xml:",chardata"`
				ProductAdsPaymentEvent struct {
					Text            string `xml:",chardata"`
					PostedDate      string `xml:"PostedDate"`
					TransactionType string `xml:"transactionType"`
					InvoiceId       string `xml:"invoiceId"`
					BaseValue       struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"baseValue"`
					TaxValue struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"taxValue"`
					TransactionValue struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"transactionValue"`
				} `xml:"ProductAdsPaymentEvent"`
			} `xml:"ProductAdsPaymentEventList"`
			RentalTransactionEventList     string `xml:"RentalTransactionEventList"`
			PayWithAmazonEventList         string `xml:"PayWithAmazonEventList"`
			ServiceFeeEventList            string `xml:"ServiceFeeEventList"`
			ServiceProviderCreditEventList string `xml:"ServiceProviderCreditEventList"`
			SellerDealPaymentEventList     struct {
				Text                   string `xml:",chardata"`
				SellerDealPaymentEvent struct {
					Text            string `xml:",chardata"`
					PostedDate      string `xml:"PostedDate"`
					DealDescription string `xml:"DealDescription"`
					DealId          string `xml:"DealId"`
					EventType       string `xml:"EventType"`
					FeeType         string `xml:"FeeType"`
					FeeAmount       struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"FeeAmount"`
					TaxAmount struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxAmount"`
				} `xml:"SellerDealPaymentEvent"`
			} `xml:"SellerDealPaymentEventList"`
			DebtRecoveryEventList string `xml:"DebtRecoveryEventList"`
			ShipmentEventList     struct {
				Text          string `xml:",chardata"`
				ShipmentEvent struct {
					Text             string `xml:",chardata"`
					ShipmentItemList struct {
						Text         string `xml:",chardata"`
						ShipmentItem struct {
							Text           string `xml:",chardata"`
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
									FeeType   string `xml:"FeeType"`
									FeeAmount struct {
										Text           string `xml:",chardata"`
										CurrencyAmount string `xml:"CurrencyAmount"`
										CurrencyCode   string `xml:"CurrencyCode"`
									} `xml:"FeeAmount"`
								} `xml:"FeeComponent"`
							} `xml:"ItemFeeList"`
							OrderItemId     string `xml:"OrderItemId"`
							QuantityShipped string `xml:"QuantityShipped"`
							SellerSKU       string `xml:"SellerSKU"`
						} `xml:"ShipmentItem"`
					} `xml:"ShipmentItemList"`
					AmazonOrderId   string `xml:"AmazonOrderId"`
					PostedDate      string `xml:"PostedDate"`
					MarketplaceName string `xml:"MarketplaceName"`
					SellerOrderId   string `xml:"SellerOrderId"`
				} `xml:"ShipmentEvent"`
			} `xml:"ShipmentEventList"`
			AffordabilityExpenseEventList struct {
				Text                      string `xml:",chardata"`
				AffordabilityExpenseEvent struct {
					Text            string `xml:",chardata"`
					PostedDate      string `xml:"PostedDate"`
					TransactionType string `xml:"TransactionType"`
					AmazonOrderId   string `xml:"AmazonOrderId"`
					MarketplaceId   string `xml:"MarketplaceId"`
					BaseExpense     struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"BaseExpense"`
					TaxTypeIGST struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxTypeIGST"`
					TaxTypeSGST struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxTypeSGST"`
					TaxTypeCGST struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxTypeCGST"`
					TotalExpense struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TotalExpense"`
				} `xml:"AffordabilityExpenseEvent"`
			} `xml:"AffordabilityExpenseEventList"`
			RetrochargeEventList                  string `xml:"RetrochargeEventList"`
			GuaranteeClaimEventList               string `xml:"GuaranteeClaimEventList"`
			ChargebackEventList                   string `xml:"ChargebackEventList"`
			LoanServicingEventList                string `xml:"LoanServicingEventList"`
			RefundEventList                       string `xml:"RefundEventList"`
			AdjustmentEventList                   string `xml:"AdjustmentEventList"`
			PerformanceBondRefundEventList        string `xml:"PerformanceBondRefundEventList"`
			AffordabilityExpenseReversalEventList struct {
				Text                              string `xml:",chardata"`
				AffordabilityExpenseReversalEvent struct {
					Text            string `xml:",chardata"`
					PostedDate      string `xml:"PostedDate"`
					TransactionType string `xml:"TransactionType"`
					AmazonOrderId   string `xml:"AmazonOrderId"`
					MarketplaceId   string `xml:"MarketplaceId"`
					BaseExpense     struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"BaseExpense"`
					TaxTypeIGST struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxTypeIGST"`
					TaxTypeSGST struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxTypeSGST"`
					TaxTypeCGST struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TaxTypeCGST"`
					TotalExpense struct {
						Text           string `xml:",chardata"`
						CurrencyAmount string `xml:"CurrencyAmount"`
						CurrencyCode   string `xml:"CurrencyCode"`
					} `xml:"TotalExpense"`
				} `xml:"AffordabilityExpenseReversalEvent"`
			} `xml:"AffordabilityExpenseReversalEventList"`
			TDSReimbursementEventList struct {
				Text                  string `xml:",chardata"`
				TDSReimbursementEvent struct {
					Text             string `xml:",chardata"`
					ReimbursedAmount struct {
						Text           string `xml:",chardata"`
						CurrencyCode   string `xml:"CurrencyCode"`
						CurrencyAmount string `xml:"CurrencyAmount"`
					} `xml:"ReimbursedAmount"`
					PostedDate string `xml:"PostedDate"`
					TdsOrderId string `xml:"TdsOrderId"`
				} `xml:"TDSReimbursementEvent"`
			} `xml:"TDSReimbursementEventList"`
		} `xml:"FinancialEvents"`
	} `xml:"ListFinancialEventsResult"`
	ResponseMetadata struct {
		Text      string `xml:",chardata"`
		RequestId string `xml:"RequestId"`
	} `xml:"ResponseMetadata"`
}
