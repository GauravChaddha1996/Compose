package user

import (
	"regexp"
)

var re = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// todo use validation library and or improve this code
func IsUserSignupRequestValid(requestModel *UserSignupRequestModel) (bool, string) {
	isValid := true
	message := ""

	if len(requestModel.Name) == 0 {
		isValid = false
		message = SIGNUP_ERROR_VALIDITY_NAME_MESSAGE
	}
	if isValid && re.MatchString(requestModel.Email) == false {
		isValid = false
		message = SIGNUP_ERROR_VALIDITY_EMAIL_MESSAGE
	}
	return isValid, message
}
