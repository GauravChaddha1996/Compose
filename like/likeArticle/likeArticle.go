package likeArticle

import (
	"compose/commons"
	"compose/dbModels"
	"compose/like/daos"
	"compose/like/likeCommons"
	"errors"
)

func likeArticle(model *RequestModel) error {
	tx := likeCommons.Database.Begin()
	likeDao := daos.GetLikeDaoDuringTransaction(tx)
	var likeEntry = dbModels.LikeEntry{
		UserId:    model.CommonModel.UserId,
		ArticleId: model.ArticleId,
	}

	previousLikeEntry, _ := likeDao.FindLikeEntry(likeEntry.ArticleId, likeEntry.UserId)
	if previousLikeEntry != nil {
		tx.Rollback()
		return errors.New("Article already liked")
	}

	err := likeCommons.ArticleServiceContract.ChangeArticleLikeCount(model.ArticleId, true, tx)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article like count can't be increased")
	}

	err = likeCommons.UserServiceContract.ChangeLikeCount(model.CommonModel.UserId, true, tx)
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
