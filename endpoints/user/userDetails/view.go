package userDetails

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/dbModels"
	"errors"
)

func getUserDetails(model *RequestModel) (*dbModels.User, error) {
	dao := daos.GetUserDao()
	user, err := dao.FindUserViaId(model.UserId)
	if commons.InError(err) {
		return nil, errors.New("User entry not found ")
	}
	return user, nil
}
