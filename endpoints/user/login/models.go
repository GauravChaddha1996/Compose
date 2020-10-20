package login

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	email    string
	password string
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		email:    r.FormValue("email"),
		password: r.FormValue("password"),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if govalidator.IsEmail(model.email) == false {
		return errors.New("Email isn't valid")
	}
	if len(model.password) == 0 {
		return errors.New("Password can't be empty")
	}
	return nil
}

type ResponseModel struct {
	Status      commons.ResponseStatus `json:"status,omitempty"`
	Message     string                 `json:"message,omitempty"`
	AccessToken string                 `json:"access_token,omitempty"`
}
