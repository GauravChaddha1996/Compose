package userDetails

import (
	"compose/commons"
	"compose/dbModels"
	dao2 "compose/user/daos"
	"errors"
)

func getUserDetails(model *RequestModel) (*dbModels.User, error) {
	dao := dao2.GetUserDao()
	user, err := dao.FindUserViaId(model.userId)
	if commons.InError(err) {
		return nil, errors.New("User entry not found ")
	}
	return user, nil
}
