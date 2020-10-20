package unlikeArticle

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"errors"
)

func unlikeArticle(model *RequestModel) error {
	tx := commons.GetDB().Begin()
	likeDao := daos.GetLikeDaoDuringTransaction(tx)
	userDao := daos.GetUserDaoUnderTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)

	previousLikeEntry, _ := likeDao.FindLikeEntry(model.ArticleId, model.CommonModel.UserId)
	if previousLikeEntry == nil {
		tx.Rollback()
		return errors.New("Article not liked")
	}

	err := articleDao.ChangeArticleLikeCount(model.ArticleId, false)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article like count can't be decreased")
	}

	err = userDao.ChangeLikeCount(model.CommonModel.UserId, false)
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
