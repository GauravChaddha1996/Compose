package daos

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"compose/dbModels"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
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

func (dao ReplyDao) GetReplies(parentId string, maxLevel int, currentLevel int, limit int, maxCreatedAtTime *time.Time, previousReplyCount int) *[]commentCommons.ReplyEntity {
	if currentLevel > maxLevel {
		return nil
	}
	replies := make([]dbModels.Reply, limit)

	queryResult := dao.db.Where("parent_id = ? && created_at < ?", parentId, maxCreatedAtTime).Order("created_at desc").Limit(limit).Find(&replies)
	if commons.InError(queryResult.Error) {
		return commentCommons.GetErrorReplies()
	}
	if queryResult.RowsAffected == 0 {
		return nil
	}

	replyResponseArr := make([]commentCommons.ReplyEntity, len(replies))
	userArr, err := commentCommons.GetUsersForReplies(&replies)
	if commons.InError(err) {
		return commentCommons.GetErrorReplies()
	}
	for index, reply := range replies {
		maxTime, _ := commons.MaxTime()
		childReplies := dao.GetReplies(reply.ReplyId, maxLevel, currentLevel+1, limit, &maxTime, 0)
		var childRepliesResponse []commentCommons.ReplyEntity
		if childReplies != nil {
			childRepliesResponse = *childReplies
			childRepliesLen := len(childRepliesResponse)
			newReplyCount := childRepliesLen + previousReplyCount

			// this means that we have more reply in this reply entry
			if uint64(newReplyCount) < reply.ReplyCount {
				lastChildReply := childRepliesResponse[childRepliesLen-1]
				continueThreadPostbackParams := commentCommons.GetContinueThreadPostbackParams(reply.ArticleId, reply.ReplyId, lastChildReply.PostedAt, newReplyCount)
				childRepliesResponse = append(childRepliesResponse, commentCommons.GetContinueReplyEntity(continueThreadPostbackParams))
			}
		} else {
			childRepliesResponse = nil
		}
		replyResponseArr[index] = commentCommons.ReplyEntity{
			ReplyType:    commentCommons.NewReplyEntityTypeWrapper().ReplyTypeNormal,
			ReplyId:      reply.ReplyId,
			Markdown:     reply.Markdown,
			PostedByUser: &(*userArr)[index],
			PostedAt:     reply.CreatedAt.Format(commons.TimeFormat),
			Replies:      childRepliesResponse,
		}
	}
	return &replyResponseArr
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