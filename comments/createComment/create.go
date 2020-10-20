package createComment

import (
	"compose/commons"
	"compose/daos"
	"compose/dbModels"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createComment(model *RequestModel) (*ResponseModel, error) {
	tx := commons.GetDB().Begin()
	commentDao := daos.GetCommentDaoDuringTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)

	err := articleDao.ChangeArticleTopCommentCount(model.ArticleId, true, tx)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in increasing comment count of article")
	}

	commentUUId := uuid.NewV4()

	comment := dbModels.Comment{
		CommentId: commentUUId.String(),
		ArticleId: model.ArticleId,
		UserId:    model.CommonModel.UserId,
		Markdown:  model.Markdown,
		LikeCount: 0,
		IsDeleted: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = commentDao.CreateComment(comment)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in saving comment")
	}

	tx.Commit()
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Comment created successfully",
	}, nil
}
