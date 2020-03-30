package article

import (
	"compose/article/articleCommons"
	"compose/article/create"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	articleCommons.SetDB(db)
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/create", create.Handler).Methods(http.MethodPost)
}
