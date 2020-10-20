package create

import (
	"compose/commons"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	articleId, err := createArticle(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:    commons.NewResponseStatus().SUCCESS,
		ArticleId: *articleId,
	}

	commons.WriteSuccessResponse(response, writer)
}
