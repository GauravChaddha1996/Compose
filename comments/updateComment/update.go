package updateComment

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"compose/daos/commentAndReply"
	"errors"
)

func updateComment(model *RequestModel) error {
	transaction := commentCommons.Database.Begin()
	commentDao := daos.GetCommentDaoDuringTransaction(transaction)

	var markdownChangeMap = make(map[string]interface{})
	markdownChangeMap["markdown"] = model.Markdown
	err := commentDao.UpdateComment(model.CommentId, markdownChangeMap)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Error updating comment")
	}
	transaction.Commit()
	return nil
}
