package cmd

import (
	"lid/router"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web service",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: make config model
		router.Start(drs, frs)
	},
}

func init() {
	// We catch signals to gracefully shutdown the server
	// giving control back to cobra for *PostRun actions
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		router.Stop()
	}()

	rootCmd.AddCommand(webCmd)
}
