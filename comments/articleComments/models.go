package articleComments

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strings"
	"time"
)

type RequestModel struct {
	ArticleId           string
	MaxCreatedAt        *time.Time
	CommonModel         *commons.CommonModel
	DefaultMaxCreatedAt time.Time
}

type ResponseModel struct {
	Status          commons.ResponseStatus         `json:"status,omitempty"`
	Message         string                         `json:"message,omitempty"`
	Comments        []commentCommons.CommentEntity `json:"comments,omitempty"`
	MaxCreatedAt    string                         `json:"max_created_at,omitempty"`
	HasMoreComments bool                           `json:"has_more_comments"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var DefaultMaxCreatedAt, _ = time.ParseInLocation("2 Jan 2006 15:04:05", "2 Jan 3000 15:04:05", time.Now().Location())
	queryMap := r.URL.Query()
	model := RequestModel{
		ArticleId:           queryMap.Get("article_id"),
		MaxCreatedAt:        &DefaultMaxCreatedAt,
		DefaultMaxCreatedAt: DefaultMaxCreatedAt,
		CommonModel:         commons.GetCommonModel(r),
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
	if len(model.ArticleId) == 0 {
		return errors.New("ArticleId can't be empty")
	}
	if govalidator.StringLength(model.ArticleId, "1", "255") == false {
		return errors.New("ArticleId should be between 1 and 255")
	}
	if model.MaxCreatedAt == nil {
		return errors.New("Max created-at time isn't valid")
	}
	return nil
}
