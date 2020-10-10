package deleteComment

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

	err = deleteComment(requestModel)
	if commons.InError(err) {
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
	if commons.InError(err) {
		return errors.New("Error finding the comment")
	}
	if model.CommonModel.UserId != comment.UserId {
		return errors.New("Comment not posted by this user")
	}
	return nil
}
