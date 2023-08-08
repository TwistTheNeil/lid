/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deviceCmd represents the device command
var hostCmd = &cobra.Command{
	Use: "host",
}

func init() {
	rootCmd.AddCommand(hostCmd)
}
