package likeArticle

import (
	"compose/commons"
	"errors"
	"net/http"
)

type RequestModel struct {
	ArticleId   string
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
		ArticleId:   r.FormValue("article_id"),
		CommonModel: commons.GetCommonRequestModel(r),
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
	return nil
}
