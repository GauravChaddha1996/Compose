package delete

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"errors"
	"github.com/rs/zerolog"
)

func deleteUser(model *RequestModel, logger *zerolog.Logger) error {
	var err error
	transaction := commons.GetDB().Begin()

	err = daos.GetPasswordDaoUnderTransaction(transaction).DeletePasswordEntryViaUserId(model.CommonModel.UserId)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Error deleting password entry of user")
	}
	logger.Info().Msg("Password entry of user is deleted")

	err = daos.GetAccessTokenDaoUnderTransaction(transaction).DeleteAccessTokenEntry(model.CommonModel.AccessToken)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Error deleting access token of user")
	}
	logger.Info().Msg("Access token entry is deleted")

	err = daos.GetUserDaoUnderTransaction(transaction).DeleteUser(model.Email)
	if commons.InError(err) {
		transaction.Rollback()
		return errors.New("Error deleting user")
	}
	logger.Info().Msg("User entry is deleted")

	transaction.Commit()
	return nil
}
