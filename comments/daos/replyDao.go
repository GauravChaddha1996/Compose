package daos

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"compose/dbModels"
	"github.com/jinzhu/gorm"
)

type ReplyDao struct {
	db *gorm.DB
}

func GetReplyDao() *ReplyDao {
	return &ReplyDao{commentCommons.Database}
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

func (dao ReplyDao) GetReplies(parentId string, maxLevel int, currentLevel int, limit int) *[]commentCommons.ReplyEntity {
	if currentLevel > maxLevel {
		return nil
	}
	replies := make([]dbModels.Reply, limit)

	queryResult := dao.db.Where("parent_id = ?", parentId).Order("created_at desc").Limit(limit).Find(&replies)
	if commons.InError(queryResult.Error) {
		return nil
	}
	if queryResult.RowsAffected == 0 {
		return nil
	}

	replyResponseArr := make([]commentCommons.ReplyEntity, len(replies))
	userArr, err := commentCommons.GetUsersForReplies(&replies)
	if commons.InError(err) {
		return nil
	}
	for index, reply := range replies {
		childReplies := dao.GetReplies(reply.ReplyId, maxLevel, currentLevel+1, limit)
		var childRepliesResponse []commentCommons.ReplyEntity
		if childReplies != nil {
			childRepliesResponse = *childReplies
		} else {
			childRepliesResponse = nil
		}
		replyResponseArr[index] = commentCommons.ReplyEntity{
			ReplyId:       reply.ReplyId,
			Markdown:      reply.Markdown,
			PostedByUser: (*userArr)[index],
			Replies:       childRepliesResponse,
		}
	}
	return &replyResponseArr
}
