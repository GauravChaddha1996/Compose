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
	markdownDao := daos.GetCommentMarkdownDaoDuringTransaction(tx)

	markdownUUId, err := uuid.NewV4()
	if commons.InError(err) {
		return nil, errors.New("Error in generating markdown uuid")
	}

	commentMarkdown := dbModels.CommentMarkdown{
		Id:       markdownUUId.String(),
		Markdown: model.Markdown,
	}
	err = markdownDao.CreateCommentMarkdown(commentMarkdown)
	if commons.InError(err) {
		tx.Rollback()
		return nil, errors.New("Error in saving comment markdown")
	}

	err = commentCommons.ArticleServiceContract.ChangeArticleCommentCount(model.ArticleId, true, tx)
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
		CommentId:     commentUUId.String(),
		UserId:        model.CommonModel.UserId,
		ArticleId:     model.ArticleId,
		MarkdownId:    commentMarkdown.Id,
		ParentId:      "",
		RootCommentId: "",
		Level:         1,
		CreatedAt:     time.Now(),
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
