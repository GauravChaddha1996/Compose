package create

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createArticle(model *RequestModel) (*string, error) {
	database := articleCommons.Database
	transaction := database.Begin()
	articleDao := daos.GetArticleDaoDuringTransaction(transaction)
	markdownDao := daos.GetMarkdownDaoDuringTransaction(transaction)

	markdownUuid, err := uuid.NewV4()
	if commons.InError(err) {
		transaction.Rollback()
		return nil, errors.New("Markdown UUID can't be generated")
	}

	markdownEntry := dbModels.Markdown{
		Id:       markdownUuid.String(),
		Markdown: model.markdown,
	}

	err = markdownDao.CreateMarkdown(markdownEntry)
	if commons.InError(err) {
		return nil, errors.New("Markdown entry can't be created")
	}

	articleUuid, err := uuid.NewV4()
	if commons.InError(err) {
		transaction.Rollback()
		return nil, errors.New("Article UUID can't be generated")
	}

	articleEntry := dbModels.Article{
		Id:          articleUuid.String(),
		UserId:      model.userId,
		Title:       model.title,
		Description: model.description,
		MarkdownId:  markdownEntry.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = articleCommons.UserServiceContract.ChangeArticleCount(model.userId, true) // change = true to increase
	if commons.InError(err) {
		transaction.Rollback()
		return nil, errors.New("User article count can't be increased")
	}

	err = articleDao.CreateArticle(articleEntry)
	if commons.InError(err) {
		transaction.Rollback()
		return nil, errors.New("Article entry can't be created")
	}

	transaction.Commit()
	return &articleEntry.Id, nil
}
