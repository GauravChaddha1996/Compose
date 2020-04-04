package update

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	article, err := daos.GetArticleDao().GetArticle(requestModel.ArticleId)
	if commons.InError(err) {
		commons.WriteFailedResponse(errors.New("No such article Id exists"), writer)
		return
	}
	err = securityClearance(requestModel, article)
	if commons.InError(err) {
		writer.WriteHeader(http.StatusForbidden)
		commons.WriteFailedResponse(err, writer)
		return
	}

	err = updateArticle(requestModel, article)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:    commons.NewResponseStatus().SUCCESS,
		ArticleId: article.Id,
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel, article *articleCommons.Article) error {
	if model.CommonModel.UserId != article.UserId {
		return errors.New("Article not posted by this user")
	}
	return nil
}