package deleteComment

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
		Str(logger.ACTION, "Delete comment").
		Str(logger.COMMENT_ID, requestModel.CommentId).
		Logger()
	subLogger := &subLoggerValue

	err = securityClearance(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	err = deleteComment(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Comment deleted successfully",
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	comment, err := daos.GetCommentDao().FindComment(model.CommentId)
	if err != nil {
		return errors.New("Error finding the comment")
	}
	if model.CommonModel.UserId != comment.UserId {
		return errors.New("Comment not posted by this user")
	}
	return nil
}
