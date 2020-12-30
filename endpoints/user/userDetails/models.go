package userDetails

import (
	"compose/commons"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestModel struct {
	UserId string `validate:"required,id"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	vars := mux.Vars(r)
	model := RequestModel{
		UserId: commons.StrictSanitizeString(vars["user_id"]),
	}

	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}

type ResponseModel struct {
	Status       commons.ResponseStatus `json:"status,omitempty"`
	Message      string                 `json:"message,omitempty"`
	Email        string                 `json:"email,omitempty"`
	Name         string                 `json:"name,omitempty"`
	Description  string                 `json:"description,omitempty"`
	PhotoUrl     string                 `json:"photo_url,omitempty"`
	ArticleCount uint64                 `json:"article_count"`
	LikeCount    uint64                 `json:"like_count"`
	MemberSince  string                 `json:"member_since,omitempty"`
	Editable     bool                   `json:"editable"`
}
