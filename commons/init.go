package commons

import (
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	setDB(db)
	initTimeCoommons()
	initValidator()
}
