package article

import (
	"compose/article/articleCommons"
	"compose/article/articleDetails"
	"compose/article/create"
	"compose/article/delete"
	"compose/article/update"
	"compose/serviceContracts"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	articleCommons.Database = db
}

func SetServiceContractImpl(userContract serviceContracts.UserServiceContract,
	commentServiceContract serviceContracts.CommentServiceContract,
	likeServiceContract serviceContracts.LikeServiceContract) {
	articleCommons.UserServiceContract = userContract
	articleCommons.CommentServiceContract = commentServiceContract
	articleCommons.LikeServiceContract = likeServiceContract
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/create", create.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/get/{article_id}", articleDetails.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/update", update.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/delete", delete.Handler).Methods(http.MethodPost)
}
