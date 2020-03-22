package user

import (
	"compose/user/delete"
	"compose/user/login"
	"compose/user/signup"
	"compose/user/userCommons"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	userCommons.SetDB(db)
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/signup", signup.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/login", login.Handler).Methods(http.MethodPost)
	//subRouter.HandleFunc("/{userId}", view.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/delete", delete.Handler).Methods(http.MethodPost)
}
