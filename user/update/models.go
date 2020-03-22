package update

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	UserId      string
	NewUserId   string
	Email       string
	Name        string
	Description string
	PhotoUrl    string
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{

		// Used to ensure operation security
		UserId: r.FormValue("user_id"),

		// Data points to be updated
		NewUserId:   r.FormValue("new_user_id"),
		Email:       r.FormValue("email"),
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		PhotoUrl:    r.FormValue("photo_url"),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if model.Name == "" {
		return errors.New("Name can't be empty")
	}
	if model.Email == "" {
		return errors.New("Email can't be empty")
	}
	if model.NewUserId == "" {
		return errors.New("User Id can't be empty")
	}
	if govalidator.StringLength(model.NewUserId, "1", "255") == false {
		return errors.New("User id should be between 1 and 255 characters")
	}
	if govalidator.IsEmail(model.Email) == false {
		return errors.New("Email isn't valid")
	}
	if govalidator.StringLength(model.Name, "1", "255") == false {
		return errors.New("Name should be not greater than 255 characters")
	}
	if govalidator.StringLength(model.Description, "0", "255") == false {
		return errors.New("Description should be not greater than 255 characters")
	}
	if govalidator.StringLength(model.PhotoUrl, "0", "255") == false {
		return errors.New("Photo url should be not greater than 255 characters")
	}

	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
