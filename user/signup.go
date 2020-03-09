package user

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func signup(requestModel *SignupRequestModel) (string, string) {
	db := Db
	transaction := db.Begin()

	// todo think of ways to break this function and use context maybe for objects and error
	var u User
	emailQueryResult := transaction.Where("email = ?", requestModel.Email).Find(&u)
	if emailQueryResult.RecordNotFound() == false {
		transaction.Rollback()
		return "", ERROR_USER_EMAIL_ALREADY_PRESENT
	}

	userId, err := uuid.NewV4()
	if err != nil {
		transaction.Rollback()
		return "", ERROR_USDERID_GENERATION_FAILURE
	}

	u = User{
		UserId:    userId.String(),
		Email:     requestModel.Email,
		Name:      requestModel.Name,
		IsActive:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	saveQueryResult := transaction.Save(u)
	if saveQueryResult.Error != nil {
		transaction.Rollback()
		log.Println(saveQueryResult.Error)
		return "", ERROR_USER_DB_SAVE_FAILURE
	}

	passwordEntry := Password{
		UserId:       u.UserId,
		PasswordHash: requestModel.Password,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	savePasswordResult := transaction.Save(passwordEntry)
	if savePasswordResult.Error != nil {
		log.Print(savePasswordResult.Error)
		transaction.Rollback()
		return "", ERROR_USER_PASSWORD_SAVE_FAILURE
	}

	accessToken, err := uuid.NewV4()
	if err != nil {
		transaction.Rollback()
		return "", ERROR_ACCESS_TOKEN_GENERATION_FAILURE
	}
	accessTokenEntry := AccessToken{
		UserId:      u.UserId,
		AccessToken: accessToken.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	saveAccessTokenResult := transaction.Save(accessTokenEntry)
	if saveAccessTokenResult.Error != nil {
		log.Print(saveAccessTokenResult.Error)
		transaction.Rollback()
		return "", ERROR_ACCESS_TOKEN_SAVE_FAILURE
	}

	transaction.Commit()
	return accessToken.String(), ""
}
