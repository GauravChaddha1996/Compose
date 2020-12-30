package likedArticles

import (
	"compose/commons"
	"net/http"
	"time"
)

type RequestModel struct {
	UserId            string                      `validate:"required,id"`
	MaxLikedAt        *time.Time                  `validate:"required"`
	CommonModel       *commons.CommonRequestModel `validate:"required"`
	DefaultMaxLikedAt time.Time                   `validate:"required"`
}

type ResponseModel struct {
	Status          commons.ResponseStatus `json:"status,omitempty"`
	Message         string                 `json:"message,omitempty"`
	LikedArticles   []LikedArticle         `json:"liked_articles,omitempty"`
	MaxLikedAt      string                 `json:"max_liked_at,omitempty"`
	HasMoreArticles bool                   `json:"has_more_articles"`
}

type LikedArticle struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var DefaultMaxLikedAt = commons.MaxTime
	queryMap := r.URL.Query()

	var maxLikedAtTime *time.Time
	maxLikedAt, err := commons.ParseTime(queryMap.Get("max_liked_at"))
	if commons.InError(err) {
		maxLikedAtTime = &DefaultMaxLikedAt
	} else {
		maxLikedAtTime = &maxLikedAt
	}
	model := RequestModel{
		UserId:            commons.StrictSanitizeString(queryMap.Get("user_id")),
		CommonModel:       commons.GetCommonRequestModel(r),
		MaxLikedAt:        maxLikedAtTime,
		DefaultMaxLikedAt: DefaultMaxLikedAt,
	}
	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}
