package deleteComment

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	CommentId   string
	CommonModel *commons.CommonModel
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
		CommentId:   r.FormValue("comment_id"),
		CommonModel: commons.GetCommonModel(r),
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if govalidator.StringLength(model.CommentId, "1", "255") == false {
		return errors.New("Comment id can't be empty")
	}
	return nil
}