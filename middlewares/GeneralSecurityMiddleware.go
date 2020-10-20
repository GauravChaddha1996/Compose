package middlewares

import (
	"compose/commons"
	"errors"
	"net/http"
)

func GeneralSecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commonModel := r.Context().Value(commons.CommonModelKey).(*commons.CommonRequestModel)
		securityError := ensureSecurity(commonModel, commons.SecurityMiddlewarePathConfigMap[r.URL.Path])
		if securityError != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte("{\"message\":\"User id associated isn't known\"}"))
			commons.PanicIfError(err)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func ensureSecurity(commonHeaders *commons.CommonRequestModel, config *commons.SecurityMiddlewarePathConfig) error {
	if config == nil {
		config = commons.GetDefaultSecurityMiddlewarePathConfig()
	}
	if config.CheckAccessToken && len(commonHeaders.AccessToken) == 0 {
		return errors.New("Access token isn't present")
	}
	if config.CheckUserId && len(commonHeaders.UserId) == 0 {
		return errors.New("User id associated isn't known")
	}
	if config.CheckUserEmail && len(commonHeaders.UserEmail) == 0 {
		return errors.New("User email associated isn't known")
	}
	return nil
}
