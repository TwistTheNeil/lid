package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqlite3DB struct {
	DB *gorm.DB
}

var db sqlite3DB

func Open(dsn string) {
	sqliteDB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db = sqlite3DB{DB: sqliteDB}
}

func Get() sqlite3DB {
	return db
}

func Close() error {
	dbCtx, err := db.DB.DB()
	if err != nil {
		panic(err)
	}
	dbCtx.Close()
	return nil
}
