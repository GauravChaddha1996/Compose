package articleLikes

import (
	"compose/commons"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strconv"
	"strings"
)

type RequestModel struct {
	ArticleId         string
	LastLikeId        *string
	CommonModel       *commons.CommonModel
	DefaultLastLikeId string
}

type ResponseModel struct {
	Status       commons.ResponseStatus `json:"status,omitempty"`
	Message      string                 `json:"message,omitempty"`
	LikedByUsers []LikedByUser          `json:"liked_by_users,omitempty"`
	LastLikeId   string                 `json:"last_like_id,omitempty"`
	HasMoreLikes bool                   `json:"has_more_likes"`
}

type LikedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	var DefaultLastLikeId = "0"
	queryMap := r.URL.Query()
	model := RequestModel{
		ArticleId:         queryMap.Get("article_id"),
		LastLikeId:        &DefaultLastLikeId,
		DefaultLastLikeId: DefaultLastLikeId,
		CommonModel:       commons.GetCommonModel(r),
	}

	for key, values := range queryMap {
		value := strings.Join(values, "")
		if key == "last_like_id" {
			lastLikeId := value
			_, err := strconv.ParseUint(lastLikeId, 10, 64)
			if commons.InError(err) {
				model.LastLikeId = nil
			} else {
				model.LastLikeId = &lastLikeId
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
	if model.LastLikeId == nil {
		return errors.New("Last Like id isn't valid")
	}
	return nil
}
