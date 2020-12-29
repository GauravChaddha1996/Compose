package delete

import (
	"compose/commons"
	"errors"
	"net/http"
)

type RequestModel struct {
	ArticleId   string
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

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsInvalidId(model.ArticleId) {
		return errors.New("Article id can't be empty")
	}
	return nil
}
