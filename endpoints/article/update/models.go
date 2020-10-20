package update

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
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
		if *model.Title == "" {
			return errors.New("Title can't be empty")
		}
		if govalidator.StringLength(*model.Title, "1", "255") == false {
			return errors.New("Title should be between 1 and 255 characters")
		}
	}

	if model.Description != nil {
		if *model.Description == "" {
			return errors.New("Description can't be empty")
		}
		if govalidator.StringLength(*model.Description, "1", "255") == false {
			return errors.New("Description cannot be greater than 255 characters")
		}
	}

	if model.Markdown != nil {
		if *model.Markdown == "" {
			return errors.New("Markdown can't be empty")
		}
		if govalidator.StringLength(*model.Markdown, "1", "65536") == false {
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
