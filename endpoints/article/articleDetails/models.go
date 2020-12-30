package articleDetails

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestModel struct {
	Id          string `validate:"required,id"`
	commonModel *commons.CommonRequestModel
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	vars := mux.Vars(r)
	model := RequestModel{
		Id:          commons.StrictSanitizeString(vars["article_id"]),
		commonModel: commons.GetCommonRequestModel(r),
	}

	err := commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}

type ResponseModel struct {
	Status       commons.ResponseStatus    `json:"status,omitempty"`
	Message      string                    `json:"message,omitempty"`
	Title        string                    `json:"title,omitempty"`
	Description  string                    `json:"description,omitempty"`
	Markdown     string                    `json:"markdown,omitempty"`
	CreatedAt    string                    `json:"created_at,omitempty"`
	LikeCount    uint64                    `json:"like_count"`
	CommentCount uint64                    `json:"comment_count"`
	PostedBy     apiEntity.SmallUserEntity `json:"posted_by,omitempty"`
	Editable     bool                      `json:"editable"`
}
