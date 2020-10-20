package user

import (
	"compose/commons"
	"compose/user/delete"
	"compose/user/likedArticles"
	"compose/user/login"
	"compose/user/logout"
	"compose/user/postedArticles"
	"compose/user/signup"
	"compose/user/update"
	"compose/user/userDetails"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func Init(*gorm.DB) {
	for path, config := range getEndpointSecurityConfigMap() {
		commons.AddEndpointSecurityConfig("/user"+path, config)
	}
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/signup", signup.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/login", login.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/logout", logout.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/update", update.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/delete", delete.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/postedArticles", postedArticles.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/likedArticles", likedArticles.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/{user_id}", userDetails.Handler).Methods(http.MethodGet)
}

func getEndpointSecurityConfigMap() map[string]*commons.EndpointSecurityConfig {
	return map[string]*commons.EndpointSecurityConfig{
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
