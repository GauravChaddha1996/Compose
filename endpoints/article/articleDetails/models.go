package articleDetails

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestModel struct {
	id          string
	commonModel *commons.CommonRequestModel
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	vars := mux.Vars(r)
	model := RequestModel{
		id:          vars["article_id"],
		commonModel: commons.GetCommonRequestModel(r),
	}

	err := model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if len(model.id) == 0 {
		return errors.New("Article id can't be empty")
	}
	return nil
}

type ResponseModel struct {
	Status       commons.ResponseStatus `json:"status,omitempty"`
	Message      string                 `json:"message,omitempty"`
	Title        string                 `json:"title,omitempty"`
	Description  string                 `json:"description,omitempty"`
	Markdown     string                 `json:"markdown,omitempty"`
	CreatedAt    string                 `json:"created_at,omitempty"`
	LikeCount    uint64                 `json:"like_count"`
	CommentCount uint64                 `json:"comment_count"`
	PostedBy     apiEntity.PostedByUser `json:"posted_by,omitempty"`
	Editable     bool                   `json:"editable"`
}
