package update

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/rs/zerolog"
)

func updateArticle(model *RequestModel, article *dbModels.Article, subLogger *zerolog.Logger) error {
	transaction := commons.GetDB().Begin()
	articleDao := daos.GetArticleDaoDuringTransaction(transaction)
	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(transaction)

	markdownEntry, err := markdownDao.GetArticleMarkdown(article.MarkdownId)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Associated markdown doesnt' exist")
	}
	subLogger.Info().Msg("Associated markdown found")

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
		subLogger.Info().Msg("Associated markdown updated")
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
	subLogger.Info().Msg("Article details updated")
	transaction.Commit()
	return nil
}
