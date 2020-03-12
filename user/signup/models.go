package signup

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"unicode"
)

type RequestModel struct {
	Email    string
	Name     string
	Password string
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	//todo learn how to use enum in golang
	model := RequestModel{
		Email:    r.FormValue("email"),
		Name:     r.FormValue("name"),
		Password: r.FormValue("password"),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (requestModel RequestModel) isInvalid() error {

	if govalidator.StringLength(requestModel.Name, "1", "255") == false {
		return errors.New("Name should be between 1 and 255 characters")
	}
	if govalidator.IsEmail(requestModel.Email) == false {
		return errors.New("Email isn't valid")
	}

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
		}
	}

	if !(hasNumber && hasLowerChar && hasUpperChar && hasSpecialChar) {
		return errors.New("Password must have at-least one lowercase, one uppercase, one number and one special character")
	}
	return nil
}

// todo add rules here for null
type ResponseModel struct {
	Status      commons.ResponseStatus `json:"Status"`
	Message     string                 `json:"Message"`
	AccessToken string                 `json:"AccessToken"`
}
