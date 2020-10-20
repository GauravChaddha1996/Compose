package userDetails

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"compose/dataLayer/models"
	"errors"
)

func getUserDetails(model *RequestModel) (*models.User, error) {
	dao := daos.GetUserDao()
	user, err := dao.FindUserViaId(model.userId)
	if commons.InError(err) {
		return nil, errors.New("User entry not found ")
	}
	return user, nil
}
