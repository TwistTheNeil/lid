package cmd

import (
	"lid/models"
	"lid/services/disk_inventory"
	"lid/services/logger"
	"os"

	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze a set of directories and/or files",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("analyze")
		var nl models.NodeList

		log.Trace("analyze file called")

		tabulate, err := cmd.Flags().GetBool("tabulate")
		if err != nil {
			log.Error("Can't read --tabulate value", err)
		}

		files := disk_inventory.BuildFileList(args...)

		for _, filename := range files {
			fileStat, err := os.Stat(filename)
			if err != nil {
				log.Error("Can't read file", err, "file", filename)
				continue
			}

			newNode := models.Node{Name: filename, Size: fileStat.Size()}
			newNode.Hash()
			nl.Append(newNode)
		}

		if tabulate {
			nl.Pretty()
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	rootCmd.PersistentFlags().BoolP("tabulate", "t", false, "Tabulate analysis")
}
