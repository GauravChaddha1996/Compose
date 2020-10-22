package create

import (
	"compose/commons"
	"errors"
	"net/http"
	"strings"
)

type RequestModel struct {
	userId      string
	title       string
	description string
	markdown    string
}

type ResponseModel struct {
	Status    commons.ResponseStatus `json:"status,omitempty"`
	Message   string                 `json:"message,omitempty"`
	ArticleId string                 `json:"article_id,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	commonModel := commons.GetCommonRequestModel(r)
	model := RequestModel{
		userId:   commonModel.UserId,
		title:    r.FormValue("title"),
		markdown: r.FormValue("markdown"),
	}

	for key, values := range r.Form {
		value := strings.Join(values, "")
		if key == "description" {
			model.description = value
		}
	}
	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsInvalidId(model.title) {
		return errors.New("Title length should be between 1 and 255")
	}
	if commons.IsInvalidDataPoint(model.markdown) {
		return errors.New("Article markdown should be between 1 and 65536")
	}
	if commons.IsInvalidId(model.description) {
		return errors.New("Description should less than 255 characters")
	}

	return nil
}
