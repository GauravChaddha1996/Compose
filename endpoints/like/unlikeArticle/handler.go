package unlikeArticle

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
		Str(logger.ACTION, "Unlike article").
		Str(logger.USER_ID, requestModel.CommonModel.UserId).
		Str(logger.ARTICLE_ID, requestModel.ArticleId).
		Logger()
	subLogger := &subLoggerValue

	err = securityClearance(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	err = unlikeArticle(requestModel, subLogger)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Article unliked successfully",
	}
	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel) error {
	articleDao := daos.GetArticleDao()
	article, err := articleDao.GetArticle(model.ArticleId)
	if commons.InError(err) {
		return errors.New("No such article id exists")
	}
	if model.CommonModel.UserId == article.UserId {
		return errors.New("You cannot unlike your own article")
	}
	return nil
}
