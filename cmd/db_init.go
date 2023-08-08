/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"lid/models"
	"lid/services/logger"
	"os"

	sqlite3db "lid/services/db"

	"github.com/spf13/cobra"
)

// dbInitCmd represents the dbInit command
var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a sqlite3 database",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// don't run the root PersistentPreRun for this command
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("db_init.run")

		// TODO: check if db exists
		f, err := os.Create(database)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		sqlite3db.Open(database)
		db := sqlite3db.Get()
		db.DB.AutoMigrate(&models.Device{}, &models.File{}, &models.Host{})
		sqlite3db.Close()

		log.Info("Database initialized at " + database)
	},
}

func init() {
	dbCmd.AddCommand(dbInitCmd)
}
