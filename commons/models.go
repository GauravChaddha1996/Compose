package commons

import (
	"compose/dbModels"
	"net/http"
)

const CommonModelKey = "common_model"

type ResponseStatus string

type ResponseStatusWrapper struct {
	SUCCESS ResponseStatus
	FAILED  ResponseStatus
}

func NewResponseStatus() ResponseStatusWrapper {
	return ResponseStatusWrapper{
		SUCCESS: "success",
		FAILED:  "failed",
	}
}

type CommonModel struct {
	AccessToken string
	UserId      string
	UserEmail   string
}

func GetCommonModel(r *http.Request) *CommonModel {
	return r.Context().Value(CommonModelKey).(*CommonModel)
}

func makeCommonModel(r *http.Request) *CommonModel {
	headers := r.Header
	accessToken := headers.Get("access_token")
	userId := getUserId(accessToken)
	userEmail := getUserEmail(userId)
	return &CommonModel{
		AccessToken: accessToken,
		UserId:      userId,
		UserEmail:   userEmail,
	}
}

func getUserId(accessToken string) string {
	var accessTokenEntry dbModels.AccessToken
	accessTokenQuery := GetDB().Where("access_token = ?", accessToken).Find(&accessTokenEntry)
	if InError(accessTokenQuery.Error) {
		return ""
	}
	return accessTokenEntry.UserId
}

func getUserEmail(userId string) string {
	var user dbModels.User
	userQuery := GetDB().Where("user_id = ?", userId).Find(&user)
	if InError(userQuery.Error) {
		return ""
	}
	return user.Email
}
