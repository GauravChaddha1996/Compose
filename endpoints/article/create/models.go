package create

import (
	"compose/commons"
	"net/http"
	"strings"
)

type RequestModel struct {
	UserId      string `validate:"required,id"`
	Title       string `validate:"required,max=255"`
	Description string `validate:"max=255"`
	Markdown    string `validate:"required,max=65536"`
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
		UserId:   commonModel.UserId,
		Title:    r.FormValue("title"),
		Markdown: r.FormValue("markdown"),
	}
	for key, values := range r.Form {
		value := strings.Join(values, "")
		if key == "description" {
			model.Description = value
		}
	}

	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}

	return &model, nil
}
