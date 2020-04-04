package likedArticles

import (
	"compose/commons"
	"compose/user/daos"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	err = securityClearance(requestModel)
	if commons.InError(err) {
		writer.WriteHeader(http.StatusForbidden)
		commons.WriteFailedResponse(err, writer)
		return
	}

	response, err := getLikedArticles(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	exist, err := daos.GetUserDao().DoesUserIdExist(model.UserId)
	if commons.InError(err) {
		return errors.New("Security validation of user id fails")
	}
	if exist == false {
		return errors.New("User doesn't exist")
	}
	return nil
}
