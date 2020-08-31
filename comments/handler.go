package comments

import (
	"compose/comments/articleComments"
	"compose/comments/commentCommons"
	"compose/comments/createComment"
	"compose/comments/createReply"
	"compose/comments/replyThread"
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
	subRouter.HandleFunc("/comment", createComment.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/replyThread", replyThread.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/reply", createReply.Handler).Methods(http.MethodPost)
}
