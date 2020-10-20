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

type CommonRequestModel struct {
	AccessToken string
	UserId      string
	UserEmail   string
}

func GetCommonRequestModel(r *http.Request) *CommonRequestModel {
	return r.Context().Value(CommonModelKey).(*CommonRequestModel)
}

var SecurityMiddlewarePathConfigMap = make(map[string]*SecurityMiddlewarePathConfig)

type SecurityMiddlewarePathConfig struct {
	CheckAccessToken bool
	CheckUserId      bool
	CheckUserEmail   bool
}

func GetDefaultSecurityMiddlewarePathConfig() *SecurityMiddlewarePathConfig {
	return &SecurityMiddlewarePathConfig{
		CheckAccessToken: true,
		CheckUserId:      true,
		CheckUserEmail:   true,
	}
}

func AddSecurityMiddlewarePathConfig(path string, config *SecurityMiddlewarePathConfig) {
	SecurityMiddlewarePathConfigMap[path] = config
}
