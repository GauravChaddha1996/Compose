package commons

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

var securityMiddlewarePathConfigMap = make(map[string]*SecurityMiddlewarePathConfig)

type SecurityMiddlewarePathConfig struct {
	CheckAccessToken bool
	CheckUserId      bool
	CheckUserEmail   bool
}

func getDefaultSecurityMiddlewarePathConfig() *SecurityMiddlewarePathConfig {
	return &SecurityMiddlewarePathConfig{
		CheckAccessToken: true,
		CheckUserId:      true,
		CheckUserEmail:   true,
	}
}

func AddSecurityMiddlewarePathConfig(path string, config *SecurityMiddlewarePathConfig) {
	securityMiddlewarePathConfigMap[path] = config
}

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

func ExtractCommonModelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commonHeaders := makeCommonRequestModel(r)
		parentContext := r.Context()
		newContext := context.WithValue(parentContext, CommonModelKey, commonHeaders)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}

func GeneralSecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		commonModel := r.Context().Value(CommonModelKey).(*CommonRequestModel)
		securityError := ensureSecurity(commonModel, securityMiddlewarePathConfigMap[r.URL.Path])
		if securityError != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte("{\"message\":\"User id associated isn't known\"}"))
			PanicIfError(err)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func ensureSecurity(commonHeaders *CommonRequestModel, config *SecurityMiddlewarePathConfig) error {
	if config == nil {
		config = getDefaultSecurityMiddlewarePathConfig()
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

func CommonResponseHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func ResponseLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		newWriter := httptest.NewRecorder()
		next.ServeHTTP(newWriter, r)
		_, err := w.Write(newWriter.Body.Bytes())
		PanicIfError(err)

		println("")
		log.Println("----------------------RESPONSE RECORDED-------------------------")
		log.Println(r.Method + " " + r.URL.String())
		println("")
		log.Println(newWriter.Body.String())
		println("")

	})
}
