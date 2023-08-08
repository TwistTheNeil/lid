package cmd

import (
	"encoding/json"
	"fmt"
	"lid/services/logger"
	"os"

	"github.com/spf13/cobra"
)

// deviceAddCmd represents the deviceAdd command
var dump = &cobra.Command{
	Use:   "dump",
	Short: "Dump all info",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("dump")
		log.Trace("dump called")

		l, err := hrs.FindAllPreload()
		if err != nil {
			log.Error("something went wrong fetching data", err)
			os.Exit(1)
		}

		b, _ := json.MarshalIndent(l, "", "  ")
		fmt.Printf("%v", string(b))
	},
}

func init() {
	rootCmd.AddCommand(dump)
}
