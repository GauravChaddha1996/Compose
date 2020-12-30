package create

import (
	"compose/commons"
	"compose/commons/logger"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLoggerValue := logger.Logger.With().
		Str(logger.ACTION, "Article create").
		Logger()
	subLogger := &subLoggerValue

	articleId, err := createArticle(requestModel, subLogger)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:    commons.NewResponseStatus().SUCCESS,
		ArticleId: *articleId,
	}

	commons.WriteSuccessResponse(response, writer)
}
