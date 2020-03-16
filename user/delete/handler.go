package delete

import (
	"compose/commons"
	"compose/user/userCommons"
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
		writer.WriteHeader(403)
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
	var user userCommons.User
	db := userCommons.GetDB()
	emailQueryResult := db.Where("email = ?", model.email).Find(&user)
	if emailQueryResult.RecordNotFound() {
		return errors.New("User doesn't exist")
	}
	var accessTokenEntry userCommons.AccessToken
	accessTokenQuery := db.Where("user_id = ?", user.UserId).Find(&accessTokenEntry)
	if commons.InError(accessTokenQuery.Error) {
		return errors.New("Can't query access token")
	}
	if commons.GetCommonHeaders(request).AccessToken != accessTokenEntry.AccessToken {
		return errors.New("Access token doesn't match")
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
