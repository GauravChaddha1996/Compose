package articleComments

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"compose/dbModels"
	"encoding/json"
	"sync"
	"time"
)

const ArticleCommentLimit = 20

func getArticleCommentsResponse(model *RequestModel) (*ResponseModel, error) {
	commentDao := daos.GetCommentDao()
	replyDao := daos.GetReplyDao()
	createdAt := model.PostbackParams["created_at"]

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
			repliesForEntry := replyDao.GetReplies(e.CommentId, 10, 1, 20)
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
				PostedByUser: (*PostedByUserArr)[i],
				Replies:      replies,
			}
			wg.Done()
		}(index, entry)
	}
	wg.Wait()

	return &ResponseModel{
		Comments:       commentsResponseArr,
		PostbackParams: getPostbackParamsForPagination(comments, commentsLen),
		HasMore:        !(commentsLen > ArticleCommentLimit),
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
		Message:        message,
		Comments:       nil,
		PostbackParams: "",
		HasMore:        false,
	}
}

func getPostbackParamsForPagination(comments *[]dbModels.Comment, commentsLen int) string {
	postbackParamsMap := make(map[string]string)
	postbackParamsMap["created_at"] = ((*comments)[commentsLen-1]).CreatedAt.Format(commons.TimeFormat)
	postbackParamsStr, err := json.Marshal(postbackParamsMap)
	var postbackParams string
	if commons.InError(err) {
		postbackParams = ""
	} else {
		postbackParams = string(postbackParamsStr)
	}
	return postbackParams
}
