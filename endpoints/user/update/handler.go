package update

import (
	"compose/commons"
	"compose/commons/logger"
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
		Str(logger.ACTION, "User update").
		Str(logger.USER_ID, requestModel.UserId).
		Logger()
	subLogger := &subLoggerValue

	err = securityClearance(requestModel, request)
	if commons.InError2(err, subLogger) {
		commons.WriteForbiddenResponse(err, writer)
		return
	}

	err = update(requestModel, subLogger)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLogger.Info().Msg("User is updated")

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "User updated successfully",
	}

	commons.WriteSuccessResponse(response, writer)
}

func securityClearance(model *RequestModel, request *http.Request) error {
	commonsModel := commons.GetCommonRequestModel(request)
	if commonsModel.UserId != model.UserId {
		return errors.New("User id doesn't match")
	}
	return nil
}
