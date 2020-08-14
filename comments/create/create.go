package create

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createComment(model *RequestModel) (*ResponseModel, error) {
	tx := commentCommons.Database.Begin()
	commentDao := daos.GetCommentDaoDuringTransaction(tx)

	err := commentCommons.ArticleServiceContract.ChangeArticleCommentCount(model.ArticleId, true, tx)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in increasing comment count of article")
	}

	commentUUId, err := uuid.NewV4()
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in generating comment uuid")
	}
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
