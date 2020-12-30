package articleDetails

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/rs/zerolog"
)

func getArticleDetailsResponse(model *RequestModel, subLogger *zerolog.Logger) (*ResponseModel, error) {
	article, err := getArticleDetails(model)
	if commons.InError(err) {
		return nil, errors.New("Can't fetch article details'")
	}
	subLogger.Info().Msg("Article details fetched")
	articleMarkdown, err := getArticleMarkdown(article.MarkdownId)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch articleMarkdown details")
	}
	subLogger.Info().Msg("Article markdown fetched")
	postedByUser, err := getPostedByUser(article)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch user who posted this article")
	}
	subLogger.Info().Msg("Article user fetched")

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
	article, err := articleDao.GetArticle(model.Id)
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

func getPostedByUser(article *dbModels.Article) (*apiEntity.SmallUserEntity, error) {
	userDao := daos.GetUserDao()
	user, err := userDao.FindUserViaId(article.UserId)
	if commons.InError(err) {
		return nil, err
	}
	return apiEntity.GetSmallUserEntity(user), nil
}
