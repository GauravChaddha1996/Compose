package replyThread

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RequestModel struct {
	ArticleId   string                      `validate:"required,id"`
	ParentId    string                      `validate:"required,id"`
	CreatedAt   *time.Time                  `validate:"required"`
	ReplyCount  int                         `validate:"required"`
	CommonModel *commons.CommonRequestModel `validate:"required"`
}

type ResponseModel struct {
	Status  commons.ResponseStatus              `json:"status,omitempty"`
	Message string                              `json:"message,omitempty"`
	Parent  *apiEntity.CommentReplyParentEntity `json:"parent,omitempty"`
	Replies []*apiEntity.ReplyEntity            `json:"replies,omitempty"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	queryMap := r.URL.Query()
	model := RequestModel{
		ArticleId:   "",
		ParentId:    "",
		CommonModel: commons.GetCommonRequestModel(r),
	}
	postbackParamsStr := queryMap.Get("postback_params")
	if len(postbackParamsStr) != 0 {
		postbackParamsStr = strings.ReplaceAll(postbackParamsStr, "\\\"", "\"")
		postbackParamsMap := make(map[string]string)
		err := json.Unmarshal([]byte(postbackParamsStr), &postbackParamsMap)
		if commons.InError(err) {
			return nil, errors.New("Postback params are faulty")
		}
		model.ParentId = commons.StrictSanitizeString(postbackParamsMap["parent_id"])
		model.ArticleId = commons.StrictSanitizeString(postbackParamsMap["article_id"])
		model.CreatedAt = getCreatedAtTimeFromPostbackParams(postbackParamsMap)
		model.ReplyCount, err = strconv.Atoi(postbackParamsMap["reply_count"])
		if commons.InError(err) {
			model.ReplyCount = 0
		}
	}
	err := commons.Validator.Struct(model)
	if commons.InError(err) {
		return nil, commons.GetValidationError(err)
	}
	return &model, nil
}

func getCreatedAtTimeFromPostbackParams(model map[string]string) *time.Time {
	createdAt := model["created_at"]
	createdAtTime, err := commons.ParseTime(createdAt)
	if commons.InError(err) {
		return &commons.MaxTime
	} else {
		return &createdAtTime
	}
}
