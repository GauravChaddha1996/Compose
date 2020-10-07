package articleComments

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"compose/dbModels"
	"encoding/json"
	"errors"
)

const ArticleCommentLimit = 20
const MaxRepliesCount = 1000
const MaxCommentReplyLevel = 5

func getArticleComments(model *RequestModel) (*ResponseModel, error) {
	commentEntityArr, err := getCommentEntityArr(model)
	if commons.InError(err) {
		return nil, err
	}
	if len(commentEntityArr) == 0 {
		createdAtTime := ""
		if model.PostbackParams != nil {
			createdAtTime = model.PostbackParams.CreatedAt
		}
		return getNoCommentsResponse(createdAtTime), nil
	}

	fillCommentEntityArrWithReplies(model, commentEntityArr)

	postbackParams, hasMore := getPaginationData(model, commentEntityArr)
	return &ResponseModel{
		Status:           commons.NewResponseStatus().SUCCESS,
		Message:          "",
		CommentsPointers: commentEntityArr,
		PostbackParams:   postbackParams,
		HasMore:          hasMore,
	}, nil
}

func getCommentEntityArr(model *RequestModel) ([]*commentCommons.CommentEntity, error) {
	commentDao := daos.GetCommentDao()
	createdAtTime, err := commons.MaxTime()
	if commons.InError(err) {
		return nil, errors.New("Error in parsing max time")
	}
	if model.PostbackParams != nil {
		createdAtTime, err = commons.ParseTime(model.PostbackParams.CreatedAt)
		if commons.InError(err) {
			return nil, errors.New("Error in parsing created at time")
		}
	}

	commentDbModels, err := commentDao.ReadComments(model.ArticleId, createdAtTime, ArticleCommentLimit)
	if commons.InError(err) {
		return nil, errors.New("Error is fetching comments")
	}

	commentEntityArr := make([]*commentCommons.CommentEntity, len(*commentDbModels))
	PostedByUserArr, err := commentCommons.GetUsersForComments(commentDbModels)
	if commons.InError(err) {
		return nil, errors.New("Error in fetching users for comments")
	}
	for index, commentDbModel := range *commentDbModels {
		commentEntityArr[index] = &commentCommons.CommentEntity{
			CommentType:  commentCommons.NewCommentEntityTypeWrapper().CommentTypeNormal,
			CommentId:    commentDbModel.CommentId,
			Markdown:     commentDbModel.Markdown,
			PostedByUser: &(*PostedByUserArr)[index],
			PostedAt:     commentDbModel.CreatedAt.Format(commons.TimeFormat),
			Replies:      []*commentCommons.ReplyEntity{},
			ReplyCount:   commentDbModel.ReplyCount,
		}
	}
	return commentEntityArr, nil
}

func fillCommentEntityArrWithReplies(model *RequestModel, commentEntityArr []*commentCommons.CommentEntity) {
	currentReplyLevel := 0
	repliesCount := 0
	parentEntityArr, parentEntryMap := getParentEntityArrAndMapFromCommentEntityArr(commentEntityArr)

	repliesFinishReached := false
	breakDueToError := false
	for currentReplyLevel < MaxCommentReplyLevel && repliesCount < MaxRepliesCount && repliesFinishReached == false {
		replyDbModels, replyEntityArr, err := getReplyEntityArr(parentEntityArr)
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
				parentComment := parentEntity.commentEntity
				parentComment.Replies = append(parentComment.Replies, replyEntity)
			}
			if parentEntity.IsReply {
				parentReply := parentEntity.replyEntity
				parentReply.Replies = append(parentReply.Replies, replyEntity)
			}
		}

		parentEntityArr, parentEntryMap = getParentEntityArrAndMapFromReplyEntityArr(replyEntityArr)
		repliesCount += len(replyEntityArr)
		currentReplyLevel += 1
	}
	checkForContinueThread(repliesFinishReached, breakDueToError, model, parentEntityArr)
}

