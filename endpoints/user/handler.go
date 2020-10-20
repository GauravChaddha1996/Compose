package user

import (
	"compose/commons"
	delete2 "compose/endpoints/user/delete"
	likedArticles2 "compose/endpoints/user/likedArticles"
	login2 "compose/endpoints/user/login"
	logout2 "compose/endpoints/user/logout"
	postedArticles2 "compose/endpoints/user/postedArticles"
	signup2 "compose/endpoints/user/signup"
	update2 "compose/endpoints/user/update"
	userDetails2 "compose/endpoints/user/userDetails"
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
	subRouter.HandleFunc("/signup", signup2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/login", login2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/logout", logout2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/update", update2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/delete", delete2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/postedArticles", postedArticles2.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/likedArticles", likedArticles2.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/{user_id}", userDetails2.Handler).Methods(http.MethodGet)
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
