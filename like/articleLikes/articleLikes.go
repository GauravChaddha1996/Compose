package articleLikes

import (
	"compose/commons"
	"compose/daos"
	"errors"
)

func getArticleLikesResponse(model *RequestModel) (*ResponseModel, error) {
	likeDao := daos.GetLikeDao()
	userDao := daos.GetUserDao()

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
	users, err := userDao.FindUserViaIds(userIdList)
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
	lastLikedAt := (*likeEntries)[likeEntriesSize-1].CreatedAt.Format(commons.TimeFormat)
	return &ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		LikedByUsers: likedByUserArr,
		MaxLikedAt:   lastLikedAt,
		HasMoreLikes: !(likeEntriesSize < articleLikeEntryLimit),
	}, nil
}
