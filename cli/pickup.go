package cli

import (
	"fmt"
	"github.com/Zyian/temporal-lab/shared"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var orderPickupCmd = &cobra.Command{
	Use: "pickup",
	Run: orderPickupFunc,
}

func init() {
	rootCmd.AddCommand(orderPickupCmd)
}

func orderPickupFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide an order ID")
		os.Exit(1)
	}
	uid, _ := uuid.NewUUID()
	ops := shared.OrderPickupSig{DriverID: uid.String()}

	err := temporalClient.SignalWorkflow(cmd.Context(), args[0], "", "order-pickup-signal", ops)
	if err != nil {
		logrus.Fatal("Could not pick up order: ", err)
	}
}
