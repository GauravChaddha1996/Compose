package user

import (
	"compose/commons"
	"encoding/json"
	"log"
	"net/http"
)

func SignupHandler(writer http.ResponseWriter, request *http.Request) {

	isSecure, message := IsSignupRequestSecure(request)
	if !isSecure {
		_writeFailedResponse(message, writer)
		return
	}

	requestModel := getUserSignupRequest(request)
	if requestModel == nil {
		_writeFailedResponse(SIGNUP_ERROR_PARSE_MESSAGE, writer)
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

	response := UserSignupResponseModel{
		Status:      commons.RESPONSE_STATUS_SUCCESS,
		AccessToken: accessToken,
	}
	jsonResponse, _ := json.Marshal(response)
	_, _ = writer.Write(jsonResponse)

}

func getUserSignupRequest(r *http.Request) *UserSignupRequestModel {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Print(err)
		return nil
	}
	//todo learn how to use enum in golang
	return &UserSignupRequestModel{
		Email: r.FormValue(SIGNUP_REQUEST_MODEL_EMAIL),
		Name:  r.FormValue(SIGNUP_REQUEST_MODEL_NAME),
	}
}

func _writeFailedResponse(message string, writer http.ResponseWriter) {
	// todo check how to do message ==nil then message default = "something went wrong"
	// todo better error reporting and handling
	failedResponse := UserSignupResponseModel{
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
