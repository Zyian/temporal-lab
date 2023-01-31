package cli

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.temporal.io/api/filter/v1"
	"go.temporal.io/api/workflowservice/v1"
)

var orderListCmd = &cobra.Command{
	Use: "order-list",
	Run: orderListFunc,
}

func init() {
	rootCmd.AddCommand(orderListCmd)
}

func orderListFunc(cmd *cobra.Command, args []string) {
	list, err := temporalClient.ListOpenWorkflow(cmd.Context(), &workflowservice.ListOpenWorkflowExecutionsRequest{
		Namespace: viper.GetString("namespace"),
		Filters: &workflowservice.ListOpenWorkflowExecutionsRequest_TypeFilter{
			TypeFilter: &filter.WorkflowTypeFilter{
				Name: "OrderingWorkflow",
			}},
	})
	if err != nil {
		logrus.Fatal("Could not find list of workflows: ", err)
	}

	orders := map[string]string{}
	for _, w := range list.Executions {
		resp, err := temporalClient.QueryWorkflow(cmd.Context(), w.Execution.WorkflowId, "", "current_state")
		if err != nil {
			logrus.Fatal("Could not query workflow: ", err)
		}
		var respStr string
		err = resp.Get(&respStr)
		if err != nil {
			logrus.Fatal("Could not read resp: ", err)
		}
		orders[w.Execution.WorkflowId] = respStr
	}

	fmt.Println("Orders and Status\n=================")
	for o, s := range orders {
		fmt.Println(o, "\t", s)
	}
}
