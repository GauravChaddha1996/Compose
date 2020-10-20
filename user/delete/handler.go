package delete

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

	err = securityClearance(requestModel)
	if commons.InError(err) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	err = deleteUser(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "User deleted successfully",
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	if model.commonModel.UserEmail != model.email {
		return errors.New("Email id doesn't match")
	}
	return nil
}
