package update

import (
	"compose/commons"
	"errors"
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
		UserId: commons.StrictSanitizeString(r.FormValue("user_id")),
	}

	for key, values := range r.Form {
		value := strings.Join(values, "")
		strictSanitizedValue := commons.StrictSanitizeString(value)
		ugcSanitizedValue := commons.UgcSanitizeString(value)
		if key == "new_user_id" {
			model.NewUserId = &strictSanitizedValue
		} else if key == "email" {
			model.Email = &strictSanitizedValue
		} else if key == "name" {
			model.Name = &strictSanitizedValue
		} else if key == "description" {
			model.Description = &ugcSanitizedValue
		} else if key == "photo_url" {
			model.PhotoUrl = &ugcSanitizedValue
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
		if commons.IsEmpty(*model.NewUserId) {
			return errors.New("New user Id can't be empty")
		}
		if commons.IsInvalidId(*model.NewUserId) {
			return errors.New("User id should be between 1 and 255 characters")
		}
	}

	if model.Name != nil {
		if commons.IsEmpty(*model.Name) {
			return errors.New("Name can't be empty")
		}
		if commons.IsInvalidId(*model.Name) {
			return errors.New("Name should be not greater than 255 characters")
		}
	}

	if model.Email != nil {
		if commons.IsEmpty(*model.Email) {
			return errors.New("Email can't be empty")
		}

		if commons.IsInvalidEmail(*model.Email) {
			return errors.New("Email isn't valid")
		}
	}

	if model.Description != nil {
		if commons.IsInvalidDataLength(*model.Description, 0, 255) {
			return errors.New("Description should be not greater than 255 characters")
		}
	}

	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
