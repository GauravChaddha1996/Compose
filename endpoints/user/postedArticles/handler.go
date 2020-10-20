package postedArticles

import (
	"compose/commons"
	"compose/daos"
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
		commons.WriteFailedResponse(err, writer)
		return
	}

	response, err := getPostedArticles(requestModel)
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
