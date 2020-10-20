package comments

import (
	"compose/commons"
	"compose/daos/commentAndReply"
	"gorm.io/gorm"
)

type CommentServiceContractImpl struct{}

func GetCommentServiceContractImpl() CommentServiceContractImpl {
	return CommentServiceContractImpl{}
}

func (impl CommentServiceContractImpl) DeleteAssociatedCommentsAndReplies(articleId string, transaction *gorm.DB) error {
	commentDao := daos.GetCommentDaoDuringTransaction(transaction)
	replyDao := daos.GetReplyDaoDuringTransaction(transaction)
	err := commentDao.DeleteCommentsForArticle(articleId)
	if commons.InError(err) {
		return err
	}
	err = replyDao.DeleteRepliesForArticle(articleId)
	if commons.InError(err) {
		return err
	}
	return nil
}
