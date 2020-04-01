package user

import (
	"compose/commons"
	"compose/user/daos"
	"compose/user/userCommons"
)

type ServiceContractImpl struct{}

func GetServiceContractImpl() ServiceContractImpl {
	return ServiceContractImpl{}
}

func (ServiceContractImpl) GetUser(userId string) (*userCommons.User, error) {
	// Convert this into a handler call
	dao := daos.GetUserDao()
	user, err := dao.FindUserViaId(userId)
	if commons.InError(err) {
		return nil, err
	}
	return user, nil
}
