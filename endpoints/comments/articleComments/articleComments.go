package articleComments

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"compose/dataLayer/daos"
	userDaos "compose/dataLayer/daos/user"
	"compose/endpoints/comments/commentCommons"
	"compose/endpoints/comments/replyThreadCommon"
	"encoding/json"
	"errors"
)

const ArticleCommentLimit = 20
const MaxRepliesCount = 1000
const MaxCommentReplyLevel = 2

func getArticleComments(model *RequestModel) (*ResponseModel, error) {
	replyDao := daos.GetReplyDao()
	userDao := daos.GetUserDao()
	commentEntityArr, err := getCommentEntityArr(model, userDao)
	if commons.InError(err) {
		return nil, err
	}
	if len(commentEntityArr) == 0 {
		createdAtTime := ""
		if model.PostbackParams != nil {
			createdAtTime = model.PostbackParams.CreatedAt
		}
		return getNoCommentsResponse(createdAtTime), nil
	}

	parentEntityArr, parentEntryMap := replyThreadCommon.GetParentEntityArrAndMapFromCommentEntityArr(commentEntityArr)
	replyThreadCommon.FillReplyTreeInParentIdArr(model.ArticleId, MaxCommentReplyLevel, MaxRepliesCount, parentEntityArr, parentEntryMap, replyDao, userDao)

	postbackParams, hasMore := getPaginationData(model, commentEntityArr)
	return &ResponseModel{
		Status:         commons.NewResponseStatus().SUCCESS,
		Message:        "",
		Comments:       commentEntityArr,
		PostbackParams: postbackParams,
		HasMore:        hasMore,
	}, nil
}

func getCommentEntityArr(model *RequestModel, userDao *userDaos.UserDao) ([]*apiEntity.CommentEntity, error) {
	commentDao := daos.GetCommentDao()
	createdAtTime := commons.MaxTime
	var err error
	if model.PostbackParams != nil {
		createdAtTime, err = commons.ParseTime(model.PostbackParams.CreatedAt)
		if commons.InError(err) {
			return nil, errors.New("Error in parsing created at time")
		}
	}

	commentDbModels, err := commentDao.ReadComments(model.ArticleId, createdAtTime, ArticleCommentLimit)
	if commons.InError(err) {
		return nil, errors.New("Error is fetching comments")
	}

	commentEntityArr := make([]*apiEntity.CommentEntity, len(*commentDbModels))
	PostedByUserArr, err := commentCommons.GetUsersForComments(commentDbModels, userDao)
	if commons.InError(err) {
		return nil, errors.New("Error in fetching users for comments")
	}
	for index, commentDbModel := range *commentDbModels {
		commentEntityArr[index] = apiEntity.GetCommentEntityFromModel(&commentDbModel, &(*PostedByUserArr)[index])
	}
	return commentEntityArr, nil
}

func getPaginationData(model *RequestModel, commentEntityArr []*apiEntity.CommentEntity) (string, bool) {
	articleDao := daos.GetArticleDao()
	totalTopCommentCount := articleDao.GetArticleTopCommentCount(model.ArticleId)
	commentEntityArrLen := len(commentEntityArr)

	commentsServedTillNowCount := commentEntityArrLen
	if commentEntityArrLen == 0 {
		return "", false
	}
	if model.PostbackParams != nil {
		commentsServedTillNowCount += model.PostbackParams.Count
	}
	postbackParams := PostbackParams{
		Count:     commentsServedTillNowCount,
		CreatedAt: (commentEntityArr[commentEntityArrLen-1]).PostedAt,
	}
	postbackParamsStr, err := json.Marshal(postbackParams)
	if commons.InError(err) {
		return "", false
	} else {
		return string(postbackParamsStr), uint64(commentsServedTillNowCount) < totalTopCommentCount
	}
}

func getNoCommentsResponse(createdAt string) *ResponseModel {
	var message string
	if createdAt == "" {
		message = "No comments to show"
	} else {
		message = "No more comments to show"
	}
	entity := apiEntity.GetNoMoreCommentEntity(message)
	return &ResponseModel{
		Status:         commons.NewResponseStatus().SUCCESS,
		Message:        "",
		Comments:       []*apiEntity.CommentEntity{&entity},
		PostbackParams: "",
		HasMore:        false,
	}
}
