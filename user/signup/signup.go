package signup

import (
	"compose/commons"
	"compose/dbModels"
	"compose/user/daos"
	"compose/user/userCommons"
	"errors"
	"github.com/raja/argon2pw"
	uuid "github.com/satori/go.uuid"
	"time"
)

func signup(requestModel *RequestModel) (string, error) {
	db := userCommons.Database
	transaction := db.Begin()
	userDao := daos.GetUserDaoUnderTransaction(transaction)
	passwordDao := daos.GetPasswordDaoUnderTransaction(transaction)
	accessTokenDao := daos.GetAccessTokenDaoUnderTransaction(transaction)

	// Email query

	_, err := userDao.FindUserViaEmail(requestModel.Email)
	if commons.InError(err) == false {
		transaction.Rollback()
		return "", errors.New("Email already present")
	}

	// Creating user

	userId, err := uuid.NewV4()
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("UUID can't be generated")
	}

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

	// Access token entry

	accessToken, err := uuid.NewV4()
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("Access token can't be generated")
	}
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

	transaction.Commit()
	return accessToken.String(), nil
}
