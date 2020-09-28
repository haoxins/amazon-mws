package mws

import (
	"fmt"
	"net/url"
	"time"

	"github.com/pkg4go/tools"
	"github.com/spf13/cast"
)

type shipmentStatus string

const (
	shipmentStatusWorking   shipmentStatus = "WORKING"
	shipmentStatusShipped                  = "SHIPPED"
	shipmentStatusInTransit                = "IN_TRANSIT"
	shipmentStatusDelivered                = "DELIVERED"
	shipmentStatusCheckedIn                = "CHECKED_IN"
	shipmentStatusReceiving                = "RECEIVING"
	shipmentStatusClosed                   = "CLOSED"
	shipmentStatusCancelled                = "CANCELLED"
	shipmentStatusDeleted                  = "DELETED"
	shipmentStatusError                    = "ERROR"
)

// ListInboundShipmentsParams ListInboundShipments API params
type ListInboundShipmentsParams struct {
	ShipmentStatusList []shipmentStatus
	LastUpdatedAfter   time.Time
	LastUpdatedBefore  time.Time
}

// ListInboundShipments ...
func (seller *Seller) ListInboundShipments(params ListInboundShipmentsParams) {
	opts := seller.genListInboundShipmentsParams(params)

	body, err := seller.requestInboundShipment(opts)
	tools.AssertError(err)

	// TODO
	fmt.Println(string(body))
}

func (seller *Seller) genListInboundShipmentsParams(params ListInboundShipmentsParams) string {
	v := url.Values{}

	v.Add("Action", "ListInboundShipments")
	v.Add("SellerId", seller.SellerID)
	v.Add("MWSAuthToken", seller.AuthToken)
	v.Add("AWSAccessKeyId", seller.AccessKey)
	v.Add("LastUpdatedAfter", params.LastUpdatedAfter.Format(time.RFC3339))
	v.Add("LastUpdatedBefore", params.LastUpdatedBefore.Format(time.RFC3339))
	v.Add("Timestamp", time.Now().UTC().Format(time.RFC3339))
	v.Add("SignatureVersion", "2")
	v.Add("SignatureMethod", "HmacSHA256")
	v.Add("Version", "2010-10-01")

	for key, val := range params.ShipmentStatusList {
		v.Add("ShipmentStatusList.member."+cast.ToString(key+1), string(val))
	}

	s := v.Encode()

	return s
}

func (seller *Seller) requestInboundShipment(params string) ([]byte, error) {
	return seller.get(InboundShipmentPath, params)
}
