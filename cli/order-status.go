package cli

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var orderStatusCmd = &cobra.Command{
	Use: "order-status",
	Run: orderStatusFunc,
}

func init() {
	rootCmd.AddCommand(orderStatusCmd)
}

func orderStatusFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide an order ID")
		os.Exit(1)
	}

	resp, err := temporalClient.QueryWorkflow(cmd.Context(), args[0], "", "current_state")
	if err != nil {
		logrus.Fatal("Could not query workflow: ", err)
	}
	var respStr string
	err = resp.Get(&respStr)
	if err != nil {
		logrus.Fatal("Could not read resp: ", err)
	}
	fmt.Println("Current status for ", args[0], ": ", respStr)
}
