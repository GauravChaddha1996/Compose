package commentCommons

import (
	"compose/commons"
	"compose/dataLayer/daos/user"
	"compose/dataLayer/models"
	"errors"
)

func GetUsersForComments(comments *[]models.Comment, userDao *user.UserDao) (*[]PostedByUser, error) {
	commentsLen := len(*comments)
	userIdList := make([]string, commentsLen)
	for index, entry := range *comments {
		userIdList[index] = entry.UserId
	}
	return getUserArr(&userIdList, userDao)
}

func GetUsersForReplies(replies []*models.Reply, userDao *user.UserDao) (*[]PostedByUser, error) {
	commentsLen := len(replies)
	userIdList := make([]string, commentsLen)
	for index, entry := range replies {
		userIdList[index] = entry.UserId
	}
	return getUserArr(&userIdList, userDao)
}

func getUserArr(userIdList *[]string, userDao *user.UserDao) (*[]PostedByUser, error) {
	userLen := len(*userIdList)
	users, err := userDao.FindUserViaIds(*userIdList)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch details via userId")
	}
	userMap := make(map[string]*models.User)
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
