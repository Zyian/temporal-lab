package shared

import "time"

const (
	OrderProductTaskQueue = "ORDER_PRODUCT_TASK_QUEUE"
)

type OrderStatus int

const (
	ChargingCard OrderStatus = iota
	Paid
	PickedUp
	Delivered
	Refunding
	Refunded
	Error
)

var orderStatusString = map[OrderStatus]string{
	ChargingCard: "Charging Card",
	Paid:         "Paid",
	PickedUp:     "Picked Up",
	Delivered:    "Delivered",
	Refunding:    "Refunding",
	Refunded:     "Refunded",
	Error:        "Error",
}

func (os OrderStatus) String() string {
	return orderStatusString[os]
}

type OrderPickupSig struct {
	DriverID string
}

type OrderDeliverSig struct {
	DeliveredTime time.Time
}
