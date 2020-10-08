package daos

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"compose/dbModels"
	"errors"
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

func (dao ReplyDao) GetRepliesInParentIds(parentIds []string) ([]*dbModels.Reply, error) {
	parentIdsLen := len(parentIds)
	if parentIdsLen == 0 {
		return []*dbModels.Reply{}, nil
	}
	parentIdQuery := ""
	for index, parentId := range parentIds {
		parentIdQuery += "\"" + parentId + "\""
		if index != parentIdsLen-1 {
			parentIdQuery += ","
		}
	}
	whereQuery := "parent_id IN (" + parentIdQuery + ")"

	var dbReplies []*dbModels.Reply
	queryResult := dao.db.Where(whereQuery).Find(&dbReplies)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return dbReplies, nil
}

func (dao ReplyDao) FindReply(replyId string) (*dbModels.Reply, error) {
	var reply dbModels.Reply
	queryResult := dao.db.Where("reply_id = ?", replyId).Limit(1).Find(&reply)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return &reply, nil
}

func (dao ReplyDao) IncreaseReplyCount(replyId string) error {
	reply, err := dao.FindReply(replyId)
	if commons.InError(err) {
		return errors.New("Cannot find reply for this reply id")
	}
	var changeMap = make(map[string]interface{})
	changeMap["reply_count"] = (*reply).ReplyCount + 1
	return dao.UpdateReply(replyId, changeMap)
}

func (dao ReplyDao) UpdateReply(replyId string, changeMap map[string]interface{}) error {
	var reply dbModels.Reply
	return dao.db.Model(reply).Where("reply_id = ?", replyId).UpdateColumns(changeMap).Error
}