package delete

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
)

func delete(model *RequestModel) error {
	db := userCommons.GetDB()
	userDeletionResult := db.Unscoped().Delete(&model.user)
	if commons.InError(userDeletionResult.Error) {
		return errors.New("Deletion of user failed")
	}
	return nil
}
