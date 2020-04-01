package delete

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"errors"
)

func deleteArticle(article *articleCommons.Article) error {
	tx := articleCommons.Database.Begin()

	markdownDao := daos.GetMarkdownDaoDuringTransaction(tx)

	err := markdownDao.DeleteMarkdown(article.MarkdownId)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Cannot delete associated markdown")
	}

	err = articleCommons.UserServiceContract.ChangeArticleCount(article.UserId, false) // change = false to decrease
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("User article count can't be decreased")
	}
	tx.Commit()
	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
