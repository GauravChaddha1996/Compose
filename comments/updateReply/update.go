package updateReply

import (
	"compose/comments/daos"
	"compose/commons"
	"errors"
)

func updateReply(model *RequestModel) error {
	replyDao := daos.GetReplyDao()

	var markdownChangeMap = make(map[string]interface{})
	markdownChangeMap["markdown"] = model.Markdown
	err := replyDao.UpdateReply(model.ReplyId, markdownChangeMap)
	if commons.InError(err) {
		return errors.New("Error updating reply")
	}
	return nil
}
