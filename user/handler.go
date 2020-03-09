package user

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func Init(db *gorm.DB) {
	Db = db
}

func AddSubRoutes(subRouter *mux.Router)  {
	subRouter.HandleFunc("/signup", SignupHandler)
}
