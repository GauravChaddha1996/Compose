package commentCommons

import (
	"compose/commons"
	"compose/dbModels"
	"compose/serviceContracts"
	"errors"
	"gorm.io/gorm"
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

func GetUsersForReplies(replies []*dbModels.Reply) (*[]PostedByUser, error) {
	commentsLen := len(replies)
	userIdList := make([]string, commentsLen)
	for index, entry := range replies {
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
	userMap := make(map[string]*dbModels.User)
	for _, user := range users {
		userMap[user.UserId] = user
	}

	PostedByUserArr := make([]PostedByUser, userLen)
	for index, userId := range *userIdList {
		user := userMap[userId]
		PostedByUserArr[index] = PostedByUser{
			UserId:   user.UserId,
			PhotoUrl: user.PhotoUrl,
			Name:     user.Name,
		}
	}
	return &PostedByUserArr, nil
}