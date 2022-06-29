package commands

import (
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client for verifiable credentials",
	Short: "client for verifiable credentials",
	Long:  "client for verifiable credentials",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	clientCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file")
	rootCmd.PersistentFlags().StringVar(
		&logFile,
		"logFile",
		"",
		"log file")
	rootCmd.AddCommand(clientCmd)
}