func getReplyEntityArr(parentEntityArr []*ParentEntity) ([]*dbModels.Reply, []*commentCommons.ReplyEntity, error) {
	replyDao := daos.GetReplyDao()
	parentEntityArrLen := len(parentEntityArr)
	parentIds := make([]string, parentEntityArrLen)
	for index, parentEntity := range parentEntityArr {
		parentIds[index] = parentEntity.Id
	}
	replyDbModels, err := replyDao.GetRepliesInParentIds(parentIds)
	if commons.InError(err) {
		return nil, nil, errors.New("Error in fetching replies for parent entity arr")
	}

	PostedByUserArr, err := commentCommons.GetUsersForRepliesCorrect(replyDbModels)
	if commons.InError(err) {
		return nil, nil, errors.New("Error in fetching users for comments")
	}

	replyDbModelsLen := len(replyDbModels)
	replyEntityArr := make([]*commentCommons.ReplyEntity, replyDbModelsLen)
	for index, replyDbModel := range replyDbModels {
		replyEntityArr[index] = &commentCommons.ReplyEntity{
			ReplyType:              commentCommons.ReplyEntityTypeWrapper{}.ReplyTypeNormal,
			ReplyId:                replyDbModel.ReplyId,
			Markdown:               replyDbModel.Markdown,
			PostedByUser:           &(*PostedByUserArr)[index],
			Replies:                []*commentCommons.ReplyEntity{},
			PostedAt:               replyDbModel.CreatedAt.Format(commons.TimeFormat),
			ContinuePostbackParams: "",
			ReplyCount:             replyDbModel.ReplyCount,
		}
	}
	return replyDbModels, replyEntityArr, nil
}

func checkForContinueThread(repliesFinishReached bool, breakDueToError bool, model *RequestModel, parentEntityArr []*ParentEntity) {
	if repliesFinishReached == false || breakDueToError {
		for _, parentEntity := range parentEntityArr {
			if parentEntity.IsComment {
				parentComment := parentEntity.commentEntity
				if parentComment.ReplyCount > 0 {
					repliesLen := len(parentComment.Replies)
					var createdAtTime string
					if repliesLen == 0 {
						createdAtTime = commons.MAX_TIME
					} else {
						createdAtTime = parentComment.Replies[repliesLen-1].PostedAt
					}
					continuePostbackParams := getContinueThreadPostbackParams(model.ArticleId, parentComment.CommentId, createdAtTime, repliesLen)
					continueReplyEntity := commentCommons.GetContinueReplyEntity(continuePostbackParams)
					parentComment.Replies = append(parentComment.Replies, &continueReplyEntity)
				}
			}
			if parentEntity.IsReply {
				parentReply := parentEntity.replyEntity
				if parentReply.ReplyCount > 0 {
					repliesLen := len(parentReply.Replies)
					var createdAtTime string
					if repliesLen == 0 {
						createdAtTime = commons.MAX_TIME
					} else {
						createdAtTime = parentReply.Replies[repliesLen-1].PostedAt
					}
					continuePostbackParams := getContinueThreadPostbackParams(model.ArticleId, parentReply.ReplyId, createdAtTime, repliesLen)
					continueReplyEntity := commentCommons.GetContinueReplyEntity(continuePostbackParams)
					parentReply.Replies = append(parentReply.Replies, &continueReplyEntity)
				}
			}
		}
	}
}

func getPaginationData(model *RequestModel, commentEntityArr []*commentCommons.CommentEntity) (string, bool) {
	totalTopCommentCount := commentCommons.ArticleServiceContract.GetArticleTopCommentCount(model.ArticleId)
	commentEntityArrLen := len(commentEntityArr)

	commentsServedTillNowCount := commentEntityArrLen
	if commentEntityArrLen == 0 {
		return "", false
	}
	if model.PostbackParams != nil {
		commentsServedTillNowCount += model.PostbackParams.Count
	}
	postbackParams := PostbackParams{
		Count:     commentsServedTillNowCount,
		CreatedAt: (commentEntityArr[commentEntityArrLen-1]).PostedAt,
	}
	postbackParamsStr, err := json.Marshal(postbackParams)
	if commons.InError(err) {
		return "", false
	} else {
		return string(postbackParamsStr), uint64(commentsServedTillNowCount) < totalTopCommentCount
	}
}
