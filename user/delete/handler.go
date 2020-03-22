package delete

import (
	"compose/commons"
	"encoding/json"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	err = securityClearance(requestModel, request)
	if commons.InError(err) {
		writer.WriteHeader(http.StatusForbidden)
		writeFailedResponse(err, writer)
		return
	}

	err = delete(requestModel)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "User deleted successfully",
	}

	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)

}

func securityClearance(model *RequestModel, request *http.Request) error {
	commonsModel := request.Context().Value(commons.CommonModelKey).(commons.CommonModel)
	if commonsModel.UserEmail != model.email {
		return errors.New("Email id doesn't match")
	}
	return nil
}

func writeFailedResponse(err error, writer http.ResponseWriter) {
	failedResponse := ResponseModel{
		Status:  commons.NewResponseStatus().FAILED,
		Message: err.Error(),
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	commons.PanicIfError(err)
	_, err = writer.Write(failedResponseJson)
	commons.PanicIfError(err)
}
