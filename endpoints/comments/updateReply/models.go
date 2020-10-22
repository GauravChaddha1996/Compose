package updateReply

import (
	"compose/commons"
	"errors"
	"net/http"
)

type RequestModel struct {
	ReplyId     string
	Markdown    string
	CommonModel *commons.CommonRequestModel
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
	ReplyId string                 `json:"reply_id,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	err := r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}
	model := RequestModel{
		ReplyId:     r.FormValue("reply_id"),
		Markdown:    r.FormValue("markdown"),
		CommonModel: commons.GetCommonRequestModel(r),
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}
	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsInvalidId(model.ReplyId) {
		return errors.New("ReplyId should be between 1 and 255 characters")
	}
	if commons.IsInvalidDataPoint(model.Markdown) {
		return errors.New("Reply markdown should be between 1 and 65536 characters")
	}
	return nil
}
