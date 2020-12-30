package signup

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
	"github.com/raja/argon2pw"
	"github.com/rs/zerolog"
	uuid "github.com/satori/go.uuid"
	"time"
)

func signup(requestModel *RequestModel, subLogger *zerolog.Logger) (string, error) {
	transaction := commons.GetDB().Begin()
	userDao := daos.GetUserDaoUnderTransaction(transaction)
	passwordDao := daos.GetPasswordDaoUnderTransaction(transaction)
	accessTokenDao := daos.GetAccessTokenDaoUnderTransaction(transaction)

	// Email query

	_, err := userDao.FindUserViaEmail(requestModel.Email)
	if commons.InError(err) == false {
		transaction.Rollback()
		return "", errors.New("Email already present")
	}
	subLogger.Info().Msg("User not present already")

	// Creating user

	userId := uuid.NewV4()

	user := dbModels.User{
		UserId:    userId.String(),
		Email:     requestModel.Email,
		Name:      requestModel.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = userDao.CreateUser(user)
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("User can't be created")
	}
	subLogger.Info().Msg("User entry created")

	// Password entry

	passwordHash, err := argon2pw.GenerateSaltedHash(requestModel.Password)
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("Password hash can't be generated")
	}

	passwordEntry := dbModels.Password{
		UserId:       user.UserId,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = passwordDao.CreatePasswordEntry(passwordEntry)
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("Password can't be saved")
	}
	subLogger.Info().Msg("Password entry created")

	// Access token entry

	accessToken := uuid.NewV4()
	accessTokenEntry := dbModels.AccessToken{
		UserId:      user.UserId,
		AccessToken: accessToken.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = accessTokenDao.CreateAccessTokenEntry(accessTokenEntry)
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("Access token can't be saved")
	}
	subLogger.Info().Msg("Access token entry created")

	transaction.Commit()
	return accessToken.String(), nil
}
