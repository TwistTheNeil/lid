/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"lid/models"
	"lid/services/logger"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// dbInitCmd represents the dbInit command
var dbInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a sqlite3 database",
	Run: func(cmd *cobra.Command, args []string) {
		log := logger.CreateLogger("db_init.run")

		if len(args) == 0 {
			log.Error("No database file path argument provided", fmt.Errorf("no database file path provided"))
			os.Exit(1)
		}

		if len(args) > 1 {
			log.Warn("Using only first argument to create database file")
		}

		// TODO: check if db exists
		f, err := os.Create(args[0])
		if err != nil {
			panic(err)
		}
		defer f.Close()

		db, err := gorm.Open(sqlite.Open(args[0]), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db.AutoMigrate(&models.Device{}, &models.Node{})
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}

		sqlDB.Close()

		log.Info("Database initialized at " + args[0])
	},
}

func init() {
	dbCmd.AddCommand(dbInitCmd)
}
