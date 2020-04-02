package likeArticle

import (
	"compose/commons"
	"compose/like/daos"
	"compose/like/likeCommons"
	"errors"
)

func likeArticle(model *RequestModel) error {
	likeDao := daos.GetLikeDao()
	var likeEntry = likeCommons.LikeEntry{
		UserId:    model.CommonModel.UserId,
		ArticleId: model.ArticleId,
	}

	previousLikeEntry, _ := likeDao.FindLikeEntry(likeEntry.ArticleId, likeEntry.UserId)
	if previousLikeEntry != nil {
		return errors.New("Article already liked")
	}

	err := likeDao.LikeArticle(&likeEntry)
	if commons.InError(err) {
		return errors.New("Article can't be liked")
	}
	return nil
}
