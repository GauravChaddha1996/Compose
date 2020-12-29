package apiEntity

import (
	"compose/commons"
	"compose/dataLayer/daos"
	commentAndReplyDaos "compose/dataLayer/daos/commentAndReply"
	"compose/dataLayer/dbModels"
)

type CommentReplyParentEntity struct {
	ParentId     string           `json:"parent_id,omitempty"`
	Markdown     string           `json:"markdown,omitempty"`
	PostedByUser *SmallUserEntity `json:"user,omitempty"`
	ReplyCount   uint64           `json:"total_replies,omitempty"`
}

func GetCommentReplyParentEntity(parentId string, replyDao *commentAndReplyDaos.ReplyDao, commentDao *commentAndReplyDaos.CommentDao) (*CommentReplyParentEntity, *CommentEntity, *ReplyEntity) {
	userDao := daos.GetUserDao()
	markdown := ""
	replyCount := uint64(0)
	user := &dbModels.User{}
	userId := ""

	parentReplyEntity, err := replyDao.FindReply(parentId)
	parentCommentEntity, err := commentDao.FindComment(parentId)

	if parentReplyEntity != nil {
		userId = parentReplyEntity.UserId
		markdown = parentReplyEntity.Markdown
		replyCount = parentReplyEntity.ReplyCount
	} else if parentCommentEntity != nil {
		userId = parentCommentEntity.UserId
		markdown = parentCommentEntity.Markdown
		replyCount = parentCommentEntity.ReplyCount
	} else {
		return nil, nil, nil
	}

	user, err = userDao.FindUserViaId(userId)
	if commons.InError(err) {
		return nil, nil, nil
	}
	postedByUser := GetSmallUserEntity(user)
	parentEntity := CommentReplyParentEntity{
		ParentId:     parentId,
		Markdown:     markdown,
		ReplyCount:   replyCount,
		PostedByUser: postedByUser,
	}
	return &parentEntity,
		GetCommentEntityFromModel(parentCommentEntity, postedByUser),
		GetReplyEntityFromModel(parentReplyEntity, postedByUser)
}
