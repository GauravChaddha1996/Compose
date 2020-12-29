package userDetails

import (
	"compose/commons"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestModel struct {
	userId string
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var err error
	vars := mux.Vars(r)
	model := RequestModel{
		userId: commons.StrictSanitizeString(vars["user_id"]),
	}

	err = model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}

	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsInvalidId(model.userId) {
		return errors.New("User id isn't valid")
	}
	return nil
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
