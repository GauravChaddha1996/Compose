package delete

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
)

type RequestModel struct {
	ArticleId   string
	MarkdownId  string
	CommonModel *commons.CommonModel
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	err = r.ParseMultipartForm(1024)
	if commons.InError(err) {
		return nil, err
	}

	model := RequestModel{
		ArticleId:   r.FormValue("article_id"),
		MarkdownId:  "",
		CommonModel: commons.GetCommonModel(r),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if govalidator.StringLength(model.ArticleId, "1", "255") == false {
		return errors.New("Article id can't be empty")
	}
	return nil
}
