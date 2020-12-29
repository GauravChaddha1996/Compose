package articleLikes

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/rs/zerolog"
)

func getArticleLikesResponse(model *RequestModel, sublogger *zerolog.Logger) (*ResponseModel, error) {
	var articleLikeEntryLimit = 3
	likeDao := daos.GetLikeDao()

	likeEntries, err := likeDao.GetArticleLikes(model.ArticleId, model.MaxLikedAt, articleLikeEntryLimit)
	if commons.InError(err) {
		return nil, errors.New("Error fetching liked userId list")
	}
	sublogger.Info().Msg("Like entries are found")

	likeEntriesSize := len(*likeEntries)
	if likeEntriesSize == 0 {
		sublogger.Info().Msg("Like entries are empty")
		return getEmptyLikeEntriesResponse(model), nil
	}

	likedByUserArr, err := getLikedByUserArr(likeEntries, likeEntriesSize)
	if commons.InError(err) {
		return nil, errors.New("Cannot fetch details via userIds")
	}
	sublogger.Info().Msg("User entries are fetched")

	lastLikedAt := (*likeEntries)[likeEntriesSize-1].CreatedAt.Format(commons.TimeFormat)
	return &ResponseModel{
		Status:       commons.NewResponseStatus().SUCCESS,
		LikedByUsers: likedByUserArr,
		MaxLikedAt:   lastLikedAt,
		HasMoreLikes: !(likeEntriesSize < articleLikeEntryLimit),
	}, nil
}

func getEmptyLikeEntriesResponse(model *RequestModel) *ResponseModel {
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
	}
}

func getLikedByUserArr(likeEntries *[]dbModels.LikeEntry, likeEntriesSize int) ([]*apiEntity.SmallUserEntity, error) {
	userDao := daos.GetUserDao()
	userIdList := make([]string, likeEntriesSize)
	for index, entry := range *likeEntries {
		userIdList[index] = entry.UserId
	}
	users, err := userDao.FindUserViaIds(userIdList)
	if commons.InError(err) {
		return nil, err
	}
	likedByUserArr := make([]*apiEntity.SmallUserEntity, likeEntriesSize)
	for index := range users {
		user := users[index]
		likedByUserArr[index] = apiEntity.GetSmallUserEntity(user)
	}
	return likedByUserArr, nil
}
