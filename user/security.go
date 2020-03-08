package user

import "net/http"

func IsSignupRequestSecure(_ *http.Request) (bool, string) {
	return true, ""
}
