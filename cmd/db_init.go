/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lid/models"
	"lid/services/logger"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// dbInitCmd represents the dbInit command
var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a sqlite3 database",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		switch debugLevel {
		case -1, 0, 1, 2, 3, 4, 5, 6, 7:
			zerolog.SetGlobalLevel(zerolog.Level(debugLevel))
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("db_init.run")

		// TODO: check if db exists
		f, err := os.Create(database)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		db, err := gorm.Open(sqlite.Open(database), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&models.Device{}, &models.File{})
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		sqlDB.Close()

		log.Info("Database initialized at " + database)
	},
}

func init() {
	dbCmd.AddCommand(dbInitCmd)
}
