package update

import (
	"compose/commons"
	"errors"
	"net/http"
	"strings"
)

type RequestModel struct {
	ArticleId   string
	Title       *string
	Description *string
	Markdown    *string
	CommonModel *commons.CommonRequestModel
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

	for key, values := range r.Form {
		value := strings.Join(values, "")
		if key == "title" {
			model.Title = &value
		} else if key == "description" {
			model.Description = &value
		} else if key == "markdown" {
			model.Markdown = &value
		}
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {

	if model.Title != nil {
		if commons.IsEmpty(*model.Title) {
			return errors.New("Title can't be empty")
		}
		if commons.IsInvalidId(*model.Title) {
			return errors.New("Title should be between 1 and 255 characters")
		}
	}

	if model.Description != nil {
		if commons.IsEmpty(*model.Description) {
			return errors.New("Description can't be empty")
		}
		if commons.IsInvalidId(*model.Description) {
			return errors.New("Description cannot be greater than 255 characters")
		}
	}

	if model.Markdown != nil {
		if commons.IsEmpty(*model.Markdown) {
			return errors.New("Markdown can't be empty")
		}
		if commons.IsInvalidDataPoint(*model.Markdown) {
			return errors.New("Markdown cannot be greater than 65536 characters")
		}
	}

	return nil
}

type ResponseModel struct {
	Status    commons.ResponseStatus `json:"status,omitempty"`
	Message   string                 `json:"message,omitempty"`
	ArticleId string                 `json:"article_id,omitempty"`
}
