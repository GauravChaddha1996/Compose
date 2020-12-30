package articleDetails

import (
	"compose/commons"
	"compose/commons/logger"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError2(err, nil) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLoggerValue := logger.Logger.With().
		Str(logger.FETCH, "Article details").
		Str(logger.ARTICLE_ID, requestModel.Id).
		Logger()
	subLogger := &subLoggerValue

	response, err := getArticleDetailsResponse(requestModel, subLogger)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	commons.WriteSuccessResponse(response, writer)
}
