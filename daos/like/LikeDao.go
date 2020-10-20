package like

import (
	"compose/commons"
	"compose/dbModels"
	"errors"
	"gorm.io/gorm"
	"time"
)

type LikeDao struct {
	DB *gorm.DB
}

func (dao LikeDao) LikeArticle(entry *dbModels.LikeEntry) error {
	return dao.DB.Create(entry).Error
}

func (dao LikeDao) FindLikeEntry(articleId string, userId string) (*dbModels.LikeEntry, error) {
	var entry dbModels.LikeEntry
	queryResult := dao.DB.Where("user_id = ? && article_id = ?", userId, articleId).Find(&entry)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return &entry, nil
}

func (dao LikeDao) GetArticleLikes(articleId string, maxLikedAt *time.Time, limit int) (*[]dbModels.LikeEntry, error) {
	var likeEntries []dbModels.LikeEntry
	queryResult := dao.DB.
		Where("article_id = ? && created_at < ?", articleId, maxLikedAt).
		Order("created_at desc").
		Limit(limit).
		Find(&likeEntries)
	if commons.InError(queryResult.Error) {
		return nil, errors.New("Error in fetching like entries")
	}
	return &likeEntries, nil
}

func (dao LikeDao) GetUserLikes(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.LikeEntry, error) {
	var likeEntries []dbModels.LikeEntry
	queryResult := dao.DB.
		Where("user_id = ? && created_at < ?", userId, maxCreatedAtTime).
		Order("created_at desc").
		Limit(limit).
		Find(&likeEntries)
	if commons.InError(queryResult.Error) {
		return nil, errors.New("Error in fetching like entries")
	}
	return &likeEntries, nil
}

func (dao LikeDao) UnlikeArticle(likeEntry *dbModels.LikeEntry) error {
	var entry dbModels.LikeEntry
	return dao.DB.Where("id = ?", likeEntry.Id).Unscoped().Delete(&entry).Error
}

func (dao LikeDao) DeleteAllLikesOfArticle(articleId string) error {
	var entries []dbModels.LikeEntry
	return dao.DB.Where("article_id = ?", articleId).Unscoped().Delete(&entries).Error
}
