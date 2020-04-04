package login

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

	accessToken, err := login(requestModel)
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
