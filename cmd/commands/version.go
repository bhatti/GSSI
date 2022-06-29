package commands

import (
	"fmt"
	"github.com/bhatti/GSSI/internal/version"

	"github.com/spf13/cobra"
)

var (
	shortened = false
	Version   = "dev"
	Commit    = "dirty"
	Date      = ""
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version",
	Long:  `Version will output the current build information`,
	Run: func(cmd *cobra.Command, _ []string) {
		v := version.New(Version, Commit, Date)
		var response string

		if shortened {
			response = v.ToYAML()
		} else {
			response = v.ToJSON()
		}

		fmt.Printf("%+v", response)
		return
	},
}

func init() {
	versionCmd.Flags().BoolVarP(
		&shortened,
		"short",
		"s",
		true,
		"Use shortened output for version information.")
	rootCmd.AddCommand(versionCmd)
}
