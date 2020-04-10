package comments

import (
	"compose/comments/articleComments"
	"compose/comments/commentCommons"
	"compose/serviceContracts"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	commentCommons.Database = db
}

func SetServiceContractImpl(articleContract serviceContracts.ArticleServiceContract, userContract serviceContracts.UserServiceContract) {
	commentCommons.ArticleServiceContract = articleContract
	commentCommons.UserServiceContract = userContract
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/comments", articleComments.Handler).Methods(http.MethodGet)
}
