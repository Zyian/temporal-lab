package cli

import (
	"fmt"
	"github.com/Zyian/temporal-lab/shared"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"os"
)

var productMap = map[string]string{
	"28e0": "Apples",
	"48b5": "Oranges",
	"b99e": "Chipotle",
	"1fff": "McDonalds",
	"4947": "Wendys",
	"bf5a": "Pizza",
}

var orderCmd = &cobra.Command{
	Use: "order",
	Run: orderFunc,
}

func init() {
	rootCmd.AddCommand(orderCmd)
	orderCmd.PersistentFlags().BoolP("list", "l", false, "List all available products")
	_ = viper.BindPFlag("useList", orderCmd.PersistentFlags().Lookup("list"))
}

func orderFunc(cmd *cobra.Command, args []string) {
	if viper.GetBool("useList") {
		fmt.Println("Here's what you can order:")
		for k, v := range productMap {
			fmt.Printf("%s:\t%v\n", k, v)
		}
		os.Exit(0)
	}

	if len(args) < 1 {
		fmt.Println("Not enough arguments provided, add a product")
		os.Exit(1)
	}

	logrus.WithFields(logrus.Fields{
		"server-hostport": viper.GetString("hostport"),
		"server-name":     viper.GetString("servername"),
		"namespace":       viper.GetString("namespace"),
	}).Info("Submitting Order")

	oid, _ := uuid.NewUUID()
	o := &shared.OrderDetails{
		CurrentState: shared.ChargingCard,
		OrderID:      oid.String(),
		ProductCode:  args[0],
	}

	workflowOpts := client.StartWorkflowOptions{
		ID:        oid.String(),
		TaskQueue: shared.OrderProductTaskQueue,
	}
	_, err := temporalClient.ExecuteWorkflow(cmd.Context(), workflowOpts, shared.OrderingWorkflow, o)
	if err != nil {
		fmt.Println("Could not execute order: ", err)
		os.Exit(1)
	}
	fmt.Println("Created Order #: ", oid.String())
}
