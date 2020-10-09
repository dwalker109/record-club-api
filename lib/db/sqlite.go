package db

import (
	"github.com/dwalker109/record-club-api/lib/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SQLiteConn *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect db")
	}

	if err := db.AutoMigrate(&model.Pick{}); err != nil {
		panic("pick migrations failed")
	}

	SQLiteConn = db
}
