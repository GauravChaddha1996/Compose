package middlewares

import (
	"compose/commons"
	"log"
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

		println("")
		log.Println("----------------------RESPONSE RECORDED-------------------------")
		log.Println(r.Method + " " + r.URL.String())
		println("")
		log.Println(newWriter.Body.String())
		println("")

	})
}
