package createReply

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createReply(model *RequestModel) (*ResponseModel, error) {
	tx := commentCommons.Database.Begin()
	replyDao := daos.GetReplyDaoDuringTransaction(tx)

	err := commentCommons.ArticleServiceContract.ChangeArticleCommentCount(model.ArticleId, true, tx)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in increasing comment count of article")
	}

	replyUUId, err := uuid.NewV4()
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in generating reply uuid")
	}
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