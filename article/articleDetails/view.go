package articleDetails

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
)

func getArticleDetailsResponse(model *RequestModel) (*ResponseModel, error) {
	article, err := getArticleDetails(model)
	if commons.InError(err) {
		return nil, errors.New("Can't fetch article details'")
	}
	articleMarkdown, err := getArticleMarkdown(article.MarkdownId)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch articleMarkdown details")
	}
	postedByUser, err := getPostedByUser(article)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch user who posted this article")
	}
	return &ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		Message:      "",
		Title:        article.Title,
		Description:  article.Description,
		Markdown:     articleMarkdown.Markdown,
		LikeCount:    article.LikeCount,
		CommentCount: article.TotalCommentCount,
		CreatedAt:    article.CreatedAt.Format("Posted on ", ),
		PostedBy:     *postedByUser,
		Editable:     model.commonModel.UserId == article.UserId,
	}, nil
}

func getArticleDetails(model *RequestModel) (*dbModels.Article, error) {
	articleDao := daos.GetArticleDao()
	article, err := articleDao.GetArticle(model.id)
	if commons.InError(err) {
		return nil, err
	}
	return article, nil
}

func getArticleMarkdown(markdownId string) (*dbModels.ArticleMarkdown, error) {
	dao := daos.GetArticleMarkdownDao()
	markdown, err := dao.GetArticleMarkdown(markdownId)
	if commons.InError(err) {
		return nil, err
	}
	return markdown, nil
}

func getPostedByUser(article *dbModels.Article) (*PostedByUser, error) {
	user, err := articleCommons.UserServiceContract.GetUser(article.UserId)
	if commons.InError(err) {
		return nil, err
	}
	return &PostedByUser{
		UserId:   article.UserId,
		Name:     user.Name,
		PhotoUrl: user.PhotoUrl,
	}, nil
}
