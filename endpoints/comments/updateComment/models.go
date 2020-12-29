package updateComment

import (
	"compose/commons"
	"errors"
	"net/http"
)

type RequestModel struct {
	CommentId   string
	Markdown    string
	CommonModel *commons.CommonRequestModel
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
		CommentId:   commons.StrictSanitizeString(r.FormValue("comment_id")),
		Markdown:    commons.UgcSanitizeString(r.FormValue("markdown")),
		CommonModel: commons.GetCommonRequestModel(r),
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}
	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsInvalidId(model.CommentId) {
		return errors.New("CommentId should be between 1 and 255 characters")
	}
	if commons.IsInvalidDataPoint(model.Markdown) == false {
		return errors.New("Comment markdown should be between 1 and 65536 characters")
	}
	return nil
}
