package main

import (
	"github.com/Zyian/temporal-lab/cli"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
	cli.ExecuteOrder()
}
