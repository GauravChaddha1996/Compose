package likeArticle

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	articleDao := daos.GetArticleDao()
	article, err := articleDao.GetArticle(requestModel.ArticleId)
	if commons.InError(err) {
		commons.WriteFailedResponse(errors.New("No such article id exists"), writer)
		return
	}
	err = securityClearance(requestModel, &article.UserId)
	if commons.InError(err) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	err = likeArticle(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Article liked successfully",
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel, articleUserId *string) error {
	if model.CommonModel.UserId == *articleUserId {
		return errors.New("You cannot like your own article")
	}
	return nil
}
