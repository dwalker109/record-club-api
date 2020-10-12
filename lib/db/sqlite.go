package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SQLiteConn *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	SQLiteConn = db
}
