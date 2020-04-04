package update

import (
	"compose/article/daos"
	"compose/commons"
	"compose/dbModels"
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
		commons.WriteForbiddenResponse(err, writer)
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

func securityClearance(model *RequestModel, article *dbModels.Article) error {
	if model.CommonModel.UserId != article.UserId {
		return errors.New("Article not posted by this user")
	}
	return nil
}