package daos

import (
	"compose/commons"
	"compose/like/likeCommons"
	"compose/dbModels"
	"errors"
	"github.com/jinzhu/gorm"
)

type LikeDao struct {
	db *gorm.DB
}

func GetLikeDao() LikeDao {
	return LikeDao{db: likeCommons.Database}
}

func GetLikeDaoDuringTransaction(db *gorm.DB) LikeDao {
	return LikeDao{db: db}
}

func (dao LikeDao) LikeArticle(entry *dbModels.LikeEntry) error {
	return dao.db.Create(entry).Error
}

func (dao LikeDao) FindLikeEntry(articleId string, userId string) (*dbModels.LikeEntry, error) {
	var entry dbModels.LikeEntry
	queryResult := dao.db.Where("user_id = ? && article_id = ?", userId, articleId).Find(&entry)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return &entry, nil
}

func (dao LikeDao) GetArticleLikes(articleId string, lastLikeId *string, limit int) (*[]dbModels.LikeEntry, error) {
	var likeEntries []dbModels.LikeEntry
	queryResult := dao.db.Where("article_id = ? && id > ?", articleId, lastLikeId).Limit(limit).Find(&likeEntries)
	if commons.InError(queryResult.Error) {
		return nil, errors.New("Error in fetching like entries")
	}
	return &likeEntries, nil
}

func (dao LikeDao) UnlikeArticle(likeEntry *dbModels.LikeEntry) error {
	var entry dbModels.LikeEntry
	return dao.db.Where("id = ?", likeEntry.Id).Unscoped().Delete(&entry).Error
}
