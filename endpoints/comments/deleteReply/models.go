package deleteReply

import (
	"compose/commons"
	"net/http"
)

type RequestModel struct {
	ReplyId     string `validate:"required,id"`
	CommonModel *commons.CommonRequestModel
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}
	model := RequestModel{
		ReplyId:     commons.StrictSanitizeString(r.FormValue("reply_id")),
		CommonModel: commons.GetCommonRequestModel(r),
	}
	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}