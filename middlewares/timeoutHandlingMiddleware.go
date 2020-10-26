package middlewares

import (
	"compose/commons/globalConfigHolders"
	"net/http"
)

func TimeoutHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeoutConfig := globalConfigHolders.GetTimeoutConfig(r.URL.Path)
		timeoutNextHandler := http.TimeoutHandler(next, timeoutConfig.TimeoutInSeconds, "Request timeout")
		timeoutNextHandler.ServeHTTP(w, r)
	})
}
