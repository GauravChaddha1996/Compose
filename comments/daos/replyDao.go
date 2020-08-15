package daos

import (
	"compose/commons"
	"compose/dbModels"
	"github.com/jinzhu/gorm"
)

type ReplyDao struct {
	db *gorm.DB
}

func GetReplyDaoDuringTransaction(db *gorm.DB) *ReplyDao {
	return &ReplyDao{db}
}

func (dao ReplyDao) CreateReply(reply dbModels.Reply) error {
	return dao.db.Create(&reply).Error
}

func (dao ReplyDao) DoesParentExist(parentId string) bool {
	var reply dbModels.Reply
	queryResult := dao.db.Select("reply_id").Where("reply_id = ?", parentId).Limit(1).Find(&reply)
	if commons.InError(queryResult.Error) {
		return false
	}
	return true
}
