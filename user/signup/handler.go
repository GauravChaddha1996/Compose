package signup

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

	accessToken, err := signup(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:      commons.NewResponseStatus().SUCCESS,
		AccessToken: accessToken,
	}

	commons.WriteSuccessResponse(response, writer)
}