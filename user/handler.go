package user

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func Init(db *gorm.DB) {
	Db = db
	Db.AutoMigrate(User{})
}

func AddApiRoutes(router *mux.Router) {
	router.HandleFunc("/user/signup", SignupHandler)
}
