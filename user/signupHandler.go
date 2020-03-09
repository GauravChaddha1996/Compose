package user

import (
	"compose/commons"
	"encoding/json"
	"log"
	"net/http"
)

func SignupHandler(writer http.ResponseWriter, request *http.Request) {
	requestModel := getUserSignupRequest(request)
	if requestModel == nil {
		_writeFailedResponse(ERROR_PARSE_MESSAGE, writer)
		return
	}

	isValid, message := IsUserSignupRequestValid(requestModel)
	if !isValid {
		_writeFailedResponse(message, writer)
		return
	}

	// todo check how to to send back null access token
	accessToken, message := signup(requestModel)
	if len(accessToken) == 0 {
		_writeFailedResponse(message, writer)
		return
	}

	response := SignupResponseModel{
		Status:      commons.RESPONSE_STATUS_SUCCESS,
		AccessToken: accessToken,
	}
	// todo find better way for marshalling
	jsonResponse, _ := json.Marshal(response)
	_, _ = writer.Write(jsonResponse)

}

func getUserSignupRequest(r *http.Request) *SignupRequestModel {
	err := r.ParseMultipartForm(1024)
	if err != nil {
		log.Print(err)
		return nil
	}
	//todo learn how to use enum in golang
	return &SignupRequestModel{
		Email:    r.FormValue(REQUEST_MODEL_EMAIL),
		Name:     r.FormValue(REQUEST_MODEL_NAME),
		Password: r.FormValue(REQUEST_MODEL_PASSWORD),
	}
}

func _writeFailedResponse(message string, writer http.ResponseWriter) {
	// todo check how to do message ==nil then message default = "something went wrong"
	// todo better error reporting and handling
	failedResponse := SignupResponseModel{
		Status:  commons.RESPONSE_STATUS_FAILED,
		Message: message,
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(failedResponseJson)
	if err != nil {
		panic(err)
	}
}
