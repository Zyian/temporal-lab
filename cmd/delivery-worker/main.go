package main

import (
	"crypto/tls"
	"fmt"
	"github.com/Zyian/temporal-lab/shared"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"os"
	"strings"
)

func init() {
	viper.SetEnvPrefix("temp")
	_ = viper.BindEnv("hostport")
	_ = viper.BindEnv("servername")
	_ = viper.BindEnv("namespace")

	switch strings.ToLower(os.Getenv("LEVEL")) {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	}
}

func main() {
	logrus.WithFields(logrus.Fields{
		"server-hostport": viper.GetString("hostport"),
		"server-name":     viper.GetString("servername"),
		"namespace":       viper.GetString("namespace"),
	}).Info("Loaded worker")
	temporalClient, err := client.Dial(client.Options{
		Namespace: viper.GetString("namespace"),
		HostPort:  viper.GetString("hostport"),
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{
				ServerName: viper.GetString("servername"),
			},
		},
	})
	if err != nil {
		fmt.Println("failed to connect to Temporal server: ", err)
		os.Exit(1)
	}
	defer temporalClient.Close()

	w := worker.New(temporalClient, shared.OrderProductTaskQueue, worker.Options{})

	w.RegisterWorkflow(shared.OrderingWorkflow)
	w.RegisterActivity(shared.ChargeCustomer)
	w.RegisterActivity(shared.RefundCustomer)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		fmt.Println("failed to run Temporal worker")
		os.Exit(1)
	}
}
