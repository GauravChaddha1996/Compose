package daos

import (
	"compose/article/articleCommons"
	"compose/commons"
	"github.com/jinzhu/gorm"
)

type ArticleDao struct {
	db *gorm.DB
}

func GetArticleDaoDuringTransaction(db *gorm.DB) *ArticleDao {
	return &ArticleDao{db: db}
}

func GetArticleDao() *ArticleDao {
	return &ArticleDao{db: articleCommons.Database}
}

func (dao ArticleDao) CreateArticle(article articleCommons.Article) error {
	return dao.db.Create(article).Error
}

func (dao ArticleDao) DoesArticleExist(articleId string) bool {
	var article articleCommons.Article
	queryResult := dao.db.
		Select("id").
		Where("id = ?", articleId).
		Limit(1).
		Find(&article)
	return queryResult.Error == nil && article.Id == articleId
}

func (dao ArticleDao) GetArticle(articleId string) (*articleCommons.Article, error) {
	var article articleCommons.Article
	articleQuery := dao.db.Where("id = ?", articleId).Find(&article)
	if commons.InError(articleQuery.Error) {
		return nil, articleQuery.Error
	}
	return &article, nil
}

func (dao ArticleDao) UpdateArticle(articleId string, changeMap map[string]interface{}) error {
	var article articleCommons.Article
	return dao.db.Model(article).Where("id = ?", articleId).UpdateColumns(changeMap).Error
}
