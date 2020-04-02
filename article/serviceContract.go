package article

import (
	"compose/article/daos"
	"compose/commons"
)

type ServiceContractImpl struct {
	dao *daos.ArticleDao
}

func GetServiceContractImpl() ServiceContractImpl {
	return ServiceContractImpl{dao: daos.GetArticleDao()}
}

func (impl ServiceContractImpl) GetArticleAuthorId(articleId string) *string {
	article, err := impl.dao.GetArticle(articleId)
	if commons.InError(err) {
		return nil
	}
	return &article.UserId
}
