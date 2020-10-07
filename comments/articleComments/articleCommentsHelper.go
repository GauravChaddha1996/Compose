package articleComments

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"encoding/json"
	"strconv"
)

func getNoCommentsResponse(createdAt string) *ResponseModel {
	var message string
	if createdAt == "" {
		message = "No comments to show"
	} else {
		message = "No more comments to show"
	}
	return &ResponseModel{
		Status:         commons.NewResponseStatus().SUCCESS,
		Message:        "",
		Comments:       []commentCommons.CommentEntity{commentCommons.GetNoMoreCommentEntity(message)},
		PostbackParams: "",
		HasMore:        false,
	}
}

func getContinueThreadPostbackParams(articleId string, parentId string, createdAt string, replyCount int) string {
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

func getParentEntityArrAndMapFromCommentEntityArr(parentEntityArr []*commentCommons.CommentEntity) ([]*ParentEntity, *map[string]int) {
	newParentEntityArrLength := len(parentEntityArr)
	newParentEntityArr := make([]*ParentEntity, newParentEntityArrLength)
	newParentEntryMap := make(map[string]int)
	for index, parentEntity := range parentEntityArr {
		newParentEntityArr[index] = &ParentEntity{
			Id:            parentEntity.CommentId,
			IsComment:     true,
			IsReply:       false,
			commentEntity: parentEntity,
		}
		newParentEntryMap[parentEntity.CommentId] = index
	}
	return newParentEntityArr, &newParentEntryMap
}

func getParentEntityArrAndMapFromReplyEntityArr(parentEntityArr []*commentCommons.ReplyEntity) ([]*ParentEntity, *map[string]int) {
	newParentEntityArrLength := len(parentEntityArr)
	newParentEntityArr := make([]*ParentEntity, newParentEntityArrLength)
	newParentEntryMap := make(map[string]int)
	for index, parentEntity := range parentEntityArr {
		newParentEntityArr[index] = &ParentEntity{
			Id:          parentEntity.ReplyId,
			IsComment:   false,
			IsReply:     true,
			replyEntity: parentEntity,
		}
		newParentEntryMap[parentEntity.ReplyId] = index
	}
	return newParentEntityArr, &newParentEntryMap
}

type ParentEntity struct {
	Id            string
	IsComment     bool
	IsReply       bool
	commentEntity *commentCommons.CommentEntity
	replyEntity   *commentCommons.ReplyEntity
}
