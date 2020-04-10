package article

import (
	"compose/article/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type ServiceContractImpl struct {
	articleDao  *daos.ArticleDao
	markdownDao *daos.ArticleMarkdownDao
}

func GetServiceContractImpl() ServiceContractImpl {
	return ServiceContractImpl{articleDao: daos.GetArticleDao(), markdownDao: daos.GetArticleMarkdownDao()}
}

func (impl ServiceContractImpl) DoesArticleExist(articleId string) (bool, error) {
	return impl.articleDao.DoesArticleExist(articleId)
}

func (impl ServiceContractImpl) GetArticleAuthorId(articleId string) *string {
	article, err := impl.articleDao.GetArticle(articleId)
	if commons.InError(err) {
		return nil
	}
	return &article.UserId
}

func (impl ServiceContractImpl) ChangeArticleLikeCount(articleId string, change bool) error {
	article, err := impl.articleDao.GetArticle(articleId)
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

	err = impl.articleDao.UpdateArticle(articleId, changeMap)
	if commons.InError(err) {
		return errors.New("Article like count can't be updated")
	}
	return nil
}

func (impl ServiceContractImpl) ChangeArticleCommentCount(articleId string, change bool, transaction *gorm.DB) error {
	articleDao := impl.getArticleDao(transaction)
	article, err := articleDao.GetArticle(articleId)
	if commons.InError(err) {
		return errors.New("Can't find any such article")
	}
	if change {
		article.CommentCount += 1
	} else {
		article.CommentCount -= 1
	}
	var changeMap = make(map[string]interface{})
	changeMap["comment_count"] = article.CommentCount

	err = articleDao.UpdateArticle(articleId, changeMap)
	if commons.InError(err) {
		return errors.New("Article comment count can't be updated")
	}
	return nil
}

func (impl ServiceContractImpl) GetAllArticlesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.Article, error) {
	return impl.articleDao.GetArticlesOfUser(userId, maxCreatedAtTime, limit)
}

func (impl ServiceContractImpl) GetAllArticles(articleIds []string) (*[]dbModels.Article, error) {
	return impl.articleDao.GetArticles(articleIds)
}

func (impl ServiceContractImpl) GetArticleMarkdown(markdownId string) (*dbModels.ArticleMarkdown, error) {
	return impl.markdownDao.GetArticleMarkdown(markdownId)
}

func (impl ServiceContractImpl) getArticleDao(transaction *gorm.DB) *daos.ArticleDao {
	if transaction != nil {
		return daos.GetArticleDaoDuringTransaction(transaction)
	} else {
		return impl.articleDao
	}
}
