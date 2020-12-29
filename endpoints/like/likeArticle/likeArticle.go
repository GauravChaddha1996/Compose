package likeArticle

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/rs/zerolog"
)

func likeArticle(model *RequestModel, sublogger *zerolog.Logger) error {
	tx := commons.GetDB().Begin()
	likeDao := daos.GetLikeDaoDuringTransaction(tx)
	userDao := daos.GetUserDaoUnderTransaction(tx)
	articleDao := daos.GetArticleDaoDuringTransaction(tx)
	var likeEntry = dbModels.LikeEntry{
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
	sublogger.Info().Msg("Article like count changed")

	err = userDao.ChangeLikeCount(model.CommonModel.UserId, true)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("User like count can't be increased")
	}
	sublogger.Info().Msg("User total like count changed")

	err = likeDao.LikeArticle(&likeEntry)
	if commons.InError(err) {
		tx.Rollback()
		return errors.New("Article can't be liked")
	}
	sublogger.Info().Msg("Like entry is made")

	tx.Commit()
	return nil
}
