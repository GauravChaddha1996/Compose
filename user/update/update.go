package update

import (
	"compose/commons"
	"compose/user/daos"
	"errors"
)

func update(requestModel *RequestModel) error {
	dao := daos.GetUserDao()

	_, err := dao.FindUserViaId(requestModel.UserId)
	if commons.InError(err) {
		return errors.New("User query failed")
	}

	var changesMap = make(map[string]interface{})
	if requestModel.NewUserId != nil {
		changesMap["user_id"] = *requestModel.NewUserId
	}
	if requestModel.Email != nil {
		changesMap["email"] = *requestModel.Email
	}
	if requestModel.Name != nil {
		changesMap["name"] = *requestModel.Name
	}
	if requestModel.Description != nil {
		changesMap["description"] = *requestModel.Description
	}
	if requestModel.PhotoUrl != nil {
		changesMap["photo_url"] = *requestModel.PhotoUrl
	}

	err = dao.UpdateUser(changesMap, requestModel.UserId)
	if commons.InError(err) {
		return errors.New("User update query failed")
	}
	return nil
}
