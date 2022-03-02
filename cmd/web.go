package cmd

import (
	"lid/router"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web service",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: make config model
		router.Start(drs)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
