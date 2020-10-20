package commons

import (
	"encoding/json"
	"net/http"
)

func NewResponseStatus() ResponseStatusWrapper {
	return ResponseStatusWrapper{
		SUCCESS: "success",
		FAILED:  "failed",
	}
}

func GetCommonRequestModel(r *http.Request) *CommonRequestModel {
	return r.Context().Value(CommonModelKey).(*CommonRequestModel)
}

func WriteSuccessResponse(response interface{}, writer http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	PanicIfError(err)
}

func WriteFailedResponse(err error, writer http.ResponseWriter) {
	failedResponse := GenericErrorResponseModel{
		Status:  NewResponseStatus().FAILED,
		Message: err.Error(),
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	PanicIfError(err)
	_, err = writer.Write(failedResponseJson)
	PanicIfError(err)
}

func WriteForbiddenResponse(err error, writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusForbidden)
	WriteFailedResponse(err, writer)
}
