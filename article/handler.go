package article

import (
	"compose/article/articleCommons"
	"compose/article/articleDetails"
	"compose/article/create"
	"compose/serviceContracts"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	articleCommons.Database = db
}

func SetServiceContractImpl(userContract serviceContracts.UserServiceContract) {
	articleCommons.UserServiceContract = userContract
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/create", create.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/{article_id}", articleDetails.Handler).Methods(http.MethodGet)
}
