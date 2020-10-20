package replyThreadCommon

import (
	"compose/commons"
	commentAndReplyDaos "compose/dataLayer/daos/commentAndReply"
	userDaos "compose/dataLayer/daos/user"
	"compose/dataLayer/models"
	commentCommons2 "compose/endpoints/comments/commentCommons"
	"encoding/json"
	"errors"
	"strconv"
)

type ReplyThreadParentModel struct {
	Id            string
	IsComment     bool
	IsReply       bool
	CommentEntity *commentCommons2.CommentEntity
	ReplyEntity   *commentCommons2.ReplyEntity
}

func GetParentEntityArrAndMapFromCommentEntityArr(parentEntityArr []*commentCommons2.CommentEntity) ([]*ReplyThreadParentModel, *map[string]int) {
	newParentEntityArrLength := len(parentEntityArr)
	newParentEntityArr := make([]*ReplyThreadParentModel, newParentEntityArrLength)
	newParentEntryMap := make(map[string]int)
	for index, parentEntity := range parentEntityArr {
		newParentEntityArr[index] = &ReplyThreadParentModel{
			Id:            parentEntity.CommentId,
			IsComment:     true,
			IsReply:       false,
			CommentEntity: parentEntity,
		}
		newParentEntryMap[parentEntity.CommentId] = index
	}
	return newParentEntityArr, &newParentEntryMap
}

func GetParentEntityArrAndMapFromReplyEntityArr(parentEntityArr []*commentCommons2.ReplyEntity) ([]*ReplyThreadParentModel, *map[string]int) {
	newParentEntityArrLength := len(parentEntityArr)
	newParentEntityArr := make([]*ReplyThreadParentModel, newParentEntityArrLength)
	newParentEntryMap := make(map[string]int)
	for index, parentEntity := range parentEntityArr {
		newParentEntityArr[index] = &ReplyThreadParentModel{
			Id:          parentEntity.ReplyId,
			IsComment:   false,
			IsReply:     true,
			ReplyEntity: parentEntity,
		}
		newParentEntryMap[parentEntity.ReplyId] = index
	}
	return newParentEntityArr, &newParentEntryMap
}

func getParentEntityArrAndMapFromReplyEntityArr(parentEntityArr []*commentCommons2.ReplyEntity) ([]*ReplyThreadParentModel, *map[string]int) {
	newParentEntityArrLength := len(parentEntityArr)
	newParentEntityArr := make([]*ReplyThreadParentModel, newParentEntityArrLength)
	newParentEntryMap := make(map[string]int)
	for index, parentEntity := range parentEntityArr {
		newParentEntityArr[index] = &ReplyThreadParentModel{
			Id:          parentEntity.ReplyId,
			IsComment:   false,
			IsReply:     true,
			ReplyEntity: parentEntity,
		}
		newParentEntryMap[parentEntity.ReplyId] = index
	}
	return newParentEntityArr, &newParentEntryMap
}

func FillReplyTreeInParentIdArr(
	articleId string,
	maxCommentReplyLevel int,
	maxRepliesCount int,
	parentEntityArr []*ReplyThreadParentModel,
	parentEntryMap *map[string]int,
	replyDao *commentAndReplyDaos.ReplyDao,
	userDao *userDaos.UserDao,
) {
	currentReplyLevel := 0
	repliesCount := 0
	repliesFinishReached := false
	breakDueToError := false
	for currentReplyLevel < maxCommentReplyLevel && repliesCount < maxRepliesCount && repliesFinishReached == false {
		replyDbModels, replyEntityArr, err := GetReplyEntityArr(parentEntityArr, replyDao, userDao)
		if len(replyDbModels) == 0 {
			repliesFinishReached = true
		}

		if commons.InError(err) {
			breakDueToError = true
			break
		}
		for index, replyEntity := range replyEntityArr {
			replyDbModel := replyDbModels[index]
			index = (*parentEntryMap)[replyDbModel.ParentId]
			parentEntity := parentEntityArr[index]
			if parentEntity.IsComment {
				parentComment := parentEntity.CommentEntity
				parentComment.Replies = append(parentComment.Replies, replyEntity)
			}
			if parentEntity.IsReply {
				parentReply := parentEntity.ReplyEntity
				parentReply.Replies = append(parentReply.Replies, replyEntity)
			}
		}

		parentEntityArr, parentEntryMap = getParentEntityArrAndMapFromReplyEntityArr(replyEntityArr)
		repliesCount += len(replyEntityArr)
		currentReplyLevel += 1
	}
	CheckForContinueThread(repliesFinishReached, breakDueToError, articleId, parentEntityArr)
}

