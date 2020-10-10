package comments

import (
	"compose/comments/articleComments"
	"compose/comments/commentCommons"
	"compose/comments/createComment"
	"compose/comments/createReply"
	"compose/comments/deleteComment"
	"compose/comments/deleteReply"
	"compose/comments/replyThread"
	"compose/comments/updateComment"
	"compose/comments/updateReply"
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
	subRouter.HandleFunc("/replyThread", replyThread.Handler).Methods(http.MethodGet)

	subRouter.HandleFunc("/comment", createComment.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/updateComment", updateComment.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/deleteComment", deleteComment.Handler).Methods(http.MethodPost)

	subRouter.HandleFunc("/reply", createReply.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/updateReply", updateReply.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/deleteReply", deleteReply.Handler).Methods(http.MethodPost)
}
