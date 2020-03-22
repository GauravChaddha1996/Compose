package update

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
)

func update(requestModel *RequestModel) error {
	db := userCommons.GetDB()
	var user userCommons.User
	userQuery := db.Where("user_id = ?", requestModel.UserId).Find(&user)
	if commons.InError(userQuery.Error) {
		return errors.New("User query failed")
	}
	var userUpdateMap = make(map[string]interface{})
	if requestModel.NewUserId != "" {
		userUpdateMap["user_id"] = requestModel.NewUserId
	}
	if requestModel.Email != "" {
		userUpdateMap["email"] = requestModel.Email
	}
	if requestModel.Name != "" {
		userUpdateMap["name"] = requestModel.Name
	}
	userUpdateMap["description"] = requestModel.Description
	userUpdateMap["photo_url"] = requestModel.PhotoUrl
	userUpdateQuery := db.Model(&user).Updates(userUpdateMap)
	if userUpdateQuery.Error != nil {
		return errors.New("User update query failed")
	}
	return nil
}
