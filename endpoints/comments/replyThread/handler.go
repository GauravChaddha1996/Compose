package replyThread

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
		Str(logger.FETCH, "Article reply thread").
		Str(logger.ARTICLE_ID, requestModel.ArticleId).
		Str(logger.PARENT_ID, requestModel.ParentId).
		Logger()
	subLogger := &subLoggerValue

	err = securityClearance(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response, err := getReplyThread(requestModel, subLogger)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	articleDao := daos.GetArticleDao()
	articleExists, err := articleDao.DoesArticleExist(model.ArticleId)
	if err != nil {
		return errors.New("Security problem. Can't confirm if article id exists")
	}
	if articleExists == false {
		return errors.New("No such article id exists")
	}
	return nil
}
