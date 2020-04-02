package like

import (
	"compose/like/likeArticle"
	"compose/like/likeCommons"
	"compose/like/unlikeArticle"
	"compose/serviceContracts"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Init(db *gorm.DB) {
	likeCommons.Database = db
}

func SetServiceContractImpl(articleContract serviceContracts.ArticleServiceContract) {
	likeCommons.ArticleServiceContract = articleContract
}

func AddSubRoutes(subRouter *mux.Router) {
	subRouter.HandleFunc("/like", likeArticle.Handler).Methods(http.MethodPost)
	subRouter.HandleFunc("/unlike", unlikeArticle.Handler).Methods(http.MethodPost)
}
