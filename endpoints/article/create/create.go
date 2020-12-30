package create

import (
	"compose/commons"
	daos "compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createArticle(model *RequestModel, subLogger *zerolog.Logger) (*string, error) {
	transaction := commons.GetDB().Begin()
	userDao := daos.GetUserDaoUnderTransaction(transaction)
	articleDao := daos.GetArticleDaoDuringTransaction(transaction)
	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(transaction)

	markdownUuid := uuid.NewV4()

	markdownEntry := dbModels.ArticleMarkdown{
		Id:       markdownUuid.String(),
		Markdown: model.Markdown,
	}

	err := markdownDao.CreateArticleMarkdown(markdownEntry)
	if commons.InError(err) {
		return nil, errors.New("ArticleMarkdown entry can't be created")
	}
	subLogger.Info().Msg("Article markdown entry created")

	articleUuid := uuid.NewV4()

	articleEntry := dbModels.Article{
		Id:          articleUuid.String(),
		UserId:      model.UserId,
		Title:       model.Title,
		Description: model.Description,
		MarkdownId:  markdownEntry.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = userDao.ChangeArticleCount(model.UserId, true) // change = true to increase
	if commons.InError(err) {
		transaction.Rollback()
		return nil, errors.New("User article count can't be increased")
	}
	subLogger.Info().Msg("User article count increased")

	err = articleDao.CreateArticle(articleEntry)
	if commons.InError(err) {
		transaction.Rollback()
		return nil, errors.New("Article entry can't be created")
	}
	subLogger.Info().Msg("Article entry created")

	transaction.Commit()
	return &articleEntry.Id, nil
}
