package cli

import (
	"crypto/tls"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"os"
)

var (
	temporalClient client.Client
)

func init() {
	cobra.OnInitialize(initClient)
}

func initClient() {
	logrus.WithFields(logrus.Fields{
		"server-hostport": viper.GetString("hostport"),
		"server-name":     viper.GetString("servername")},
	).Debug("Connecting client")

	var err error
	temporalClient, err = client.Dial(client.Options{
		Namespace: viper.GetString("namespace"),
		HostPort:  viper.GetString("hostport"),
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{
				ServerName: viper.GetString("servername"),
			},
		},
	})
	if err != nil {
		fmt.Println("Could not connect to Temporal Instance: ", err)
		os.Exit(1)
	}
}
