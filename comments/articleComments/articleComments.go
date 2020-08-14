package articleComments

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"encoding/json"
	"errors"
	"time"
)

func getArticleCommentsResponse(model *RequestModel) (*ResponseModel, error) {
	tx := commentCommons.Database.Begin()
	commentDao := daos.GetCommentDaoDuringTransaction(tx)

	commentsLimit := 8
	defaultCreatedAt := commons.MAX_TIME
	createdAt := model.PostbackParams["created_at"]
	if len(createdAt) == 0 {
		createdAt = defaultCreatedAt
	}

	createdAtTime, _ := time.ParseInLocation(commons.TimeFormat, createdAt, time.Now().Location())

	comments, err := commentDao.ReadComments(model.ArticleId, createdAtTime, commentsLimit)
	if commons.InError(err) {
		return nil, err
	}

	commentsLen := len(*comments)

	if commentsLen == 0 {
		var message string
		if createdAt == defaultCreatedAt {
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
		}, nil
	}

	userIdList := make([]string, commentsLen)
	for index, entry := range *comments {
		userIdList[index] = entry.UserId
	}
	users, err := commentCommons.UserServiceContract.GetUsers(userIdList)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch details via userId")
	}

	commentByUserArr := make([]CommentByUser, commentsLen)
	for index := range users {
		user := users[index]
		commentByUserArr[index] = CommentByUser{
			UserId:   user.UserId,
			PhotoUrl: user.PhotoUrl,
			Name:     user.Name,
		}
	}

	commentsResponseArr := make([]CommentResponse, commentsLen)
	for index, entry := range *comments {
		commentsResponseArr[index] = CommentResponse{
			CommentId:     entry.CommentId,
			Markdown:      entry.Markdown,
			CommentByUser: commentByUserArr[index],
		}
	}

	postbackParamsMap := make(map[string]string)
	postbackParamsMap["created_at"] = ((*comments)[commentsLen-1]).CreatedAt.Format(commons.TimeFormat)
	postbackParamsStr, err := json.Marshal(postbackParamsMap)
	var postbackParams string
	if commons.InError(err) {
		postbackParams = ""
	} else {
		postbackParams = string(postbackParamsStr)
	}
	return &ResponseModel{
		Comments:       commentsResponseArr,
		PostbackParams: postbackParams,
		HasMore:        !(commentsLen > commentsLimit),
	}, nil
}
