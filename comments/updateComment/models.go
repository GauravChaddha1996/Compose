package updateComment

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	CommentId   string
	Markdown    string
	CommonModel *commons.CommonModel
}

type ResponseModel struct {
	Status    commons.ResponseStatus `json:"status,omitempty"`
	Message   string                 `json:"message,omitempty"`
	CommentId string                 `json:"comment_id,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	err := r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}
	model := RequestModel{
		CommentId:   r.FormValue("comment_id"),
		Markdown:    r.FormValue("markdown"),
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
		return errors.New("CommentId should be between 1 and 255 characters")
	}
	if govalidator.StringLength(model.Markdown, "1", "65536") == false {
		return errors.New("Comment markdown should be between 1 and 65536 characters")
	}
	return nil
}
