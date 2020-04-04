package postedArticles

import (
	"compose/commons"
	"errors"
	"net/http"
	"strings"
	"time"
)

type RequestModel struct {
	UserId              string
	MaxCreatedAt        *time.Time
	CommonModel         *commons.CommonModel
	DefaultMaxCreatedAt time.Time
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
	var DefaultMaxCreatedAt, _ = time.ParseInLocation("2 Jan 2006 15:04:05", "2 Jan 3000 15:04:05", time.Now().Location())
	queryMap := r.URL.Query()
	model := RequestModel{
		UserId:              queryMap.Get("user_id"),
		CommonModel:         commons.GetCommonModel(r),
		MaxCreatedAt:        &DefaultMaxCreatedAt,
		DefaultMaxCreatedAt: DefaultMaxCreatedAt,
	}

	for key, values := range queryMap {
		value := strings.Join(values, "")
		if key == "max_created_at" {
			maxCreatedAt, err := time.ParseInLocation("2 Jan 2006 15:04:05", value, time.Now().Location())
			if commons.InError(err) {
				model.MaxCreatedAt = nil
			} else {
				model.MaxCreatedAt = &maxCreatedAt
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
	if model.MaxCreatedAt == nil {
		return errors.New("Created_at time isn't valid")
	}
	return nil
}
