package createComment

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createComment(model *RequestModel, subLogger *zerolog.Logger) (*ResponseModel, error) {
	tx := commons.GetDB().Begin()
	commentDao := daos.GetCommentDaoDuringTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)

	err := articleDao.ChangeArticleTopCommentCount(model.ArticleId, true, tx)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in increasing comment count of article")
	}
	subLogger.Info().Msg("Article top comment count increased")

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
	subLogger.Info().Msg("Comment entry is created")

	tx.Commit()
	return &ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Comment created successfully",
	}, nil
}
