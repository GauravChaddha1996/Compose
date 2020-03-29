package update

import (
	"compose/commons"
	dao2 "compose/user/daos"
	"errors"
)

func update(requestModel *RequestModel) error {
	dao := dao2.GetUserDao()

	user, err := dao.FindUserViaId(requestModel.UserId)
	if commons.InError(err) {
		return errors.New("User query failed")
	}

	var changesMap = make(map[string]interface{})
	if requestModel.NewUserId != "" {
		changesMap["user_id"] = requestModel.NewUserId
	}
	if requestModel.Email != "" {
		changesMap["email"] = requestModel.Email
	}
	if requestModel.Name != "" {
		changesMap["name"] = requestModel.Name
	}
	changesMap["description"] = requestModel.Description
	changesMap["photo_url"] = requestModel.PhotoUrl

	err = dao.UpdateUser(changesMap, user)
	if commons.InError(err) {
		return errors.New("User update query failed")
	}
	return nil
}
