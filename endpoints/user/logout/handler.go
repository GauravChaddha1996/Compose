package logout

import (
	"compose/commons"
	"compose/commons/logger"
	"compose/dataLayer/daos"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	commonModel := commons.GetCommonRequestModel(request)
	accessTokenDao := daos.GetAccessTokenDao()
	subLoggerValue := logger.Logger.With().
		Str(logger.ACTION, "Logout").
		Str(logger.USER_ID, commonModel.UserId).
		Logger()
	subLogger := &subLoggerValue

	err := accessTokenDao.DeleteAccessTokenEntry(commonModel.AccessToken)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(errors.New("Error deleting access token entry"), writer)
		return
	}
	subLogger.Info().Msg("Access token entry deleted")
	responseModel := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Succcessfully logged out",
	}
	commons.WriteSuccessResponse(responseModel, writer)
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
