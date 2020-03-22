package update

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

	err = update(requestModel)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "User updated successfully",
	}

	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)

}

func securityClearance(model *RequestModel, request *http.Request) error {
	commonsModel := commons.GetCommonModel(request)
	if commonsModel.UserId != model.UserId {
		return errors.New("User id doesn't match")
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
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(failedResponseJson)
	commons.PanicIfError(err)
}
