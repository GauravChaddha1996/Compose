package articleComments

import (
	"compose/commons"
	"compose/daos"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	err = securityClearance(requestModel)
	if commons.InError(err) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	response, err := getArticleComments(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	articleDao := daos.GetArticleDao()
	articleExists, err := articleDao.DoesArticleExist(model.ArticleId)
	if commons.InError(err) {
		return errors.New("Security problem. Can't confirm if article id exists")
	}
	if articleExists == false {
		return errors.New("No such article id exists")
	}
	return nil
}
