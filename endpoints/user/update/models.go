package update

import (
	"compose/commons"
	"net/http"
	"strings"
)

type RequestModel struct {
	UserId      string  `validate:"required,id"`
	NewUserId   *string `validate:"id"`
	Email       *string `validate:"email"`
	Name        *string `validate:"max=255"`
	Description *string `validate:"max=255"`
	PhotoUrl    *string `validate:"url"`
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
	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
