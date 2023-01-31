package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "tempdel",
	Short: "Tempdel a temporal based food delivery CLI",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hi, bye!")
	},
}

func ExecuteOrder() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ExecuteDelivery() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(namespaceRegisterCmd)
}
