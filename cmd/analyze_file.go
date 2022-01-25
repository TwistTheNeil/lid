package cmd

import (
	"lid/services/file_ops"
	"lid/services/logger"

	"github.com/spf13/cobra"
)

// analyzeFileCmd represents the analyze command
var analyzeFileCmd = &cobra.Command{
	Use:   "file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("analyze")
		log.Trace("analyze file called")

		log.Info(file_ops.MD5Hash(args[0]))
	},
}

func init() {
	analyzeCmd.AddCommand(analyzeFileCmd)
}
