package db

import (
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Activate(db *gorm.DB) {
	Conn = db
}
