package likedArticles

import (
	"compose/commons"
	"errors"
	"net/http"
	"strings"
	"time"
)

type RequestModel struct {
	UserId            string
	MaxLikedAt        *time.Time
	CommonModel       *commons.CommonModel
	DefaultMaxLikedAt time.Time
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
	var DefaultMaxLikedAt, _ = time.ParseInLocation("2 Jan 2006 15:04:05", "2 Jan 3000 15:04:05", time.Now().Location())
	queryMap := r.URL.Query()
	model := RequestModel{
		UserId:            queryMap.Get("user_id"),
		CommonModel:       commons.GetCommonModel(r),
		MaxLikedAt:        &DefaultMaxLikedAt,
		DefaultMaxLikedAt: DefaultMaxLikedAt,
	}

	for key, values := range queryMap {
		value := strings.Join(values, "")
		if key == "max_liked_at" {
			maxLikedAt, err := time.ParseInLocation("2 Jan 2006 15:04:05", value, time.Now().Location())
			if commons.InError(err) {
				model.MaxLikedAt = nil
			} else {
				model.MaxLikedAt = &maxLikedAt
			}
		}
	}

	err2 := model.isInvalid()
	if commons.InError(err2) {
		return nil, err2
	}
	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if len(model.UserId) == 0 {
		return errors.New("User id can't be empty")
	}
	if model.MaxLikedAt == nil {
		return errors.New("Liked-at time isn't valid")
	}
	return nil
}
