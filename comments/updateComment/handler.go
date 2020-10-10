package updateComment

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
	err = updateComment(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:    commons.NewResponseStatus().SUCCESS,
		CommentId: requestModel.CommentId,
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	comment, err := daos.GetCommentDao().FindComment(model.CommentId)
	if commons.InError(err) {
		return errors.New("Unable to find this comment")
	}
	if model.CommonModel.UserId != comment.UserId {
		return errors.New("Comment not posted by this user")
	}
	return nil
}
