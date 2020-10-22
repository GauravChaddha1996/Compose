package commons

import (
	"github.com/asaskevich/govalidator"
)

func IsEmpty(data string) bool {
	return len(data) == 0
}

func IsInvalidEmail(email string) bool {
	return !govalidator.IsEmail(email)
}

func IsInvalidId(data string) bool {
	return !govalidator.StringLength(data, "1", "255")
}

func IsInvalidDataPoint(data string) bool {
	return !govalidator.StringLength(data, "1", "65536")
}

func IsInvalidDataLength(data string, minLength int, maxLength int) bool {
	return !govalidator.StringLength(data, string(minLength), string(maxLength))
}
