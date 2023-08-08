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
var hrs interfaces.HostRepository

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lid",
	Short: "Utility to help associate files with running drives and backup drives",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		sqlite3db.Open(database)
		db := sqlite3db.Get()
		drs = repositories.NewGORMDeviceRepositoryService(db.DB)
		frs = repositories.NewGORMNodeRepositoryService(db.DB)
		hrs = repositories.NewGORMHostRepositoryService(db.DB)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		sqlite3db.Close()
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
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	switch debugLevel {
	case -1, 0, 1, 2, 3, 4, 5, 6, 7:
		zerolog.SetGlobalLevel(zerolog.Level(debugLevel))
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	rootCmd.PersistentFlags().Int8VarP(&debugLevel, "debuglevel", "d", 1, "debug level")
	rootCmd.PersistentFlags().StringVar(&database, "db", "", "database (sqlite) file")
	rootCmd.MarkPersistentFlagRequired("db")
}
