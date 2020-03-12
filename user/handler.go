package user

import (
	"compose/user/signup"
	"compose/user/userCommons"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Init(db *gorm.DB) {
	userCommons.SetDB(db)
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/signup", signup.Handler)
}
