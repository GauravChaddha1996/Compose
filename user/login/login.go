package login

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
	"github.com/raja/argon2pw"
)

func login(model *RequestModel) (string, error) {
	db := userCommons.GetDB()

	// Check if email exists
	var user userCommons.User
	emailQueryResult := db.Where("email = ?", model.email).Find(&user)
	if emailQueryResult.RecordNotFound() {
		return "", errors.New("Email doesn't exist")
	}

	// Deny login if account is inactive
	if user.IsActive == 0 {
		return "", errors.New("User is inactive")
	}

	// Match password
	var passwordEntry userCommons.Password
	passwordEntryResult := db.Where("user_id = ?", user.UserId).Find(&passwordEntry)
	if passwordEntryResult.RecordNotFound() {
		return "", errors.New("Password entry not found ")
	}

	_, err := argon2pw.CompareHashWithPassword(passwordEntry.PasswordHash, model.password)
	if commons.InError(err) {
		return "", errors.New("Password matching operation failure")
	}

	var accessTokenEntry userCommons.AccessToken
	accessTokenResult := db.Where("user_id = ?", user.UserId).Find(&accessTokenEntry)
	if accessTokenResult.RecordNotFound() {
		return "", errors.New("Access token not found")
	}

	return accessTokenEntry.AccessToken, nil
}
