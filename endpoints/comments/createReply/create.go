package createReply

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createReply(model *RequestModel) (*ResponseModel, error) {
	tx := commons.GetDB().Begin()
	replyDao := daos.GetReplyDaoDuringTransaction(tx)
	commentDao := daos.GetCommentDaoDuringTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)

	err := articleDao.ChangeArticleReplyCommentCount(model.ArticleId, true, tx)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in increasing comment count of article")
	}

	if model.ParentIsComment {
		// increase parent comment reply count
		err := commentDao.IncreaseReplyCount(model.ParentId)
		if commons.InError(err) {
			tx.Rollback()
			return nil, errors.New("Error in increasing reply count of parent comment")
		}
	} else if model.ParentIsReply {
		// increase parent reply - child reply count
		err := replyDao.IncreaseReplyCount(model.ParentId)
		if commons.InError(err) {
			tx.Rollback()
			return nil, errors.New("Error in increasing reply count of parent reply")
		}
	}

	replyUUId := uuid.NewV4()
	reply := dbModels.Reply{
		ReplyId:   replyUUId.String(),
		ParentId:  model.ParentId,
		ArticleId: model.ArticleId,
		UserId:    model.CommonModel.UserId,
		Markdown:  model.Markdown,
		LikeCount: 0,
		IsDeleted: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = replyDao.CreateReply(reply)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in saving reply")
	}

	tx.Commit()
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Reply created successfully",
	}, nil
}
