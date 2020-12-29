package articleLikes

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"net/http"
	"time"
)

type RequestModel struct {
	ArticleId         string                      `validate:"required,id"`
	MaxLikedAt        *time.Time                  `validate:"required"`
	CommonModel       *commons.CommonRequestModel `validate:"required"`
	DefaultMaxLikedAt time.Time                   `validate:"required"`
}

type ResponseModel struct {
	Status       commons.ResponseStatus      `json:"status,omitempty"`
	Message      string                      `json:"message,omitempty"`
	LikedByUsers []*apiEntity.SmallUserEntity `json:"liked_by_users,omitempty"`
	MaxLikedAt   string                      `json:"max_liked_at,omitempty"`
	HasMoreLikes bool                        `json:"has_more_likes"`
}

type LikedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
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
		ArticleId:         commons.StrictSanitizeString(queryMap.Get("article_id")),
		MaxLikedAt:        maxLikedAtTime,
		DefaultMaxLikedAt: DefaultMaxLikedAt,
		CommonModel:       commons.GetCommonRequestModel(r),
	}

	err = commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}
