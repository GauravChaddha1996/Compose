package deleteReply

import (
	"compose/comments/daos"
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

	err = deleteReply(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Reply deleted successfully",
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	reply, err := daos.GetReplyDao().FindReply(model.ReplyId)
	if commons.InError(err) {
		return errors.New("Error finding the reply")
	}
	if model.CommonModel.UserId != reply.UserId {
		return errors.New("Reply not posted by this user")
	}
	return nil
}
