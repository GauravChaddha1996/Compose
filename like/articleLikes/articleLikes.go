package articleLikes

import (
	"compose/commons"
	"compose/like/daos"
	"compose/like/likeCommons"
	"errors"
	"strconv"
)

func getArticleLikesResponse(model *RequestModel) (*ResponseModel, error) {
	likeDao := daos.GetLikeDao()

	var articleLikeEntryLimit = 3
	likeEntries, err := likeDao.GetArticleLikes(model.ArticleId, model.LastLikeId, articleLikeEntryLimit)
	if commons.InError(err) {
		return nil, errors.New("Error fetching liked userId list")
	}
	likeEntriesSize := len(*likeEntries)
	if likeEntriesSize == 0 {
		var message string
		if *model.LastLikeId == model.DefaultLastLikeId {
			message = "No likes to show"
		} else {
			message = "No more likes to show"
		}
		return &ResponseModel{
			Status:       commons.ResponseStatusWrapper{}.SUCCESS,
			Message:      message,
			HasMoreLikes: false,
		}, nil
	}
	lastLikeEntry := (*likeEntries)[likeEntriesSize-1]

	userIdList := make([]string, likeEntriesSize)
	for index, entry := range *likeEntries {
		userIdList[index] = entry.UserId
	}
	users, err := likeCommons.UserServiceContract.GetUsers(userIdList)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch details via userId")
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
	lastLikeId := strconv.FormatUint(lastLikeEntry.Id, 10)
	return &ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		LikedByUsers: likedByUserArr,
		LastLikeId:   lastLikeId,
		HasMoreLikes: !(likeEntriesSize < articleLikeEntryLimit),
	}, nil
}
