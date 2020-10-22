package createReply

import (
	"compose/commons"
	"errors"
	"net/http"
)

type RequestModel struct {
	ArticleId       string
	ParentId        string
	Markdown        string
	CommonModel     *commons.CommonRequestModel
	ParentIsComment bool
	ParentIsReply   bool
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
		ArticleId:       r.FormValue("article_id"),
		ParentId:        r.FormValue("parent_id"),
		Markdown:        r.FormValue("markdown"),
		CommonModel:     commons.GetCommonRequestModel(r),
		ParentIsComment: false,
		ParentIsReply:   false,
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}
	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsEmpty(model.ArticleId) {
		return errors.New("ArticleId can't be empty")
	}
	if commons.IsInvalidId(model.ArticleId) {
		return errors.New("ArticleId should be between 1 and 255")
	}
	if commons.IsEmpty(model.ParentId) {
		return errors.New("ParentId can't be empty")
	}
	if commons.IsInvalidId(model.ParentId) {
		return errors.New("ParentId should be between 1 and 255")
	}
	if commons.IsEmpty(model.Markdown) {
		return errors.New("Markdown can't be empty")
	}
	if commons.IsInvalidDataPoint(model.Markdown) {
		return errors.New("Markdown should be between 1 and 65536")
	}
	return nil
}
