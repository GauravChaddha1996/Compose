package deleteReply

import (
	"compose/comments/daos"
	"compose/commons"
	"errors"
)

func deleteReply(model *RequestModel) error {
	dao := daos.GetReplyDao()
	var changeMap = make(map[string]interface{})
	changeMap["is_deleted"] = 1

	err := dao.UpdateReply(model.ReplyId, changeMap)
	if commons.InError(err) {
		return errors.New("Error in deleting reply record")
	}
	return nil
}
