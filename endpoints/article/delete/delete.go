package delete

import (
	"compose/commons"
	"compose/daos"
	"compose/dbModels"
	"errors"
	"gorm.io/gorm"
)

func deleteArticle(article *dbModels.Article) error {
	tx := commons.GetDB().Begin()

	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)
	likeDao := daos.GetLikeDaoDuringTransaction(tx)
	userDao := daos.GetUserDaoUnderTransaction(tx)

	err := markdownDao.DeleteArticleMarkdown(article.MarkdownId)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Cannot delete associated markdown")
	}

	err = userDao.ChangeArticleCount(article.UserId, false) // change = false to decrease
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("User article count can't be decreased")
	}

	err = deleteAssociatedCommentsAndReplies(article.Id, tx)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Cannot delete associated comments and replies")
	}
	err = likeDao.DeleteAllLikesOfArticle(article.Id)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Cannot delete associated like entries")
	}

	err = articleDao.DeleteArticle(article)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Error in deleting article")
	}

	tx.Commit()
	return nil
}

func deleteAssociatedCommentsAndReplies(articleId string, transaction *gorm.DB) error {
	commentDao := daos.GetCommentDaoDuringTransaction(transaction)
	replyDao := daos.GetReplyDaoDuringTransaction(transaction)
	err := commentDao.DeleteCommentsForArticle(articleId)
	if commons.InError(err) {
		return err
	}
	err = replyDao.DeleteRepliesForArticle(articleId)
	if commons.InError(err) {
		return err
	}
	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
