package userDetails

import (
	"compose/user/userCommons"
	"errors"
)

func getUserDetails(model *RequestModel) (*userCommons.User, error) {
	db := userCommons.GetDB()

	var user userCommons.User
	userQueryResult := db.Where("user_id = ?", model.userId).Find(&user)
	if userQueryResult.RecordNotFound() {
		return nil, errors.New("User entry not found ")
	}

	return &user, nil
}
