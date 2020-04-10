package articleLikes

import (
	"compose/commons"
	"compose/like/daos"
	"compose/like/likeCommons"
	"errors"
)

func getArticleLikesResponse(model *RequestModel) (*ResponseModel, error) {
	likeDao := daos.GetLikeDao()

	var articleLikeEntryLimit = 3
	likeEntries, err := likeDao.GetArticleLikes(model.ArticleId, model.MaxLikedAt, articleLikeEntryLimit)
	if commons.InError(err) {
		return nil, errors.New("Error fetching liked userId list")
	}
	likeEntriesSize := len(*likeEntries)
	if likeEntriesSize == 0 {
		var message string
		if *model.MaxLikedAt == model.DefaultMaxLikedAt {
			message = "No likes to show"
		} else {
			message = "No more likes to show"
		}
		return &ResponseModel{
			Status:       commons.NewResponseStatus().SUCCESS,
			Message:      message,
			HasMoreLikes: false,
		}, nil
	}

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
	lastLikedAt := (*likeEntries)[likeEntriesSize-1].CreatedAt.Format("2 Jan 2006 15:04:05")
	return &ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		LikedByUsers: likedByUserArr,
		MaxLikedAt:   lastLikedAt,
		HasMoreLikes: !(likeEntriesSize < articleLikeEntryLimit),
	}, nil
}
