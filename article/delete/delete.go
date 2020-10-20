package delete

import (
	"compose/article/articleCommons"
	"compose/commons"
	"compose/daos/article"
	"compose/dbModels"
	"errors"
)

func deleteArticle(article *dbModels.Article) error {
	tx := articleCommons.Database.Begin()

	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)

	err := markdownDao.DeleteArticleMarkdown(article.MarkdownId)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Cannot delete associated markdown")
	}

	err = articleCommons.UserServiceContract.ChangeArticleCount(article.UserId, false, tx) // change = false to decrease
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("User article count can't be decreased")
	}

	err = articleCommons.CommentServiceContract.DeleteAssociatedCommentsAndReplies(article.Id, tx)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Cannot delete associated comments and replies")
	}
	err = articleCommons.LikeServiceContract.DeleteAllLikeEntriesOfArticle(article.Id, tx)
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

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
