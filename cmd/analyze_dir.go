package cmd

import (
	"lid/services/disk_inventory"
	"lid/services/logger"

	"github.com/spf13/cobra"
)

// analyzeDirCmd represents the analyze command
var analyzeDirCmd = &cobra.Command{
	Use:   "dir",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("analyze")
		log.Debug("analyze called")

		list := disk_inventory.List(args[0])
		list.Print()
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeDirCmd)
}
