package createComment

import (
	"compose/commons"
	"net/http"
)

type RequestModel struct {
	ArticleId   string `validate:"required,id"`
	Markdown    string `validate:"required,max=65336"`
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
		ArticleId:   commons.StrictSanitizeString(r.FormValue("article_id")),
		Markdown:    commons.UgcSanitizeString(r.FormValue("markdown")),
		CommonModel: commons.GetCommonRequestModel(r),
	}

	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}
