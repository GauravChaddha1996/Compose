package replyThread

import (
	"compose/commons"
	"compose/dataLayer/daos"
	commentAndReplyDaos "compose/dataLayer/daos/commentAndReply"
	userDaos "compose/dataLayer/daos/user"
	"compose/dataLayer/dbModels"
	commentCommons2 "compose/endpoints/comments/commentCommons"
	replyThreadCommon2 "compose/endpoints/comments/replyThreadCommon"
	"errors"
)

const ReplyThreadMaxLevel = 10
const ReplyThreadRepliesMaxCount = 2000

func getReplyThread(model *RequestModel) (*ResponseModel, error) {
	replyDao := daos.GetReplyDao()
	commentDao := daos.GetCommentDao()
	userDao := daos.GetUserDao()
	parentEntity, parentCommentEntity, parentReplyEntity := getParentEntity(model, replyDao, commentDao)

	replyEntityArr, err := getReplyEntityArr(model, parentCommentEntity, parentReplyEntity, replyDao, userDao)
	if commons.InError(err) {
		return nil, err
	}
	if len(replyEntityArr) == 0 {
		return getNoReplyResponse(parentEntity), nil
	}

	parentEntityArr, parentEntryMap := replyThreadCommon2.GetParentEntityArrAndMapFromReplyEntityArr(replyEntityArr)
	replyThreadCommon2.FillReplyTreeInParentIdArr(model.ArticleId, ReplyThreadMaxLevel, ReplyThreadRepliesMaxCount, parentEntityArr, parentEntryMap, replyDao, userDao)

	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Parent:  parentEntity,
		Replies: replyEntityArr,
	}, nil
}

func getReplyEntityArr(
	model *RequestModel,
	parentCommentEntity *commentCommons2.CommentEntity,
	parentReplyEntity *commentCommons2.ReplyEntity,
	replyDao *commentAndReplyDaos.ReplyDao, userDao *userDaos.UserDao) ([]*commentCommons2.ReplyEntity, error) {

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

func getNoReplyResponse(parentEntity *commentCommons2.ParentEntity) *ResponseModel {
	message := "No more replies to show"
	noMoreReplyEntity := commentCommons2.GetNoMoreReplyEntity(message)
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "",
		Replies: []*commentCommons2.ReplyEntity{&noMoreReplyEntity},
		Parent:  parentEntity,
	}
}

func getParentEntity(model *RequestModel, replyDao *commentAndReplyDaos.ReplyDao, commentDao *commentAndReplyDaos.CommentDao) (*commentCommons2.ParentEntity, *commentCommons2.CommentEntity, *commentCommons2.ReplyEntity) {
	userDao := daos.GetUserDao()
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

	user, err = userDao.FindUserViaId(userId)
	if commons.InError(err) {
		return nil, nil, nil
	}
	postedByUser := &commentCommons2.PostedByUser{
		UserId:   user.UserId,
		PhotoUrl: user.PhotoUrl,
		Name:     user.Name,
	}
	return &commentCommons2.ParentEntity{
			ParentId:     model.ParentId,
			Markdown:     markdown,
			ReplyCount:   replyCount,
			PostedByUser: postedByUser,
		}, commentCommons2.GetCommentEntityFromModel(parentCommentEntity, postedByUser),
		commentCommons2.GetReplyEntityFromModel(parentReplyEntity, postedByUser)
}
