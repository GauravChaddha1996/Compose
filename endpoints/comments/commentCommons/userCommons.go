package commentCommons

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	userDaos "compose/dataLayer/daos/user"
	"compose/dataLayer/dbModels"
	"errors"
)

func GetUsersForComments(comments *[]dbModels.Comment, userDao *userDaos .UserDao) (*[]apiEntity.PostedByUser, error) {
	commentsLen := len(*comments)
	userIdList := make([]string, commentsLen)
	for index, entry := range *comments {
		userIdList[index] = entry.UserId
	}
	return getUserArr(&userIdList, userDao)
}

func GetUsersForReplies(replies []*dbModels.Reply, userDao *userDaos.UserDao) (*[]apiEntity.PostedByUser, error) {
	commentsLen := len(replies)
	userIdList := make([]string, commentsLen)
	for index, entry := range replies {
		userIdList[index] = entry.UserId
	}
	return getUserArr(&userIdList, userDao)
}

func getUserArr(userIdList *[]string, userDao *userDaos .UserDao) (*[]apiEntity.PostedByUser, error) {
	userLen := len(*userIdList)
	users, err := userDao.FindUserViaIds(*userIdList)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch details via userId")
	}
	userMap := make(map[string]*dbModels.User)
	for _, user := range users {
		userMap[user.UserId] = user
	}

	PostedByUserArr := make([]apiEntity.PostedByUser, userLen)
	for index, userId := range *userIdList {
		user := userMap[userId]
		PostedByUserArr[index] = apiEntity.PostedByUser{
			UserId:   user.UserId,
			PhotoUrl: user.PhotoUrl,
			Name:     user.Name,
		}
	}
	return &PostedByUserArr, nil
}
