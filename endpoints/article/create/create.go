package create

import (
	"compose/commons"
	daos "compose/daos"
	"compose/dbModels"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

func createArticle(model *RequestModel) (*string, error) {
	transaction := commons.GetDB().Begin()
	userDao := daos.GetUserDaoUnderTransaction(transaction)
	articleDao := daos.GetArticleDaoDuringTransaction(transaction)
	markdownDao := daos.GetArticleMarkdownDaoDuringTransaction(transaction)

	markdownUuid := uuid.NewV4()

	markdownEntry := dbModels.ArticleMarkdown{
		Id:       markdownUuid.String(),
		Markdown: model.markdown,
	}

	err := markdownDao.CreateArticleMarkdown(markdownEntry)
	if commons.InError(err) {
		return nil, errors.New("ArticleMarkdown entry can't be created")
	}

	articleUuid := uuid.NewV4()

	articleEntry := dbModels.Article{
		Id:          articleUuid.String(),
		UserId:      model.userId,
		Title:       model.title,
		Description: model.description,
		MarkdownId:  markdownEntry.Id,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = userDao.ChangeArticleCount(model.userId, true) // change = true to increase
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
