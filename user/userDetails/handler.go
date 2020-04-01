package userDetails

import (
	"compose/commons"
	"encoding/json"
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		_writeFailedResponse(err, writer)
		return
	}

	user, err := getUserDetails(requestModel)
	if commons.InError(err) {
		_writeFailedResponse(err, writer)
		return
	}

	createdAtTime := user.CreatedAt
	response := ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.Description,
		PhotoUrl:     user.PhotoUrl,
		ArticleCount: user.ArticleCount,
		MemberSince:  fmt.Sprint("Member since: ", createdAtTime.Day(), createdAtTime.Month(), createdAtTime.Year()),
		// Only make editable if details requested of userId = user requesting the details
		Editable: user.UserId == commons.GetCommonModel(request).UserId,
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
