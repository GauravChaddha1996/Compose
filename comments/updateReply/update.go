package updateReply

import (
	"compose/commons"
	"compose/daos/commentAndReply"
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
