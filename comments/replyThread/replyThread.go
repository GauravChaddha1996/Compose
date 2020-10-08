package replyThread

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/comments/replyThreadCommon"
	"compose/commons"
	"compose/dbModels"
	"errors"
)

const ReplyThreadMaxLevel = 10
const ReplyThreadRepliesMaxCount = 2000

func getReplyThread(model *RequestModel) (*ResponseModel, error) {
	replyDao := daos.GetReplyDao()
	commentDao := daos.GetCommentDao()
	parentEntity, parentCommentEntity, parentReplyEntity := getParentEntity(model, replyDao, commentDao)

	replyEntityArr, err := getReplyEntityArr(model, parentCommentEntity, parentReplyEntity, replyDao)
	if commons.InError(err) {
		return nil, err
	}
	if len(replyEntityArr) == 0 {
		return getNoReplyResponse(parentEntity), nil
	}

	parentEntityArr, parentEntryMap := replyThreadCommon.GetParentEntityArrAndMapFromReplyEntityArr(replyEntityArr)
	replyThreadCommon.FillReplyTreeInParentIdArr(model.ArticleId, ReplyThreadMaxLevel, ReplyThreadRepliesMaxCount, parentEntityArr, parentEntryMap, replyDao)

	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Parent:  parentEntity,
		Replies: replyEntityArr,
	}, nil
}

func getReplyEntityArr(
	model *RequestModel,
	parentCommentEntity *commentCommons.CommentEntity,
	parentReplyEntity *commentCommons.ReplyEntity,
	replyDao *daos.ReplyDao) ([]*commentCommons.ReplyEntity, error) {

	replyThreadParentModel := &replyThreadCommon.ReplyThreadParentModel{
		Id:            model.ParentId,
		IsComment:     parentCommentEntity != nil,
		IsReply:       parentReplyEntity != nil,
		CommentEntity: parentCommentEntity,
		ReplyEntity:   parentReplyEntity,
	}
	_, replyEntityArr, err := replyThreadCommon.GetReplyEntityArr([]*replyThreadCommon.ReplyThreadParentModel{replyThreadParentModel}, replyDao)
	if commons.InError(err) {
		return nil, errors.New("Error in fetching replies to parent")
	}
	return replyEntityArr, nil
}

func getNoReplyResponse(parentEntity *commentCommons.ParentEntity) *ResponseModel {
	message := "No more replies to show"
	noMoreReplyEntity := commentCommons.GetNoMoreReplyEntity(message)
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "",
		Replies: []*commentCommons.ReplyEntity{&noMoreReplyEntity},
		Parent:  parentEntity,
	}
}

func getParentEntity(model *RequestModel, replyDao *daos.ReplyDao, commentDao *daos.CommentDao) (*commentCommons.ParentEntity, *commentCommons.CommentEntity, *commentCommons.ReplyEntity) {
	markdown := ""
	replyCount := uint64(0)
	user := &dbModels.User{}
	userId := ""

	parentReplyEntity, err := replyDao.FindReply(model.ParentId)
	parentCommentEntity, err := commentDao.FindComment(model.ParentId)
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

	user, err = commentCommons.UserServiceContract.GetUser(userId)
	if commons.InError(err) {
		return nil, nil, nil
	}
	postedByUser := &commentCommons.PostedByUser{
		UserId:   user.UserId,
		PhotoUrl: user.PhotoUrl,
		Name:     user.Name,
	}
	return &commentCommons.ParentEntity{
			ParentId:     model.ParentId,
			Markdown:     markdown,
			ReplyCount:   replyCount,
			PostedByUser: postedByUser,
		}, commentCommons.GetCommentEntityFromModel(parentCommentEntity, postedByUser),
		commentCommons.GetReplyEntityFromModel(parentReplyEntity, postedByUser)
}
