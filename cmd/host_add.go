/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lid/services/logger"

	"github.com/spf13/cobra"
)

// hostAddCmd represents the hostAdd command
var hostAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a host",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("host_add")
		log.Trace("host_add called")

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal("something went wrong specifying host --name", err)
			return
		}

		err = hrs.Create(name)
		if err != nil {
			log.Fatal("something went wrong saving host details", err)
		}
	},
}

func init() {
	hostCmd.AddCommand(hostAddCmd)
	hostAddCmd.Flags().String("name", "", "host name")
	hostAddCmd.MarkFlagRequired("name")
}
