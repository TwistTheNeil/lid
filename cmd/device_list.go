package cmd

import (
	"fmt"
	"lid/services/logger"
	"os"

	"github.com/spf13/cobra"
)

// deviceAddCmd represents the deviceAdd command
var deviceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List storage devices",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("device_list")
		log.Trace("device_list called")

		l, err := drs.FindAll()
		if err != nil {
			log.Error("something went wrong fetching device list", err)
			os.Exit(1)
		}

		fmt.Printf("%+v", l)
	},
}

func init() {
	deviceCmd.AddCommand(deviceListCmd)
}
