package daos

import (
	"compose/article/articleCommons"
	"github.com/jinzhu/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func GetArticleDaoDuringTransaction(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db: db}
}

func (dao ArticleDao) CreateArticle(article articleCommons.Article) error {
	return dao.db.Create(article).Error
}
