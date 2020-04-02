package unlikeArticle

import (
	"compose/commons"
	"compose/like/daos"
	"errors"
)

func unlikeArticle(model *RequestModel) error {
	likeDao := daos.GetLikeDao()

	previousLikeEntry, _ := likeDao.FindLikeEntry(model.ArticleId, model.CommonModel.UserId)
	if previousLikeEntry == nil {
		return errors.New("Article not liked")
	}

	err := likeDao.UnlikeArticle(previousLikeEntry)
	if commons.InError(err) {
		return errors.New("Article can't be un-liked")
	}
	return nil
}
