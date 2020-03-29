package userDetails

import (
	"compose/commons"
	dao2 "compose/user/daos"
	"compose/user/userCommons"
	"errors"
)

func getUserDetails(model *RequestModel) (*userCommons.User, error) {
	dao := dao2.GetUserDao()
	user, err := dao.FindUserViaId(model.userId)
	if commons.InError(err) {
		return nil, errors.New("User entry not found ")
	}
	return user, nil
}
