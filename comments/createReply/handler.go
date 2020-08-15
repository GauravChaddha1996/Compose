package createReply

import (
	"compose/comments/commentCommons"
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

	response, err := createReply(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	articleExists, err := commentCommons.ArticleServiceContract.DoesArticleExist(model.ArticleId)
	if commons.InError(err) {
		return errors.New("Security problem. Can't confirm if article id exists")
	}
	if articleExists == false {
		return errors.New("No such article id exists")
	}
	tx := commentCommons.Database.Begin()
	commentDao := daos.GetCommentDaoDuringTransaction(tx)
	replyDao := daos.GetReplyDaoDuringTransaction(tx)

	commentExists := commentDao.DoesCommentExist(model.ParentId)
	replyParentExists := replyDao.DoesParentExist(model.ParentId)
	if commentExists == false && replyParentExists == false {
		// Neither the parent id is a top level comment or a reply
		return errors.New("No such parent id exists")
	}
	return nil
}