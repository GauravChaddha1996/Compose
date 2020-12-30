package deleteReply

import (
	"compose/commons"
	"compose/commons/logger"
	"compose/dataLayer/daos"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError2(err, nil) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLoggerValue := logger.Logger.With().
		Str(logger.ACTION, "Delete reply").
		Str(logger.REPLY_ID, requestModel.ReplyId).
		Logger()
	subLogger := &subLoggerValue

	err = securityClearance(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	err = deleteReply(requestModel)
	if commons.InError2(err, subLogger) {
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
	if err != nil {
		return errors.New("Error finding the reply")
	}
	if model.CommonModel.UserId != reply.UserId {
		return errors.New("Reply not posted by this user")
	}
	return nil
}
