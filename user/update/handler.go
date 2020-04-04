package update

import (
	"compose/commons"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	err = securityClearance(requestModel, request)
	if commons.InError(err) {
		writer.WriteHeader(http.StatusForbidden)
		commons.WriteFailedResponse(err, writer)
		return
	}

	err = update(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "User updated successfully",
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel, request *http.Request) error {
	commonsModel := commons.GetCommonModel(request)
	if commonsModel.UserId != model.UserId {
		return errors.New("User id doesn't match")
	}
	return nil
}
