package serviceContracts

import (
	"compose/dbModels"
	"github.com/jinzhu/gorm"
	"time"
)

type ArticleServiceContract interface {
	DoesArticleExist(articleId string) (bool, error)
	GetArticleAuthorId(articleId string) *string
	ChangeArticleLikeCount(articleId string, change bool, transaction *gorm.DB) error    // send change true to increase and false to decrease
	ChangeArticleCommentCount(articleId string, change bool, transaction *gorm.DB) error // send change true to increase and false to decrease
	GetAllArticlesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.Article, error)
	GetAllArticles(articleIds []string) (*[]dbModels.Article, error)
	GetArticleMarkdown(markdownId string) (*dbModels.ArticleMarkdown, error)
}
