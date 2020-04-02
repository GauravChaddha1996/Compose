package serviceContracts

type ArticleServiceContract interface {
	GetArticleAuthorId(articleId string) *string
}
