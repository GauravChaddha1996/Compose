package replyThread

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
)

const ReplyThreadLimit = 20
const ReplyThreadMaxLevel = 10

func getReplyThread(model *RequestModel) (*ResponseModel, error) {
	replyDao := daos.GetReplyDao()
	commentDao := daos.GetCommentDao()
	replies := replyDao.GetReplies(model.ParentId, ReplyThreadMaxLevel, 1, ReplyThreadLimit, model.CreatedAt, model.ReplyCount)

	if replies == nil {
		return &ResponseModel{
			Status:  commons.NewResponseStatus().SUCCESS,
			Message: "No more replies in this thread",
		}, nil
	}
	repliesEntity := *replies
	repliesLen := len(repliesEntity)
	newReplyCount := repliesLen + model.ReplyCount
	parentEntity := getParentEntity(model, replyDao, commentDao)

	// this means that we have more reply in this comment entry
	if uint64(newReplyCount) < parentEntity.ReplyCount {
		lastReplyEntity := repliesEntity[repliesLen-1]
		continueThreadPostbackParams := commentCommons.GetContinueThreadPostbackParams(model.ArticleId, model.ParentId, lastReplyEntity.PostedAt, newReplyCount)
		repliesEntity = append(repliesEntity, commentCommons.GetContinueReplyEntity(continueThreadPostbackParams))
	}
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Parent:  parentEntity,
		Replies: repliesEntity,
	}, nil
}

func getParentEntity(model *RequestModel, replyDao *daos.ReplyDao, commentDao *daos.CommentDao) *commentCommons.ParentEntity {
	parentReplyEntity, err := replyDao.FindReply(model.ParentId)
	if !commons.InError(err) {
		user, err := commentCommons.UserServiceContract.GetUser(parentReplyEntity.UserId)
		if commons.InError(err) {
			return nil
		}
		return &commentCommons.ParentEntity{
			ParentId:   model.ParentId,
			Markdown:   parentReplyEntity.Markdown,
			ReplyCount: parentReplyEntity.ReplyCount,
			PostedByUser: &commentCommons.PostedByUser{
				UserId:   user.UserId,
				PhotoUrl: user.PhotoUrl,
				Name:     user.Name,
			},
		}
	}
	parentCommentEntity, err := commentDao.FindComment(model.ParentId)
	if !commons.InError(err) {
		user, err := commentCommons.UserServiceContract.GetUser(parentCommentEntity.UserId)
		if commons.InError(err) {
			return nil
		}
		return &commentCommons.ParentEntity{
			ParentId:   model.ParentId,
			Markdown:   parentCommentEntity.Markdown,
			ReplyCount: parentCommentEntity.ReplyCount,
			PostedByUser: &commentCommons.PostedByUser{
				UserId:   user.UserId,
				PhotoUrl: user.PhotoUrl,
				Name:     user.Name,
			},
		}
	}
	return nil
}
