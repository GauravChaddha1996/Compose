package signup

import (
	"compose/commons"
	"encoding/json"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		_writeFailedResponse(err, writer)
		return
	}

	// todo check how to to send back null access token
	accessToken, err := signup(requestModel)
	if commons.InError(err) {
		_writeFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:      commons.RESPONSE_STATUS_SUCCESS,
		AccessToken: accessToken,
	}

	// todo find better way for marshalling
	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)
}

func _writeFailedResponse(err error, writer http.ResponseWriter) {
	// todo better error reporting and handling
	failedResponse := ResponseModel{
		Status:  commons.RESPONSE_STATUS_FAILED,
		Message: err.Error(),
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	commons.PanicIfError(err)
	_, err = writer.Write(failedResponseJson)
	commons.PanicIfError(err)
}
