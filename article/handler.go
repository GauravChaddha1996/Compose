package article

import (
	"compose/article/articleDetails"
	"compose/article/create"
	"compose/article/delete"
	"compose/article/update"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/create", create.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/get/{article_id}", articleDetails.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/update", update.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/delete", delete.Handler).Methods(http.MethodPost)
}
