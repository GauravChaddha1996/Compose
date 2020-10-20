package article

import (
	articleDetails2 "compose/endpoints/article/articleDetails"
	create2 "compose/endpoints/article/create"
	delete2 "compose/endpoints/article/delete"
	update2 "compose/endpoints/article/update"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/create", create2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/get/{article_id}", articleDetails2.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/update", update2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/delete", delete2.Handler).Methods(http.MethodPost)
}
