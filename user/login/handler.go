package login

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

	accessToken, err := login(requestModel)
	if commons.InError(err) {
		_writeFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:      commons.NewResponseStatus().SUCCESS,
		AccessToken: accessToken,
	}

	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)

}

func _writeFailedResponse(err error, writer http.ResponseWriter) {
	failedResponse := ResponseModel{
		Status:  commons.NewResponseStatus().FAILED,
		Message: err.Error(),
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	commons.PanicIfError(err)
	_, err = writer.Write(failedResponseJson)
	commons.PanicIfError(err)
}
