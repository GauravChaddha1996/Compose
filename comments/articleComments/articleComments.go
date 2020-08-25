package articleComments

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"compose/dbModels"
	"encoding/json"
	"strconv"
	"sync"
	"time"
)

const ArticleCommentLimit = 20
const CommentReplyLimit = 20
const CommentReplyMaxLevel = 10

func getArticleCommentsResponse(model *RequestModel) (*ResponseModel, error) {
	commentDao := daos.GetCommentDao()
	replyDao := daos.GetReplyDao()
	createdAt := model.PostbackParams["created_at"]
	commentsCountServedTillNow, err := strconv.ParseUint(model.PostbackParams["count"], 10, 64)
	if commons.InError(err) {
		commentsCountServedTillNow = 0
	}
	totalTopCommentCount := commentCommons.ArticleServiceContract.GetArticleTopCommentCount(model.ArticleId)

	comments, err := commentDao.ReadComments(model.ArticleId, getCreatedAtTimeFromPostbackParams(createdAt), ArticleCommentLimit)
	if commons.InError(err) {
		return nil, err
	}

	commentsLen := len(*comments)

	if commentsLen == 0 {
		return getNoCommentsResponse(createdAt), nil
	}

	PostedByUserArr, err := commentCommons.GetUsersForComments(comments)
	if commons.InError(err) {
		return nil, err
	}

	commentsResponseArr := make([]commentCommons.CommentEntity, commentsLen)
	wg := sync.WaitGroup{}
	for index, entry := range *comments {
		wg.Add(1)
		go func(i int, e dbModels.Comment) {
			repliesForEntry := replyDao.GetReplies(e.CommentId, CommentReplyMaxLevel, 1, CommentReplyLimit)
			var replies []commentCommons.ReplyEntity
			if repliesForEntry == nil {
				replies = nil
			} else {
				replies = *repliesForEntry
			}
			commentsResponseArr[i] = commentCommons.CommentEntity{
				CommentType:  commentCommons.NewCommentEntityTypeWrapper().CommentTypeNormal,
				CommentId:    e.CommentId,
				Markdown:     e.Markdown,
				PostedByUser: &(*PostedByUserArr)[i],
				Replies:      replies,
			}
			wg.Done()
		}(index, entry)
	}
	wg.Wait()

	return &ResponseModel{
		Comments:       commentsResponseArr,
		PostbackParams: getPostbackParamsForPagination(comments, commentsLen, commentsCountServedTillNow),
		HasMore:        (commentsCountServedTillNow + uint64(commentsLen)) < totalTopCommentCount,
	}, nil
}

func getCreatedAtTimeFromPostbackParams(createdAt string) time.Time {
	if len(createdAt) == 0 {
		createdAt = commons.MAX_TIME
	}
	createdAtTime, _ := commons.ParseTime(createdAt)
	return createdAtTime
}

func getNoCommentsResponse(createdAt string) *ResponseModel {
	var message string
	if createdAt == commons.MAX_TIME {
		message = "No comments to show"
	} else {
		message = "No more comments to show"
	}
	return &ResponseModel{
		Status:         commons.NewResponseStatus().SUCCESS,
		Message:        "",
		Comments:       []commentCommons.CommentEntity{commentCommons.GetNoMoreCommentEntity(message)},
		PostbackParams: "",
		HasMore:        false,
	}
}

func getPostbackParamsForPagination(comments *[]dbModels.Comment, commentsLen int, count uint64) string {
	postbackParamsMap := make(map[string]string)
	postbackParamsMap["created_at"] = ((*comments)[commentsLen-1]).CreatedAt.Format(commons.TimeFormat)
	postbackParamsMap["count"] = strconv.FormatUint(count+uint64(commentsLen), 10)
	postbackParamsStr, err := json.Marshal(postbackParamsMap)
	var postbackParams string
	if commons.InError(err) {
		postbackParams = ""
	} else {
		postbackParams = string(postbackParamsStr)
	}
	return postbackParams
}