func GetReplyEntityArr(parentEntityArr []*ReplyThreadParentModel, replyDao *commentAndReplyDaos.ReplyDao, userDao *userDaos.UserDao) ([]*models.Reply, []*commentCommons2.ReplyEntity, error) {
	parentEntityArrLen := len(parentEntityArr)
	parentIds := make([]string, parentEntityArrLen)
	for index, parentEntity := range parentEntityArr {
		parentIds[index] = parentEntity.Id
	}
	replyDbModels, err := replyDao.GetRepliesInParentIds(parentIds)
	if commons.InError(err) {
		return nil, nil, errors.New("Error in fetching replies for parent entity arr")
	}

	PostedByUserArr, err := commentCommons2.GetUsersForReplies(replyDbModels, userDao)
	if commons.InError(err) {
		return nil, nil, errors.New("Error in fetching users for comments")
	}

	replyDbModelsLen := len(replyDbModels)
	replyEntityArr := make([]*commentCommons2.ReplyEntity, replyDbModelsLen)
	for index, replyDbModel := range replyDbModels {
		replyEntityArr[index] = commentCommons2.GetReplyEntityFromModel(replyDbModel, &(*PostedByUserArr)[index])
	}
	return replyDbModels, replyEntityArr, nil
}

func CheckForContinueThread(repliesFinishReached bool, breakDueToError bool, articleId string, parentEntityArr []*ReplyThreadParentModel) {
	if repliesFinishReached == false || breakDueToError {
		for _, parentEntity := range parentEntityArr {
			if parentEntity.IsComment {
				parentComment := parentEntity.CommentEntity
				if parentComment.ReplyCount > 0 {
					repliesLen := len(parentComment.Replies)
					var createdAtTime string
					if repliesLen == 0 {
						createdAtTime = commons.MAX_TIME
					} else {
						createdAtTime = parentComment.Replies[repliesLen-1].PostedAt
					}
					continuePostbackParams := GetContinueThreadPostbackParams(articleId, parentComment.CommentId, createdAtTime, repliesLen)
					continueReplyEntity := commentCommons2.GetContinueReplyEntity(continuePostbackParams)
					parentComment.Replies = append(parentComment.Replies, continueReplyEntity)
				}
			}
			if parentEntity.IsReply {
				parentReply := parentEntity.ReplyEntity
				if parentReply.ReplyCount > 0 {
					repliesLen := len(parentReply.Replies)
					var createdAtTime string
					if repliesLen == 0 {
						createdAtTime = commons.MAX_TIME
					} else {
						createdAtTime = parentReply.Replies[repliesLen-1].PostedAt
					}
					continuePostbackParams := GetContinueThreadPostbackParams(articleId, parentReply.ReplyId, createdAtTime, repliesLen)
					continueReplyEntity := commentCommons2.GetContinueReplyEntity(continuePostbackParams)
					parentReply.Replies = append(parentReply.Replies, continueReplyEntity)
				}
			}
		}
	}
}

func GetContinueThreadPostbackParams(articleId string, parentId string, createdAt string, replyCount int) string {
	var postbackParams string
	postbackParamsMap := make(map[string]string)
	postbackParamsMap["parent_id"] = parentId
	postbackParamsMap["article_id"] = articleId
	postbackParamsMap["created_at"] = createdAt
	postbackParamsMap["reply_count"] = strconv.Itoa(replyCount)
	postbackParamsStr, err := json.Marshal(postbackParamsMap)
	if commons.InError(err) {
		postbackParams = ""
	} else {
		postbackParams = string(postbackParamsStr)
	}
	return postbackParams
}
