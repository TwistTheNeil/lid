package cmd

import (
	"encoding/json"
	"fmt"
	"lid/services/logger"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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

		if j, _ := cmd.Flags().GetBool("json"); j {
			b, _ := json.MarshalIndent(l, "", "  ")
			fmt.Printf("%v", string(b))
			return
		}

		t := table.NewWriter()
		t.AppendHeader(table.Row{
			"UUID",
			"Name",
			"Size",
		})

		for _, v := range l {
			t.AppendRow(table.Row{
				v.UUID,
				v.Name,
				v.Size,
			})
		}

		fmt.Println(t.Render())
	},
}

func init() {
	deviceCmd.AddCommand(deviceListCmd)
	deviceListCmd.Flags().Bool("json", false, "print list as json")
}
