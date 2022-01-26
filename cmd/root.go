package cmd

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var debugLevel int16

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lid",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		switch debugLevel {
		case -1:
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		case 0:
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case 1:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case 2:
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case 3:
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().Int16VarP(&debugLevel, "debuglevel", "d", 1, "debug level")
}
