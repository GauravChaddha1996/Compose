package article

import (
	"compose/commons"
	"compose/dbModels"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ArticleDao struct {
	DB *gorm.DB
}

func (dao ArticleDao) CreateArticle(article dbModels.Article) error {
	return dao.DB.Create(article).Error
}

func (dao ArticleDao) DoesArticleExist(articleId string) (bool, error) {
	var article dbModels.Article
	queryResult := dao.DB.
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
	articleQuery := dao.DB.Where("id = ?", articleId).Find(&article)
	if commons.InError(articleQuery.Error) {
		return nil, articleQuery.Error
	}
	return &article, nil
}

func (dao ArticleDao) GetArticles(articleIds []string) (*[]dbModels.Article, error) {
	var articles []dbModels.Article
	orderByExpr := fmt.Sprintf("FIELD(id, '%s') asc", strings.Join(articleIds, "','"))
	articleQuery := dao.DB.
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
	articleQuery := dao.DB.
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
	return dao.DB.Model(article).Where("id = ?", articleId).UpdateColumns(changeMap).Error
}

func (dao ArticleDao) DeleteArticle(article *dbModels.Article) error {
	return dao.DB.Unscoped().Delete(&article).Error
}

/* Helper f()s */

func (dao ArticleDao) ChangeArticleLikeCount(articleId string, change bool) error {
	article, err := dao.GetArticle(articleId)
	if commons.InError(err) {
		return errors.New("Can't find any such article")
	}
	if change {
		article.LikeCount += 1
	} else {
		article.LikeCount -= 1
	}
	var changeMap = make(map[string]interface{})
	changeMap["like_count"] = article.LikeCount

	err = dao.UpdateArticle(articleId, changeMap)
	if commons.InError(err) {
		return errors.New("Article like count can't be updated")
	}
	return nil
}

func (dao ArticleDao) ChangeArticleReplyCommentCount(articleId string, change bool, transaction *gorm.DB) error {
	article, err := dao.GetArticle(articleId)
	if commons.InError(err) {
		return errors.New("Can't find any such article")
	}
	if change {
		article.TotalCommentCount += 1
	} else {
		article.TotalCommentCount -= 1
	}
	var changeMap = make(map[string]interface{})
	changeMap["total_comment_count"] = article.TotalCommentCount

	err = dao.UpdateArticle(articleId, changeMap)
	if commons.InError(err) {
		return errors.New("Article reply count can't be updated")
	}
	return nil
}

func (dao ArticleDao) ChangeArticleTopCommentCount(articleId string, change bool, transaction *gorm.DB) error {
	article, err := dao.GetArticle(articleId)
	if commons.InError(err) {
		return errors.New("Can't find any such article")
	}
	if change {
		article.TopCommentCount += 1
		article.TotalCommentCount += 1
	} else {
		article.TopCommentCount -= 1
		article.TotalCommentCount -= 1
	}
	var changeMap = make(map[string]interface{})
	changeMap["top_comment_count"] = article.TopCommentCount
	changeMap["total_comment_count"] = article.TotalCommentCount

	err = dao.UpdateArticle(articleId, changeMap)
	if commons.InError(err) {
		return errors.New("Article top comment count can't be updated")
	}
	return nil
}

func (dao ArticleDao) GetArticleTopCommentCount(articleId string) uint64 {
	article, err := dao.GetArticle(articleId)
	if commons.InError(err) {
		return 0
	}
	return article.TopCommentCount
}
