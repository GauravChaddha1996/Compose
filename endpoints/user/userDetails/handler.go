package userDetails

import (
	"compose/commons"
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	user, err := getUserDetails(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

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
		// Only make editable if details requested of userId = user requesting the details
		Editable: user.UserId == commons.GetCommonRequestModel(request).UserId,
	}

	commons.WriteSuccessResponse(response, writer)
}
