package create

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
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
	if govalidator.StringLength(model.title, "1", "255") == false {
		return errors.New("Title length should be between 1 and 255")
	}
	if govalidator.StringLength(model.markdown, "1", "65536") == false {
		return errors.New("Article markdown should be between 1 and 65536")
	}
	if govalidator.StringLength(model.description, "0", "255") == false {
		return errors.New("Description should less than 255 characters")
	}

	return nil
}
