package commons

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate
var Translator ut.Translator

func initValidator() {
	t := en.New()
	uni := ut.New(t, t)
	translator, _ := uni.GetTranslator("en")

	Translator = translator
	Validator = validator.New()
	registerValidations()
	registerTranslations()
}

func registerValidations() {
	err := Validator.RegisterValidation("id", func(fl validator.FieldLevel) bool {
		data := fl.Field().String()
		result := govalidator.StringLength(data, "1", "255")
		return result
	})
	PanicIfError(err)
}

func registerTranslations() {
	err := Validator.RegisterTranslation("id", Translator, func(ut ut.Translator) error {
		return ut.Add("id", "{0} must be a valid id", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("id", fe.Field())
		return t
	})
	PanicIfError(err)
	err = Validator.RegisterTranslation("required", Translator, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	PanicIfError(err)
	err = Validator.RegisterTranslation("max", Translator, func(ut ut.Translator) error {
		return ut.Add("max", "{0} can't be longer than {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})
	PanicIfError(err)
}

func GetValidationError(err error) error {
	errString := ""
	for i, e := range err.(validator.ValidationErrors) {
		if i != 0 {
			errString += ", "
		}
		errString += e.Translate(Translator)
	}
	return errors.New(errString)
}

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
