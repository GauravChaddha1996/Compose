package article

import (
	"compose/article/daos"
	"compose/commons"
	"errors"
)

type ServiceContractImpl struct {
	dao *daos.ArticleDao
}

func GetServiceContractImpl() ServiceContractImpl {
	return ServiceContractImpl{dao: daos.GetArticleDao()}
}

func (impl ServiceContractImpl) DoesArticleExist(articleId string) bool {
	return impl.dao.DoesArticleExist(articleId)
}

func (impl ServiceContractImpl) GetArticleAuthorId(articleId string) *string {
	article, err := impl.dao.GetArticle(articleId)
	if commons.InError(err) {
		return nil
	}
	return &article.UserId
}

func (impl ServiceContractImpl) ChangeArticleLikeCount(articleId string, change bool) error {
	article, err := impl.dao.GetArticle(articleId)
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

	err = impl.dao.UpdateArticle(articleId, changeMap)
	if commons.InError(err) {
		return errors.New("Article like count can't be updated")
	}
	return nil
}
