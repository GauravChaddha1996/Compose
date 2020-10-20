package deleteComment

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"errors"
)

func deleteComment(model *RequestModel) error {
	dao := daos.GetCommentDao()
	var changeMap = make(map[string]interface{})
	changeMap["is_deleted"] = 1

	err := dao.UpdateComment(model.CommentId, changeMap)
	if commons.InError(err) {
		return errors.New("Error in deleting comment record")
	}
	return nil
}
