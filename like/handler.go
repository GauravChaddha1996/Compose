package like

import (
	"compose/like/articleLikes"
	"compose/like/likeArticle"
	"compose/like/unlikeArticle"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/like", likeArticle.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/unlike", unlikeArticle.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/get_likes", articleLikes.Handler).Methods(http.MethodGet)
}
