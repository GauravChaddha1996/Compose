package signup

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
	"github.com/raja/argon2pw"
	uuid "github.com/satori/go.uuid"
	"time"
)

func signup(requestModel *RequestModel) (string, error) {
	db := userCommons.GetDB()
	transaction := db.Begin()

	var user userCommons.User

	// Email query

	emailQueryResult := transaction.Where("email = ?", requestModel.Email).Find(&user)
	if emailQueryResult.RecordNotFound() == false {
		transaction.Rollback()
		return "", errors.New("Email already present")
	}

	// Creating user

	userId, err := uuid.NewV4()
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("UUID can't be generated")
	}

	user = userCommons.User{
		UserId:    userId.String(),
		Email:     requestModel.Email,
		Name:      requestModel.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createUserResult := transaction.Create(user)
	if createUserResult.Error != nil {
		transaction.Rollback()
		return "", errors.New("User can't be created")
	}

	// Password entry

	passwordHash, err := argon2pw.GenerateSaltedHash(requestModel.Password)
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("Password hash can't be generated")
	}

	passwordEntry := userCommons.Password{
		UserId:       user.UserId,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	savePasswordResult := transaction.Save(passwordEntry)
	if commons.InError(savePasswordResult.Error) {
		transaction.Rollback()
		return "", errors.New("Password can't be saved")
	}

	// Access token entry

	accessToken, err := uuid.NewV4()
	if commons.InError(err) {
		transaction.Rollback()
		return "", errors.New("Access token can't be generated")
	}
	accessTokenEntry := userCommons.AccessToken{
		UserId:      user.UserId,
		AccessToken: accessToken.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	saveAccessTokenResult := transaction.Save(accessTokenEntry)
	if commons.InError(saveAccessTokenResult.Error) {
		transaction.Rollback()
		return "", errors.New("Access token can't be saved")
	}

	transaction.Commit()
	return accessToken.String(), nil
}
