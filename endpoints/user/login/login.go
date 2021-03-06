package login

import (
	"compose/commons"
	"compose/dataLayer/daos"
	userDaos "compose/dataLayer/daos/user"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/raja/argon2pw"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"time"
)

func login(model *RequestModel, subLogger *zerolog.Logger) (string, error) {
	userDao := daos.GetUserDao()
	passwordDao := daos.GetPasswordDao()
	accessTokenDao := daos.GetAccessTokenDao()

	// Check if Email exists
	user, err := userDao.FindUserViaEmail(model.Email)
	if commons.InError(err) {
		return "", errors.New("Email doesn't exist")
	}
	subLogger.Info().Msg("User exists")

	// Match Password
	passwordEntry, err := passwordDao.FindPasswordEntryViaUserId(user.UserId)
	if commons.InError(err) {
		return "", errors.New("Password entry not found ")
	}

	_, err = argon2pw.CompareHashWithPassword(passwordEntry.PasswordHash, model.Password)
	if commons.InError(err) {
		return "", errors.New("Password matching operation failure")
	}
	subLogger.Info().Msg("Password matches")

	accessTokenEntry, err := accessTokenDao.FindAccessTokenEntryViaUserId(user.UserId)
	if commons.InError(err) || accessTokenEntry == nil {
		accessTokenEntry, err = createAccessTokenEntry(user, accessTokenDao)
		if commons.InError(err) {
			return "", errors.New("Access token generation failed")
		}
	}
	subLogger.Info().Msg("Access token found")

	return accessTokenEntry.AccessToken, nil
}

func createAccessTokenEntry(user *dbModels.User, accessTokenDao userDaos.AccessTokenDao) (*dbModels.AccessToken, error) {
	accessToken := uuid.NewV4()
	accessTokenEntry := dbModels.AccessToken{
		UserId:      user.UserId,
		AccessToken: accessToken.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := accessTokenDao.CreateAccessTokenEntry(accessTokenEntry)
	if commons.InError(err) {
		return nil, errors.New("Access token can't be saved")
	}
	return &accessTokenEntry, nil
}
