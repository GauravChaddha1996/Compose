package commentAndReply

import (
	"compose/commons"
	"compose/dataLayer/dbModels"
	"errors"
	"gorm.io/gorm"
)

type ReplyDao struct {
	DB *gorm.DB
}

func (dao ReplyDao) CreateReply(reply dbModels.Reply) error {
	return dao.DB.Create(&reply).Error
}

func (dao ReplyDao) DoesParentExist(parentId string) bool {
	var reply dbModels.Reply
	queryResult := dao.DB.Select("reply_id").Where("reply_id = ?", parentId).Limit(1).Find(&reply)
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
	queryResult := dao.DB.Where(whereQuery).Find(&dbReplies)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return dbReplies, nil
}

func (dao ReplyDao) FindReply(replyId string) (*dbModels.Reply, error) {
	var reply dbModels.Reply
	queryResult := dao.DB.Where("reply_id = ?", replyId).Limit(1).Find(&reply)
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
	return dao.DB.Model(reply).Where("reply_id = ?", replyId).UpdateColumns(changeMap).Error
}

func (dao ReplyDao) DeleteReply(replyId string) error {
	var reply dbModels.Reply
	return dao.DB.Where("reply_id = ?", replyId).Find(&reply).Unscoped().Delete(reply).Error
}

func (dao ReplyDao) DeleteRepliesForArticle(articleId string) error {
	var replies []dbModels.Reply
	return dao.DB.Where("article_id = ?", articleId).Find(&replies).Unscoped().Delete(replies).Error
}
