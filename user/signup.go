package user

import (
	uuid "github.com/satori/go.uuid"
	"log"
	"time"
)

func signup(requestModel *UserSignupRequestModel) (string, string) {
	/*
		- Check email shouldn't exist already
		- Check if email is marked inactive
		- Make user entry
		- Make password entry
		- Make access token entry
		- Revert if any transaction fails
		- Do db entry using goroutine so db isn’t blocked
	*/

	db := Db

	var u User
	emailQueryResult := db.Where("email = ?", requestModel.Email).Find(&u)
	if emailQueryResult.RecordNotFound() == false {
		return "", SIGNUP_ERROR_USER_EMAIL_ALREADY_PRESENT
	}

	userId, err := uuid.NewV4()
	if err != nil {
		return "", SIGNUP_ERROR_USDERID_GENERATION_FAILURE
	}

	u = User{
		UserId:    userId.String(),
		Email:     requestModel.Email,
		Name:      requestModel.Name,
		IsActive:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	saveQueryResult := db.Save(u)
	if saveQueryResult.Error != nil {
		log.Println(saveQueryResult.Error)
		return "", SIGNUP_ERROR_USER_DB_SAVE_FAILURE
	}

	accessToken, err := uuid.NewV4()
	if err != nil {
		return "", SIGNUP_ERROR_ACCESS_TOKEN_GENERATION_FAILURE
	}
	return accessToken.String(), ""
}
