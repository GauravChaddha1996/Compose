package update

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strings"
)

type RequestModel struct {
	UserId      string
	NewUserId   *string
	Email       *string
	Name        *string
	Description *string
	PhotoUrl    *string
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		UserId: r.FormValue("user_id"),
	}

	for key, values := range r.Form {
		value := strings.Join(values, "")
		if key == "new_user_id" {
			model.NewUserId = &value
		} else if key == "email" {
			model.Email = &value
		} else if key == "name" {
			model.Name = &value
		} else if key == "description" {
			model.Description = &value
		} else if key == "photo_url" {
			model.PhotoUrl = &value
		}
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {

	if model.NewUserId != nil {
		if *model.NewUserId == "" {
			return errors.New("New user Id can't be empty")
		}
		if govalidator.StringLength(*model.NewUserId, "1", "255") == false {
			return errors.New("User id should be between 1 and 255 characters")
		}
	}

	if model.Name != nil {
		if *model.Name == "" {
			return errors.New("Name can't be empty")
		}
		if govalidator.StringLength(*model.Name, "1", "255") == false {
			return errors.New("Name should be not greater than 255 characters")
		}
	}

	if model.Email != nil {
		if *model.Email == "" {
			return errors.New("Email can't be empty")
		}

		if govalidator.IsEmail(*model.Email) == false {
			return errors.New("Email isn't valid")
		}
	}

	if model.Description != nil {
		if govalidator.StringLength(*model.Description, "0", "255") == false {
			return errors.New("Description should be not greater than 255 characters")
		}
	}

	if model.PhotoUrl != nil {
		if govalidator.StringLength(*model.PhotoUrl, "0", "255") == false {
			return errors.New("Photo url should be not greater than 255 characters")
		}
	}

	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
