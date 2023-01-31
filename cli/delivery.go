package cli

import (
	"fmt"
	"github.com/Zyian/temporal-lab/shared"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var orderDeliverCmd = &cobra.Command{
	Use: "deliver",
	Run: orderDeliverFunc,
}

func init() {
	rootCmd.AddCommand(orderDeliverCmd)
}

func orderDeliverFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide an order ID")
		os.Exit(1)
	}
	ops := shared.OrderDeliverSig{DeliveredTime: time.Now()}

	err := temporalClient.SignalWorkflow(cmd.Context(), args[0], "", "order-delivery-signal", ops)
	if err != nil {
		logrus.Fatal("Could not deliver order: ", err)
	}
}
