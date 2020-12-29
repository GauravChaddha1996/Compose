package delete

import (
	"compose/commons"
	"errors"
	"net/http"
)

type RequestModel struct {
	email       string
	commonModel *commons.CommonRequestModel
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		email:       commons.StrictSanitizeString(r.FormValue("email")),
		commonModel: commons.GetCommonRequestModel(r),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsInvalidEmail(model.email) {
		return errors.New("Email isn't valid")
	}
	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
