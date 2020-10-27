package middlewares

import (
	"compose/commons"
	"compose/commons/logger"
	"net/http"
	"net/http/httptest"
)

func ResponseLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		newWriter := httptest.NewRecorder()
		next.ServeHTTP(newWriter, r)
		_, err := w.Write(newWriter.Body.Bytes())
		commons.PanicIfError(err)

		event := logger.RequestResponseLogger.Debug()
		event.Str("Endpoint", r.Method+" "+r.URL.Path)
		event.Str("Response", newWriter.Body.String())
		event.Msg("RESPONSE RECORDED")
		println()
	})
}
