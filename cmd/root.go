package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sendm",
	Short: "sendm is the root command",
	Long:  `Easy way to send mail`,
}

func Execute() {
	rootCmd.AddCommand(createSendCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
