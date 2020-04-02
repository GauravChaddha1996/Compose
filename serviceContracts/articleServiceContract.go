package serviceContracts

type ArticleServiceContract interface {
	DoesArticleExist(articleId string) bool
	GetArticleAuthorId(articleId string) *string
	ChangeArticleLikeCount(articleId string, change bool) error // send change true to increase and false to decrease
}
