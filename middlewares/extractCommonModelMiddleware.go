package middlewares

import (
	"compose/commons"
	"compose/dataLayer/models"
	"context"
	"net/http"
)

func ExtractCommonModelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commonHeaders := MakeCommonRequestModel(r)
		parentContext := r.Context()
		newContext := context.WithValue(parentContext, commons.CommonModelKey, commonHeaders)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}

func MakeCommonRequestModel(r *http.Request) *commons.CommonRequestModel {
	headers := r.Header
	accessToken := headers.Get("access_token")
	userId := getUserId(accessToken)
	userEmail := getUserEmail(userId)
	return &commons.CommonRequestModel{
		AccessToken: accessToken,
		UserId:      userId,
		UserEmail:   userEmail,
	}
}

func getUserId(accessToken string) string {
	var accessTokenEntry models.AccessToken
	accessTokenQuery := commons.GetDB().Where("access_token = ?", accessToken).Find(&accessTokenEntry)
	if commons.InError(accessTokenQuery.Error) {
		return ""
	}
	return accessTokenEntry.UserId
}

func getUserEmail(userId string) string {
	var user models.User
	userQuery := commons.GetDB().Where("user_id = ?", userId).Find(&user)
	if commons.InError(userQuery.Error) {
		return ""
	}
	return user.Email
}
