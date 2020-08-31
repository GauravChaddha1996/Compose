package commentCommons

import (
	"compose/commons"
	"compose/dbModels"
	"compose/serviceContracts"
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
)

var Database *gorm.DB
var ArticleServiceContract serviceContracts.ArticleServiceContract
var UserServiceContract serviceContracts.UserServiceContract

func GetUsersForComments(comments *[]dbModels.Comment) (*[]PostedByUser, error) {
	commentsLen := len(*comments)
	userIdList := make([]string, commentsLen)
	for index, entry := range *comments {
		userIdList[index] = entry.UserId
	}
	return getUserArr(&userIdList)
}

func GetUsersForReplies(replies *[]dbModels.Reply) (*[]PostedByUser, error) {
	commentsLen := len(*replies)
	userIdList := make([]string, commentsLen)
	for index, entry := range *replies {
		userIdList[index] = entry.UserId
	}
	return getUserArr(&userIdList)
}

func getUserArr(userIdList *[]string) (*[]PostedByUser, error) {
	userLen := len(*userIdList)
	users, err := UserServiceContract.GetUsers(*userIdList)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch details via userId")
	}

	PostedByUserArr := make([]PostedByUser, userLen)
	for index := range users {
		user := users[index]
		PostedByUserArr[index] = PostedByUser{
			UserId:   user.UserId,
			PhotoUrl: user.PhotoUrl,
			Name:     user.Name,
		}
	}
	return &PostedByUserArr, nil
}

func GetContinueThreadPostbackParams(articleId string, parentId string, createdAt string, replyCount int) string {
	var postbackParams string
	postbackParamsMap := make(map[string]string)
	postbackParamsMap["parent_id"] = parentId
	postbackParamsMap["article_id"] = articleId
	postbackParamsMap["created_at"] = createdAt
	postbackParamsMap["reply_count"] = strconv.Itoa(replyCount)
	postbackParamsStr, err := json.Marshal(postbackParamsMap)
	if commons.InError(err) {
		postbackParams = ""
	} else {
		postbackParams = string(postbackParamsStr)
	}
	return postbackParams
}