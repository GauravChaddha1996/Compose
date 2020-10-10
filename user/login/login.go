package login

import (
	"compose/commons"
	"compose/dbModels"
	"compose/user/daos"
	"errors"
	"github.com/raja/argon2pw"
	uuid "github.com/satori/go.uuid"
	"time"
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
	if commons.InError(err) || accessTokenEntry == nil {
		accessTokenEntry, err = createAccessTokenEntry(user, accessTokenDao)
		if commons.InError(err) {
			return "", errors.New("Access token generation failed")
		}
	}

	return accessTokenEntry.AccessToken, nil
}

func createAccessTokenEntry(user *dbModels.User, accessTokenDao daos.AccessTokenDao) (*dbModels.AccessToken, error) {
	accessToken, err := uuid.NewV4()
	if commons.InError(err) {
		return nil, nil
	}
	accessTokenEntry := dbModels.AccessToken{
		UserId:      user.UserId,
		AccessToken: accessToken.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = accessTokenDao.CreateAccessTokenEntry(accessTokenEntry)
	if commons.InError(err) {
		return nil, errors.New("Access token can't be saved")
	}
	return &accessTokenEntry, nil
}
