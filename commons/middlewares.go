package commons

import (
	"log"
	"net/http"
	"strings"
)

func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("")
		log.Println("-----------------------------------------------------------------------")
		log.Println("REQUEST RECIEVED")
		log.Println("-----------------------------------------------------------------------")
		log.Println(r.Method + " " + r.URL.String())
		log.Println("")
		log.Println("HEADERS")
		log.Println(".................................")
		for headerKey, headerValues := range r.Header {
			log.Println(headerKey + " : " + strings.Join(headerValues, ","))
		}

		if r.Method == http.MethodGet {
			log.Println("")
			log.Println("QUERY PARAMS")
			log.Println(".................................")
		}
		if r.Method == http.MethodPost {
			log.Println("")
			log.Println("POST BODY PARAMS:")
			log.Println(".................................")
			_ = r.ParseMultipartForm(0)
			for bodyKey, bodyValues := range r.Form {
				log.Println(bodyKey + " : " + strings.Join(bodyValues, ","))
			}
		}
		log.Println("")
		log.Println("-----------------------------------------------------------------------")
		log.Println("REQUEST ENDED")
		log.Println("-----------------------------------------------------------------------")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
