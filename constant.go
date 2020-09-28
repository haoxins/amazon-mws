package mws

const (
	// ReportTypeSettlement ...
	ReportTypeSettlement = "_GET_V2_SETTLEMENT_REPORT_DATA_FLAT_FILE_V2_"
	// ReportsPath The reports path
	ReportsPath = "/Reports/2009-01-01"
	// OrdersPath The orders path
	OrdersPath = "/Orders/2013-09-01"
	// FeedsPath The feeds path
	FeedsPath = "/Feeds/2009-01-01"
	// InboundShipmentPath The inbound shipment path
	InboundShipmentPath = "/FulfillmentInboundShipment/2010-10-01"
	// FinancesPath The finances path
	FinancesPath = "/Finances/2015-05-01"
)

// https://docs.developer.amazonservices.com/en_US/dev_guide/DG_Endpoints.html

// MarketplaceID The MWS aarketplace id
var MarketplaceID = map[string]string{
	"BR": "A2Q3Y263D00KWC",
	"CA": "A2EUQ1WTGCTBG2",
	"MX": "A1AM78C64UM0Y8",
	"US": "ATVPDKIKX0DER",
	"AE": "A2VIGQ35RCS4UG",
	"DE": "A1PA6795UKMFR9",
	"EG": "ARBP9OOSHTCHU",
	"ES": "A1RKKUPIHCS9HS",
	"FR": "A13V1IB3VIYZZH",
	"GB": "A1F83G8C2ARO7P",
	"IN": "A21TJRUUN4KGV",
	"IT": "APJ6JRA9NG5V4",
	"NL": "A1805IZSGTT6HS",
	"SA": "A17E79C6D8DWNP",
	"TR": "A33AVAJ2PDY3EV",
	"SG": "A19VAU5U5O7RUS",
	"AU": "A39IBJ37TRP1C6",
	"JP": "A1VC38T7YXB528",
}

// Endpoint The MWS endpoint
var Endpoint = map[string]string{
	"BR": "https://mws.amazonservices.com",
	"CA": "https://mws.amazonservices.ca",
	"MX": "https://mws.amazonservices.com.mx",
	"US": "https://mws.amazonservices.com",
	"AE": "https://mws.amazonservices.ae",
	"DE": "https://mws-eu.amazonservices.com",
	"EG": "https://mws-eu.amazonservices.com",
	"ES": "https://mws-eu.amazonservices.com",
	"FR": "https://mws-eu.amazonservices.com",
	"GB": "https://mws-eu.amazonservices.com",
	"IN": "https://mws.amazonservices.in",
	"IT": "https://mws-eu.amazonservices.com",
	"NL": "https://mws-eu.amazonservices.com",
	"SA": "https://mws-eu.amazonservices.com",
	"TR": "https://mws-eu.amazonservices.com",
	"SG": "https://mws-fe.amazonservices.com",
	"AU": "https://mws.amazonservices.com.au",
	"JP": "https://mws.amazonservices.jp",
}
