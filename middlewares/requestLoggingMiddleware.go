package middlewares

import (
	"log"
	"net/http"
	"strings"
)

func RequestLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("----------------------REQUEST RECEIVED-------------------------")
		log.Println(r.Method + " " + r.URL.String())
		log.Println("----------------------HEADERS----------------------------------")
		for headerKey, headerValues := range r.Header {
			log.Println(headerKey + " : " + strings.Join(headerValues, ","))
		}
		if r.Method == http.MethodGet {
			queryMap := r.URL.Query()
			if len(queryMap) > 0 {
				log.Println("----------------------QUERY PARAMS-------------------------")
				for key, value := range queryMap {
					log.Println(key + " : " + strings.Join(value, ","))
				}
			}
		}
		if r.Method == http.MethodPost {
			_ = r.ParseMultipartForm(1024)
			form := r.Form
			if len(form) > 0 {
				log.Println("----------------------POST BODY PARAMS---------------------")
				for bodyKey, bodyValues := range r.Form {
					log.Println(bodyKey + " : " + strings.Join(bodyValues, ","))
				}
			}
		}
		log.Println("----------------------REQUEST ENDED-------------------------")
		println("")

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}