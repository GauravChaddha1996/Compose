package articleLikes

import (
	"compose/commons"
	"compose/like/daos"
	"compose/like/likeCommons"
	"errors"
)

func getArticleLikeUserList(model *RequestModel) (*[]LikedByUser, *uint64, error) {
	likeDao := daos.GetLikeDao()
	likeEntries, err := likeDao.GetArticleLikes(model.ArticleId, model.LastLikeId, 0) // Send 0 for using default limit
	if commons.InError(err) {
		return nil, nil, errors.New("Error fetching liked userId list")
	}
	likeEntriesSize := len(*likeEntries)
	if likeEntriesSize == 0 {
		return nil, nil, nil
	}
	lastLikeEntry := (*likeEntries)[likeEntriesSize-1]

	userIdList := make([]string, likeEntriesSize)
	for index, entry := range *likeEntries {
		userIdList[index] = entry.UserId
	}
	users, err := likeCommons.UserServiceContract.GetUsers(userIdList)
	if commons.InError(err) {
		return nil, nil, errors.New("Cannot fetch details via userId")
	}

	likedByUserArr := make([]LikedByUser, likeEntriesSize)
	for index := range users {
		user := users[index]
		likedByUserArr[index] = LikedByUser{
			UserId:   user.UserId,
			PhotoUrl: user.PhotoUrl,
			Name:     user.Name,
		}
	}
	return &likedByUserArr, &lastLikeEntry.Id, nil
}
