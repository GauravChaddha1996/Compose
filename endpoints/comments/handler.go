package comments

import (
	articleComments2 "compose/endpoints/comments/articleComments"
	createComment2 "compose/endpoints/comments/createComment"
	createReply2 "compose/endpoints/comments/createReply"
	deleteComment2 "compose/endpoints/comments/deleteComment"
	deleteReply2 "compose/endpoints/comments/deleteReply"
	replyThread2 "compose/endpoints/comments/replyThread"
	updateComment2 "compose/endpoints/comments/updateComment"
	updateReply2 "compose/endpoints/comments/updateReply"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/comments", articleComments2.Handler).Methods(http.MethodGet)
	subRouter.HandleFunc("/replyThread", replyThread2.Handler).Methods(http.MethodGet)

	subRouter.HandleFunc("/comment", createComment2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/updateComment", updateComment2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/deleteComment", deleteComment2.Handler).Methods(http.MethodPost)

	subRouter.HandleFunc("/reply", createReply2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/updateReply", updateReply2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/deleteReply", deleteReply2.Handler).Methods(http.MethodPost)
}
