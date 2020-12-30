package userDetails

import (
	"compose/commons"
	"compose/commons/logger"
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError2(err, nil) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLoggerValue := logger.Logger.With().
		Str(logger.FETCH, "User details").
		Str(logger.USER_ID, requestModel.UserId).
		Logger()
	subLogger := &subLoggerValue

	user, err := getUserDetails(requestModel)
	if commons.InError2(err, subLogger) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	subLogger.Info().Msg("User details is fetched")

	createdAtTime := user.CreatedAt
	response := ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		Email:        user.Email,
		Name:         user.Name,
		Description:  user.Description,
		PhotoUrl:     user.PhotoUrl,
		ArticleCount: user.ArticleCount,
		LikeCount:    user.LikeCount,
		MemberSince:  fmt.Sprint("Member since: ", createdAtTime.Day(), createdAtTime.Month(), createdAtTime.Year()),
		// Only make editable if details requested of UserId = user requesting the details
		Editable: user.UserId == commons.GetCommonRequestModel(request).UserId,
	}

	commons.WriteSuccessResponse(response, writer)
}
