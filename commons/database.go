package commons

import (
	"gorm.io/gorm"
)

var database *gorm.DB

func Init(db *gorm.DB) {
	setDB(db)
}

func setDB(db *gorm.DB) {
	database = db
}

func GetDB() *gorm.DB {
	return database
}
