/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lid/services/logger"
	"os"

	"github.com/spf13/cobra"
)

// deviceAddCmd represents the deviceAdd command
var deviceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a device",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("device_add")
		log.Trace("device_add called")

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Error("something went wrong specifying device --name", err)
			os.Exit(1)
		}

		uuid, err := cmd.Flags().GetString("uuid")
		if err != nil {
			log.Error("something went wrong specifying device --uuid", err)
			os.Exit(1)
		}

		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Error("something went wrong specifying device --host", err)
			os.Exit(1)
		}

		size, err := cmd.Flags().GetInt64("size")
		if err != nil {
			log.Error("something went wrong specifying device --size", err)
			os.Exit(1)
		}

		err = drs.Create(name, uuid, host, size)
		if err != nil {
			log.Error("something went wrong saving device details", err)
		}
	},
}

func init() {
	deviceCmd.AddCommand(deviceAddCmd)
	deviceAddCmd.Flags().String("name", "", "device name")
	deviceAddCmd.Flags().String("uuid", "", "device uuid")
	deviceAddCmd.Flags().String("host", "unnamed", "host name")
	deviceAddCmd.Flags().Int64("size", 0, "device size")
	deviceAddCmd.MarkFlagRequired("name")
	deviceAddCmd.MarkFlagRequired("uuid")
	deviceAddCmd.MarkFlagRequired("size")
}
