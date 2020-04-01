package user

import (
	"compose/commons"
	"compose/user/delete"
	"compose/user/login"
	"compose/user/signup"
	"compose/user/update"
	"compose/user/userCommons"
	"compose/user/userDetails"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	userCommons.Database = db
	for path, config := range getSecurityMiddlewareConfigMap() {
		commons.AddSecurityMiddlewarePathConfig("/user"+path, config)
	}
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/signup", signup.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/login", login.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/{user_id}", userDetails.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/update", update.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/delete", delete.Handler).Methods(http.MethodPost)
}

func getSecurityMiddlewareConfigMap() map[string]*commons.SecurityMiddlewarePathConfig {
	return map[string]*commons.SecurityMiddlewarePathConfig{
		"/signup": {
			CheckAccessToken: false,
			CheckUserId:      false,
			CheckUserEmail:   false,
		}, "/login": {
			CheckAccessToken: false,
			CheckUserId:      false,
			CheckUserEmail:   false,
		},
	}
}
