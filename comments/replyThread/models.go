package replyThread

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"encoding/json"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RequestModel struct {
	ArticleId      string
	ParentId       string
	CreatedAt      *time.Time
	ReplyCount     int
	PostbackParams map[string]string
	CommonModel    *commons.CommonModel
}

type ResponseModel struct {
	Status  commons.ResponseStatus        `json:"status,omitempty"`
	Message string                        `json:"message,omitempty"`
	Parent  *commentCommons.ParentEntity  `json:"parent,omitempty"`
	Replies []*commentCommons.ReplyEntity `json:"replies,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	queryMap := r.URL.Query()
	model := RequestModel{
		ArticleId:      "",
		ParentId:       "",
		PostbackParams: nil,
		CommonModel:    commons.GetCommonModel(r),
	}
	postbackParamsStr := queryMap.Get("postback_params")
	if len(postbackParamsStr) != 0 {
		postbackParamsStr = strings.ReplaceAll(postbackParamsStr, "\\\"", "\"")
		postbackParamsMap := make(map[string]string)
		err := json.Unmarshal([]byte(postbackParamsStr), &postbackParamsMap)
		if commons.InError(err) {
			return nil, errors.New("Postback params are faulty")
		}
		model.PostbackParams = postbackParamsMap
		model.ParentId = postbackParamsMap["parent_id"]
		model.ArticleId = postbackParamsMap["article_id"]
		model.CreatedAt = getCreatedAtTimeFromPostbackParams(&model)
		model.ReplyCount, err = strconv.Atoi(postbackParamsMap["reply_count"])
		if commons.InError(err) {
			model.ReplyCount = 0
		}
	}
	err := model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}
	return &model, nil
}

func getCreatedAtTimeFromPostbackParams(model *RequestModel) *time.Time {
	maxTime, _ := commons.MaxTime()
	createdAt := model.PostbackParams["created_at"]
	createdAtTime, err := commons.ParseTime(createdAt)
	if commons.InError(err) {
		return &maxTime
	} else {
		return &createdAtTime
	}
}

func (model RequestModel) isInvalid() error {
	if len(model.ArticleId) == 0 {
		return errors.New("ArticleId can't be empty")
	}
	if govalidator.StringLength(model.ArticleId, "1", "255") == false {
		return errors.New("ArticleId should be between 1 and 255")
	}
	if len(model.ParentId) == 0 {
		return errors.New("ParentId can't be empty")
	}
	if govalidator.StringLength(model.ParentId, "1", "255") == false {
		return errors.New("ParentId should be between 1 and 255")
	}
	return nil
}
