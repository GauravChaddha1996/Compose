package articleComments

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type RequestModel struct {
	ArticleId      string
	PostbackParams *PostbackParams
	CommonModel    *commons.CommonRequestModel
}

type ResponseModel struct {
	Status         commons.ResponseStatus     `json:"status,omitempty"`
	Message        string                     `json:"message,omitempty"`
	Comments       []*apiEntity.CommentEntity `json:"comments,omitempty"`
	PostbackParams string                     `json:"postback_params,omitempty"`
	HasMore        bool                       `json:"has_more"`
}

type PostbackParams struct {
	Count     int    `json:"count"`
	CreatedAt string `json:"created_at"`
}

func getRequestModel(r *http.Request) (*RequestModel, error) {
	queryMap := r.URL.Query()
	model := RequestModel{
		ArticleId:      commons.StrictSanitizeString(queryMap.Get("article_id")),
		PostbackParams: nil,
		CommonModel:    commons.GetCommonRequestModel(r),
	}
	postbackParamsStr := queryMap.Get("postback_params")
	if len(postbackParamsStr) != 0 {
		postbackParamsStr = strings.ReplaceAll(postbackParamsStr, "\\\"", "\"")
		var postbackParams PostbackParams
		err := json.Unmarshal([]byte(postbackParamsStr), &postbackParams)
		if commons.InError(err) {
			return nil, errors.New("Postback params are faulty")
		}
		model.PostbackParams = &postbackParams
	}
	err := model.isInvalid()
	if commons.InError(err) {
		return nil, err
	}
	return &model, nil
}

func (model RequestModel) isInvalid() error {
	if commons.IsEmpty(model.ArticleId) {
		return errors.New("ArticleId can't be empty")
	}
	if commons.IsInvalidId(model.ArticleId) {
		return errors.New("ArticleId should be between 1 and 255")
	}
	return nil
}
