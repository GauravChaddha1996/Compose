package update

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
)

func updateArticle(model *RequestModel, article *dbModels.Article) error {
	transaction := articleCommons.Database.Begin()
	articleDao := daos.GetArticleDaoDuringTransaction(transaction)
	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(transaction)

	markdownEntry, err := markdownDao.GetArticleMarkdown(article.MarkdownId)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Associated markdown doesnt' exist")
	}

	var markdownChangeMap = make(map[string]interface{})
	if model.Markdown != nil {
		markdownChangeMap["markdown"] = *model.Markdown
	}
	if len(markdownChangeMap) > 0 {
		err := markdownDao.UpdateArticleMarkdown(markdownEntry.Id, markdownChangeMap)
		if commons.InError(err) {
			transaction.Rollback()
			return errors.New("ArticleMarkdown update operation failure")
		}
	}

	var articleChangeMap = make(map[string]interface{})
	if model.Title != nil {
		articleChangeMap["title"] = *model.Title
	}
	if model.Description != nil {
		articleChangeMap["description"] = *model.Description
	}

	if len(articleChangeMap) > 0 {
		err := articleDao.UpdateArticle(article.Id, articleChangeMap)
		if commons.InError(err) {
			transaction.Rollback()
			return errors.New("Article update operation failure")
		}
	}

	transaction.Commit()
	return nil
}
