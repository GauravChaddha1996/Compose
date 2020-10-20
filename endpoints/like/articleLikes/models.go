package articleLikes

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strings"
	"time"
)

type RequestModel struct {
	ArticleId         string
	MaxLikedAt        *time.Time
	CommonModel       *commons.CommonRequestModel
	DefaultMaxLikedAt time.Time
}

type ResponseModel struct {
	Status       commons.ResponseStatus `json:"status,omitempty"`
	Message      string                 `json:"message,omitempty"`
	LikedByUsers []LikedByUser          `json:"liked_by_users,omitempty"`
	MaxLikedAt   string                 `json:"max_liked_at,omitempty"`
	HasMoreLikes bool                   `json:"has_more_likes"`
}

type LikedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var DefaultMaxLikedAt, _ = commons.MaxTime()
	queryMap := r.URL.Query()
	model := RequestModel{
		ArticleId:         queryMap.Get("article_id"),
		MaxLikedAt:        &DefaultMaxLikedAt,
		DefaultMaxLikedAt: DefaultMaxLikedAt,
		CommonModel:       commons.GetCommonRequestModel(r),
	}

	for key, values := range queryMap {
		value := strings.Join(values, "")
		if key == "max_liked_at" {
			maxLikedAt, err := commons.ParseTime(value)
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
	if len(model.ArticleId) == 0 {
		return errors.New("ArticleId can't be empty")
	}
	if govalidator.StringLength(model.ArticleId, "1", "255") == false {
		return errors.New("ArticleId should be between 1 and 255")
	}
	if model.MaxLikedAt == nil {
		return errors.New("Liked-at time isn't valid")
	}
	return nil
}
