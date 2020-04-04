package delete

import (
	"compose/commons"
	"compose/dbModels"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	email string
	user  *dbModels.User
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		email: r.FormValue("email"),
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
	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
