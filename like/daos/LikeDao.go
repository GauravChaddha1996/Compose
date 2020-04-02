package daos

import (
	"compose/commons"
	"compose/like/likeCommons"
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

func (dao LikeDao) LikeArticle(entry *likeCommons.LikeEntry) error {
	return dao.db.Create(entry).Error
}

func (dao LikeDao) FindLikeEntry(articleId string, userId string) (*likeCommons.LikeEntry, error) {
	var entry likeCommons.LikeEntry
	queryResult := dao.db.Where("user_id = ? && article_id = ?", userId, articleId).Find(&entry)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return &entry, nil
}

func (dao LikeDao) UnlikeArticle(likeEntry *likeCommons.LikeEntry) error {
	var entry likeCommons.LikeEntry
	return dao.db.Where("id = ?", likeEntry.Id).Unscoped().Delete(&entry).Error
}
