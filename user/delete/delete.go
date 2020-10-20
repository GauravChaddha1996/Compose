package delete

import (
	"compose/commons"
	"compose/daos/user"
	"compose/user/userCommons"
	"errors"
)

func deleteUser(model *RequestModel) error {
	var err error
	transaction := userCommons.Database.Begin()
	err = daos.GetPasswordDaoUnderTransaction(transaction).DeletePasswordEntryViaUserId(model.commonModel.UserId)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Error deleting password entry of user")
	}

	err = daos.GetAccessTokenDaoUnderTransaction(transaction).DeleteAccessTokenEntry(model.commonModel.AccessToken)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Error deleting access token of user")
	}
	transaction.Commit()
	return nil
}
