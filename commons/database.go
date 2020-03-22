package commons

import (
	"github.com/jinzhu/gorm"
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
