package cmd

import (
	"lid/interfaces"
	"lid/repositories"
	sqlite3db "lid/services/db"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var debugLevel int8
var database string
var drs interfaces.DeviceRepository
var frs interfaces.FileRepository

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lid",
	Short: "Utility to help associate files with running drives and backup drives",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		switch debugLevel {
		case -1, 0, 1, 2, 3, 4, 5, 6, 7:
			zerolog.SetGlobalLevel(zerolog.Level(debugLevel))
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}

		sqlite3db.Open(database)
		db := sqlite3db.Get()
		drs = repositories.NewGORMDeviceRepositoryService(db.DB)
		frs = repositories.NewGORMNodeRepositoryService(db.DB)
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
	rootCmd.PersistentFlags().Int8VarP(&debugLevel, "debuglevel", "d", 1, "debug level")
	rootCmd.PersistentFlags().StringVar(&database, "db", "", "database (sqlite) file")
	rootCmd.MarkPersistentFlagRequired("db")
}
