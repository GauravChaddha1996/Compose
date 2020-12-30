package delete

import (
	"compose/commons"
	"compose/commons/logger"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
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
		Str(logger.ACTION, "Article delete").
		Str(logger.ARTICLE_ID, requestModel.ArticleId).
		Logger()
	subLogger := &subLoggerValue

	article, err := daos.GetArticleDao().GetArticle(requestModel.ArticleId)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(errors.New("No such article id exists"), writer)
		return
	}
	err = securityClearance(requestModel, article)
	if commons.InError2(err, subLogger) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	err = deleteArticle(article, subLogger)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Article deleted successfully",
	}

	commons.WriteSuccessResponse(response, writer)

}

func securityClearance(model *RequestModel, article *dbModels.Article) error {
	if model.CommonModel.UserId != article.UserId {
		return errors.New("Article not posted by this user")
	}
	return nil
}
