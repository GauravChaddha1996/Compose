package deleteReply

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	ReplyId   string
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
		ReplyId:   r.FormValue("reply_id"),
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
		return errors.New("Reply id can't be empty")
	}
	return nil
}
