package createComment

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	ArticleId   string
	Markdown    string
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
		Markdown:    r.FormValue("markdown"),
		CommonModel: commons.GetCommonRequestModel(r),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}
	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if len(model.ArticleId) == 0 {
		return errors.New("ArticleId can't be empty")
	}
	if govalidator.StringLength(model.ArticleId, "1", "255") == false {
		return errors.New("ArticleId should be between 1 and 255")
	}

	if len(model.Markdown) == 0 {
		return errors.New("Markdown can't be empty")
	}
	if govalidator.StringLength(model.Markdown, "1", "65536") == false {
		return errors.New("Markdown should be between 1 and 65536")
	}
	return nil
}