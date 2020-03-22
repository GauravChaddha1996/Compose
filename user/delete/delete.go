package delete

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
)

func deleteUser(model *RequestModel) error {
	db := userCommons.GetDB()
	var user userCommons.User
	userDeletionResult := db.Where("email = ?", model.email).Unscoped().Delete(&user)
	if commons.InError(userDeletionResult.Error) {
		return errors.New("Deletion of user failed")
	}
	return nil
}
