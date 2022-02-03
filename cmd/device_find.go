/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"lid/models"
	"lid/services/logger"
	"os"

	"github.com/spf13/cobra"
)

var query string
var queryFunc func(string) (models.Device, error)

// deviceAddCmd represents the deviceAdd command
var deviceFindCmd = &cobra.Command{
	Use:   "find",
	Short: "Add a device",
	PreRun: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("device_prerun_find")
		var err error

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Error("something went wrong when fetching --name arg", err)
			os.Exit(1)
		}

		uuid, err := cmd.Flags().GetString("uuid")
		if err != nil {
			log.Error("something went wrong when fetching --uuid arg", err)
			os.Exit(1)
		}

		if name != "" && uuid != "" {
			onlyOneArgError := "only --name or --uuid allowed"
			log.Error(onlyOneArgError, fmt.Errorf(onlyOneArgError))
			os.Exit(1)
		}

		if name == "" && uuid == "" {
			needArgError := "need --name or --uuid args"
			log.Error(needArgError, fmt.Errorf(needArgError))
			os.Exit(1)
		}

		if name != "" {
			query = name
			queryFunc = drs.FindByName
		}

		if uuid != "" {
			query = uuid
			queryFunc = drs.FindByUUID
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("device_find")
		log.Trace("device_find called")

		d, err := queryFunc(query)
		if err != nil {
			log.Error("something went wrong fetching device details", err)
			os.Exit(1)
		}

		fmt.Print(d)
	},
}

func init() {
	deviceCmd.AddCommand(deviceFindCmd)
	deviceFindCmd.Flags().String("name", "", "device name")
	deviceFindCmd.Flags().String("uuid", "", "device uuid")
}
