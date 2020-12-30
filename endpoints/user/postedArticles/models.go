package postedArticles

import (
	"compose/commons"
	"net/http"
	"time"
)

type RequestModel struct {
	UserId              string                      `validate:"required,id"`
	MaxCreatedAt        *time.Time                  `validate:"required"`
	CommonModel         *commons.CommonRequestModel `validate:"required"`
	DefaultMaxCreatedAt time.Time                   `validate:"required"`
}

type ResponseModel struct {
	Status          commons.ResponseStatus `json:"status,omitempty"`
	Message         string                 `json:"message,omitempty"`
	PostedArticles  []PostedArticle        `json:"posted_articles,omitempty"`
	MaxCreatedAt    string                 `json:"max_created_at,omitempty"`
	HasMoreArticles bool                   `json:"has_more_articles"`
}

type PostedArticle struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var DefaultMaxCreatedAt = commons.MaxTime
	queryMap := r.URL.Query()
	var maxCreatedAtTime *time.Time
	maxCreatedAt, err := commons.ParseTime(queryMap.Get("max_created_at"))
	if commons.InError(err) {
		maxCreatedAtTime = &DefaultMaxCreatedAt
	} else {
		maxCreatedAtTime = &maxCreatedAt
	}
	model := RequestModel{
		UserId:              commons.StrictSanitizeString(queryMap.Get("user_id")),
		CommonModel:         commons.GetCommonRequestModel(r),
		MaxCreatedAt:        maxCreatedAtTime,
		DefaultMaxCreatedAt: DefaultMaxCreatedAt,
	}

	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}
