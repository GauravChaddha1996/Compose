package user

import (
	"github.com/asaskevich/govalidator"
	"unicode"
)

// todo use validation library and or improve this code
// todo retturn nil message here somehow
func IsUserSignupRequestValid(requestModel *SignupRequestModel) (bool, string) {
	isValid := true
	message := ""

	if govalidator.StringLength(requestModel.Name, "1", "255") == false {
		isValid = false
		message = ERROR_VALIDITY_NAME_MESSAGE
	}
	if isValid && govalidator.IsEmail(requestModel.Email) == false {
		isValid = false
		message = ERROR_VALIDITY_EMAIL_MESSAGE
	}

	if isValid {
		hasNumber := false
		hasLowerChar := false
		hasUpperChar := false
		hasSpecialChar := false
		for _, char := range requestModel.Password {
			switch {
			case unicode.IsNumber(char) || unicode.IsDigit(char):
				hasNumber = true
			case unicode.IsLower(char):
				hasLowerChar = true
			case unicode.IsUpper(char):
				hasUpperChar = true
			case unicode.IsSymbol(char) || unicode.IsPunct(char) || unicode.IsMark(char):
				hasSpecialChar = true
			default:

			}
		}

		if !(hasNumber && hasLowerChar && hasUpperChar && hasSpecialChar) {
			isValid = false
			message = ERROR_VALIDITY_PASSWORD_MESSAGE
		}
	}
	return isValid, message
}
