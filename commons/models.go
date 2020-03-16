package commons

import "net/http"

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

type CommonHeaders struct {
	AccessToken string
}

func GetCommonHeaders(r *http.Request) *CommonHeaders {
	headers := r.Header
	return &CommonHeaders{
		AccessToken: headers.Get("access_token"),
	}
}
