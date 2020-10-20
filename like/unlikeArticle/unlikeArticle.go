package unlikeArticle

import (
	"compose/commons"
	"compose/daos/like"
	"compose/like/likeCommons"
	"errors"
)

func unlikeArticle(model *RequestModel) error {
	tx := likeCommons.Database.Begin()
	likeDao := daos.GetLikeDaoDuringTransaction(tx)

	previousLikeEntry, _ := likeDao.FindLikeEntry(model.ArticleId, model.CommonModel.UserId)
	if previousLikeEntry == nil {
		tx.Rollback()
		return errors.New("Article not liked")
	}

	err := likeCommons.ArticleServiceContract.ChangeArticleLikeCount(model.ArticleId, false, tx)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article like count can't be decreased")
	}

	err = likeCommons.UserServiceContract.ChangeLikeCount(model.CommonModel.UserId, false, tx)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("User like count can't be decreased")
	}

	err = likeDao.UnlikeArticle(previousLikeEntry)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article can't be un-liked")
	}

	tx.Commit()
	return nil
}
