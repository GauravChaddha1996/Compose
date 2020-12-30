package update

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"errors"
	"github.com/rs/zerolog"
)

func update(requestModel *RequestModel, subLogger *zerolog.Logger) error {
	dao := daos.GetUserDao()

	_, err := dao.FindUserViaId(requestModel.UserId)
	if commons.InError2(err, subLogger) {
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
	if commons.InError2(err, subLogger) {
		return errors.New("User update query failed")
	}
	return nil
}
