package login

import (
	"compose/commons"
	"compose/user/daos"
	"errors"
	"github.com/raja/argon2pw"
)

func login(model *RequestModel) (string, error) {
	userDao := daos.GetUserDao()
	passwordDao := daos.GetPasswordDao()
	accessTokenDao := daos.GetAccessTokenDao()

	// Check if email exists
	user, err := userDao.FindUserViaEmail(model.email)
	if commons.InError(err) {
		return "", errors.New("Email doesn't exist")
	}

	// Match password
	passwordEntry, err := passwordDao.FindPasswordEntryViaUserId(user.UserId)
	if commons.InError(err) {
		return "", errors.New("Password entry not found ")
	}

	_, err = argon2pw.CompareHashWithPassword(passwordEntry.PasswordHash, model.password)
	if commons.InError(err) {
		return "", errors.New("Password matching operation failure")
	}

	accessTokenEntry, err := accessTokenDao.FindAccessTokenEntryViaUserId(user.UserId)
	if commons.InError(err) {
		return "", errors.New("Access token not found")
	}

	return accessTokenEntry.AccessToken, nil
}
