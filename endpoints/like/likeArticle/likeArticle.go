package likeArticle

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/models"
	"errors"
)

func likeArticle(model *RequestModel) error {
	tx := commons.GetDB().Begin()
	likeDao := daos.GetLikeDaoDuringTransaction(tx)
	userDao := daos.GetUserDaoUnderTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)
	var likeEntry = models.LikeEntry{
		UserId:    model.CommonModel.UserId,
		ArticleId: model.ArticleId,
	}

	previousLikeEntry, _ := likeDao.FindLikeEntry(likeEntry.ArticleId, likeEntry.UserId)
	if previousLikeEntry != nil {
		tx.Rollback()
		return errors.New("Article already liked")
	}

	err := articleDao.ChangeArticleLikeCount(model.ArticleId, true)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article like count can't be increased")
	}

	err = userDao.ChangeLikeCount(model.CommonModel.UserId, true)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("User like count can't be increased")
	}

	err = likeDao.LikeArticle(&likeEntry)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article can't be liked")
	}

	tx.Commit()
	return nil
}
