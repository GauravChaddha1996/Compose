package delete

import (
	"compose/commons"
	"net/http"
)

type RequestModel struct {
	Email       string `validate:"required,email"`
	CommonModel *commons.CommonRequestModel
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		Email:       commons.StrictSanitizeString(r.FormValue("email")),
		CommonModel: commons.GetCommonRequestModel(r),
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
