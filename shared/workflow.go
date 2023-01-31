package shared

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
	"time"
)

type OrderDetails struct {
	CurrentState OrderStatus
	StateMsg     string
	OrderID      string
	ProductCode  string
}

func OrderingWorkflow(ctx workflow.Context, od *OrderDetails) (*OrderDetails, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	err := workflow.SetQueryHandler(ctx, "current_state", func() (string, error) {
		return fmt.Sprintf("%s: %s", od.CurrentState, od.StateMsg), nil
	})
	if err != nil {
		od.CurrentState = Error
		od.StateMsg = "Could not register query handler"
		return od, err
	}

	var chargeResult string
	err = workflow.ExecuteActivity(ctx, ChargeCustomer).Get(ctx, &chargeResult)
	if err != nil {
		od.CurrentState = Error
		od.StateMsg = "Could not charge customer"
		return od, err
	}

	od.CurrentState = Paid
	od.StateMsg = "Order has been paid, waiting for pickup"

	var pickupSig OrderPickupSig
	sigChan := workflow.GetSignalChannel(ctx, "order-pickup-signal")
	sigChan.Receive(ctx, &pickupSig)

	od.CurrentState = PickedUp
	od.StateMsg = fmt.Sprintf("The order has been picked up by %s", pickupSig.DriverID)

	var deliverSig OrderDeliverSig
	deliverChan := workflow.GetSignalChannel(ctx, "order-delivery-signal")
	ok, _ := deliverChan.ReceiveWithTimeout(ctx, time.Minute, &deliverSig)
	if !ok {
		od.CurrentState = Refunding
		od.StateMsg = "Order was not completed in time, refunding"

		var refundResult string
		err := workflow.ExecuteActivity(ctx, RefundCustomer).Get(ctx, &refundResult)
		if err != nil {
			od.CurrentState = Error
			od.StateMsg = "Could not charge customer"
			return od, err
		}
		od.CurrentState = Refunded
		od.StateMsg = "The order was not delivered and payment was refunded"
		return od, nil
	}
	od.CurrentState = Delivered
	od.StateMsg = "The order has been delivered"

	workflow.Sleep(ctx, time.Minute)
	return od, nil
}
