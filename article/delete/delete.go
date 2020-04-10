package delete

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
)

func deleteArticle(article *dbModels.Article) error {
	tx := articleCommons.Database.Begin()

	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(tx)

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
	tx.Commit()
	return nil
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
