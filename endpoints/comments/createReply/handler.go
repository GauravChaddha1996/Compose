package createReply

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
		Str(logger.ACTION, "Create reply").
		Str(logger.ARTICLE_ID, requestModel.ArticleId).
		Str(logger.PARENT_ID, requestModel.ParentId).
		Logger()
	subLogger := &subLoggerValue

	err = securityClearance(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response, err := createReply(requestModel, subLogger)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	articleDao := daos.GetArticleDao()
	articleExists, err := articleDao.DoesArticleExist(model.ArticleId)
	if err!=nil {
		return errors.New("Security problem. Can't confirm if article id exists")
	}
	if articleExists == false {
		return errors.New("No such article id exists")
	}
	commentDao := daos.GetCommentDao()
	replyDao := daos.GetReplyDao()

	commentExists := commentDao.DoesCommentExist(model.ParentId)
	replyParentExists := replyDao.DoesParentExist(model.ParentId)
	model.ParentIsComment = commentExists
	model.ParentIsReply = replyParentExists
	if commentExists == false && replyParentExists == false {
		// Neither the parent id is a top level comment or a reply
		return errors.New("No such parent id exists")
	}
	return nil
}
