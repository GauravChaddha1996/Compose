package likedArticles

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

	response, err := getLikedArticles(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	commons.WriteSuccessResponse(response, writer)
}
