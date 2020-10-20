package like

import (
	articleLikes2 "compose/endpoints/like/articleLikes"
	likeArticle2 "compose/endpoints/like/likeArticle"
	unlikeArticle2 "compose/endpoints/like/unlikeArticle"
	"github.com/gorilla/mux"
	"net/http"
)

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/like", likeArticle2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/unlike", unlikeArticle2.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/get_likes", articleLikes2.Handler).Methods(http.MethodGet)
}
