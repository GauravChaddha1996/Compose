package updateReply

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
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
	if govalidator.StringLength(model.ReplyId, "1", "255") == false {
		return errors.New("ReplyId should be between 1 and 255 characters")
	}
	if govalidator.StringLength(model.Markdown, "1", "65536") == false {
		return errors.New("Reply markdown should be between 1 and 65536 characters")
	}
	return nil
}
