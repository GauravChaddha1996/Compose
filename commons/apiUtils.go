package commons

import (
	"encoding/json"
	"net/http"
)

func WriteSuccessResponse(response interface{}, writer http.ResponseWriter) {
	jsonResponse, err := json.Marshal(response)
	PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	PanicIfError(err)
}

func WriteFailedResponse(err error, writer http.ResponseWriter) {
	failedResponse := genericErrorResponseModel{
		Status:  ResponseStatusWrapper{}.FAILED,
		Message: err.Error(),
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	PanicIfError(err)
	_, err = writer.Write(failedResponseJson)
	PanicIfError(err)
}

type genericErrorResponseModel struct {
	Status  ResponseStatus `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
}
