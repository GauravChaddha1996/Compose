package update

import (
	"compose/commons"
	"net/http"
	"strings"
)

type RequestModel struct {
	ArticleId   string  `validate:"required,id"`
	Title       *string `validate:"max=255"`
	Description *string `validate:"max=255"`
	Markdown    *string `validate:"max=65536"`
	CommonModel *commons.CommonRequestModel
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		ArticleId:   commons.StrictSanitizeString(r.FormValue("article_id")),
		CommonModel: commons.GetCommonRequestModel(r),
	}

	for key, values := range r.Form {
		value := strings.Join(values, "")
		strictSanitizedValue := commons.StrictSanitizeString(value)
		if key == "title" {
			model.Title = &strictSanitizedValue
		} else if key == "description" {
			model.Description = &strictSanitizedValue
		} else if key == "markdown" {
			model.Markdown = &value
		}
	}
	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}

type ResponseModel struct {
	Status    commons.ResponseStatus `json:"status,omitempty"`
	Message   string                 `json:"message,omitempty"`
	ArticleId string                 `json:"article_id,omitempty"`
}
