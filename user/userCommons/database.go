package userCommons

import (
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func SetDB(db *gorm.DB) {
	database = db
}

func GetDB() *gorm.DB {
	return database
}