package replyThread

import (
	"compose/commons"
	"compose/dataLayer/apiEntity"
	"compose/dataLayer/daos"
	commentAndReplyDaos "compose/dataLayer/daos/commentAndReply"
	userDaos "compose/dataLayer/daos/user"
	replyThreadCommon2 "compose/endpoints/comments/replyThreadCommon"
	"errors"
	"github.com/rs/zerolog"
)

const ReplyThreadMaxLevel = 10
const ReplyThreadRepliesMaxCount = 2000

func getReplyThread(model *RequestModel, subLogger *zerolog.Logger) (*ResponseModel, error) {
	replyDao := daos.GetReplyDao()
	commentDao := daos.GetCommentDao()
	userDao := daos.GetUserDao()
	parentEntity, parentCommentEntity, parentReplyEntity := apiEntity.GetCommentReplyParentEntity(model.ParentId, replyDao, commentDao)
	subLogger.Info().Msg("Parent entity found")

	replyEntityArr, err := getReplyEntityArr(model, parentCommentEntity, parentReplyEntity, replyDao, userDao)
	if commons.InError(err) {
		return nil, err
	}
	subLogger.Info().Msg("Reply entity arr is found")
	if len(replyEntityArr) == 0 {
		subLogger.Info().Msg("No more replies")
		return getNoReplyResponse(parentEntity), nil
	}

	parentEntityArr, parentEntryMap := replyThreadCommon2.GetParentEntityArrAndMapFromReplyEntityArr(replyEntityArr)
	replyThreadCommon2.FillReplyTreeInParentIdArr(model.ArticleId, ReplyThreadMaxLevel, ReplyThreadRepliesMaxCount, parentEntityArr, parentEntryMap, replyDao, userDao)
	subLogger.Info().Msg("Reply tree is filled")

	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Parent:  parentEntity,
		Replies: replyEntityArr,
	}, nil
}

func getReplyEntityArr(
	model *RequestModel,
	parentCommentEntity *apiEntity.CommentEntity,
	parentReplyEntity *apiEntity.ReplyEntity,
	replyDao *commentAndReplyDaos.ReplyDao, userDao *userDaos.UserDao) ([]*apiEntity.ReplyEntity, error) {

	replyThreadParentModel := &replyThreadCommon2.ReplyThreadParentModel{
		Id:            model.ParentId,
		IsComment:     parentCommentEntity != nil,
		IsReply:       parentReplyEntity != nil,
		CommentEntity: parentCommentEntity,
		ReplyEntity:   parentReplyEntity,
	}
	_, replyEntityArr, err := replyThreadCommon2.GetReplyEntityArr([]*replyThreadCommon2.ReplyThreadParentModel{replyThreadParentModel}, replyDao, userDao)
	if commons.InError(err) {
		return nil, errors.New("Error in fetching replies to parent")
	}
	return replyEntityArr, nil
}

func getNoReplyResponse(parentEntity *apiEntity.CommentReplyParentEntity) *ResponseModel {
	message := "No more replies to show"
	noMoreReplyEntity := apiEntity.GetNoMoreReplyEntity(message)
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "",
		Replies: []*apiEntity.ReplyEntity{&noMoreReplyEntity},
		Parent:  parentEntity,
	}
}