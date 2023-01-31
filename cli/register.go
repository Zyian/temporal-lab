package cli

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.temporal.io/api/workflowservice/v1"
	"go.temporal.io/sdk/client"
	"os"
	"time"
)

var namespaceRegisterCmd = &cobra.Command{
	Use:   "namespace",
	Short: "",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires namespace arg")
		}
		return nil
	},
	Run: registerNamespace,
}

func registerNamespace(cmd *cobra.Command, args []string) {
	namespaceClient, err := client.NewNamespaceClient(client.Options{
		HostPort: viper.GetString("hostport"),
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{
				ServerName: viper.GetString("servername"),
			},
		},
	})
	if err != nil {
		fmt.Println("Could not build client: ", err)
		os.Exit(1)
	}

	logrus.WithFields(logrus.Fields{
		"server-hostport": viper.GetString("hostport"),
		"server-name":     viper.GetString("servername")},
	).Debug("Registering namespace")

	retention := time.Hour * 24 * 7
	err = namespaceClient.Register(cmd.Context(), &workflowservice.RegisterNamespaceRequest{
		Namespace:                        args[0],
		WorkflowExecutionRetentionPeriod: &retention,
	})
	if err != nil {
		fmt.Println("Could not register namespace: ", err)
		os.Exit(1)
	}
}
