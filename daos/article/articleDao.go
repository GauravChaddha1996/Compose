package daos

import (
	"compose/article/articleCommons"
	"compose/commons"
	"compose/dbModels"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
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

func (dao ArticleDao) CreateArticle(article dbModels.Article) error {
	return dao.db.Create(article).Error
}

func (dao ArticleDao) DoesArticleExist(articleId string) (bool, error) {
	var article dbModels.Article
	queryResult := dao.db.
		Select("id").
		Where("id = ?", articleId).
		Limit(1).
		Find(&article)
	if commons.InError(queryResult.Error) {
		return false, queryResult.Error
	} else {
		return true, nil
	}
}

func (dao ArticleDao) GetArticle(articleId string) (*dbModels.Article, error) {
	var article dbModels.Article
	articleQuery := dao.db.Where("id = ?", articleId).Find(&article)
	if commons.InError(articleQuery.Error) {
		return nil, articleQuery.Error
	}
	return &article, nil
}

func (dao ArticleDao) GetArticles(articleIds []string) (*[]dbModels.Article, error) {
	var articles []dbModels.Article
	orderByExpr := fmt.Sprintf("FIELD(id, '%s') asc", strings.Join(articleIds, "','"))
	articleQuery := dao.db.
		Where("id IN (?)", articleIds).
		Order(gorm.Expr(orderByExpr)).
		Find(&articles)
	if commons.InError(articleQuery.Error) {
		return nil, articleQuery.Error
	}
	return &articles, nil
}

func (dao ArticleDao) GetArticlesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.Article, error) {
	var articles []dbModels.Article
	articleQuery := dao.db.
		Where("user_id = ? && created_at < ?", userId, maxCreatedAtTime).
		Order("created_at desc").
		Limit(limit).
		Find(&articles)
	if commons.InError(articleQuery.Error) {
		return nil, articleQuery.Error
	}
	return &articles, nil
}

func (dao ArticleDao) UpdateArticle(articleId string, changeMap map[string]interface{}) error {
	var article dbModels.Article
	return dao.db.Model(article).Where("id = ?", articleId).UpdateColumns(changeMap).Error
}

func (dao ArticleDao) DeleteArticle(article *dbModels.Article) error {
	return dao.db.Unscoped().Delete(&article).Error
}
