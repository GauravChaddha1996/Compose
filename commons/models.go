package commons

import (
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
	headers := r.Header
	accessToken := headers.Get("access_token")
	return &CommonModel{
		AccessToken: accessToken,
	}
}
