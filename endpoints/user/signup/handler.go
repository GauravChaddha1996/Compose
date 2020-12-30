package signup

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
		Str(logger.ACTION, "User signup").
		Str(logger.EMAIL, requestModel.Email).
		Logger()
	subLogger := &subLoggerValue

	accessToken, err := signup(requestModel, subLogger)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLogger.Info().Msg("User signup done")

	response := ResponseModel{
		Status:      commons.NewResponseStatus().SUCCESS,
		AccessToken: accessToken,
	}

	commons.WriteSuccessResponse(response, writer)
}
