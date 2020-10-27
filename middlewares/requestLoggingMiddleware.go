package middlewares

import (
	"compose/commons/logger"
	"github.com/rs/zerolog"
	"net/http"
	"strings"
)

func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headerDict, requestGetParams, requestPostParams := GetDataFromRequest(r)

		event := logger.RequestResponseLogger.Debug()
		event.Str("Endpoint", r.Method+" "+r.URL.Path)
		if headerDict != nil {
			event.Dict("Headers", headerDict)
		}
		if requestGetParams != nil {
			event.Dict("GET params", requestGetParams)
		}
		if requestPostParams != nil {
			event.Dict("POST params", requestPostParams)
		}
		event.Msg("REQUEST RECEIVED")
		println()

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func GetDataFromRequest(r *http.Request) (*zerolog.Event, *zerolog.Event, *zerolog.Event) {
	var headerDict *zerolog.Event = nil
	var requestGetParams *zerolog.Event = nil
	var requestPostParams *zerolog.Event = nil

	for headerKey, headerValues := range r.Header {
		if headerDict == nil {
			headerDict = zerolog.Dict()
		}
		headerDict.Str(headerKey, strings.Join(headerValues, ","))
	}
	if r.Method == http.MethodGet {
		queryMap := r.URL.Query()
		if len(queryMap) > 0 {
			requestGetParams = zerolog.Dict()
			for key, value := range queryMap {
				requestGetParams.Str(key, strings.Join(value, ","))
			}
		}
	}
	if r.Method == http.MethodPost {
		_ = r.ParseMultipartForm(1024)
		form := r.Form
		if len(form) > 0 {
			requestPostParams = zerolog.Dict()
			for bodyKey, bodyValues := range r.Form {
				requestPostParams.Str(bodyKey, strings.Join(bodyValues, ","))
			}
		}
	}
	return headerDict, requestGetParams, requestPostParams
}
