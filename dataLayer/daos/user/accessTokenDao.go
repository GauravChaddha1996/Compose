package user

import (
	"compose/commons"
	"compose/dataLayer/dbModels"
	"gorm.io/gorm"
)

type AccessTokenDao struct {
	DB *gorm.DB
}

func (dao AccessTokenDao) CreateAccessTokenEntry(token dbModels.AccessToken) error {
	return dao.DB.Create(token).Error
}

func (dao AccessTokenDao) FindAccessTokenEntryViaUserId(userId string) (*dbModels.AccessToken, error) {
	var accessTokenEntry dbModels.AccessToken
	accessTokenQuery := dao.DB.Where("user_id = ?", userId).Find(&accessTokenEntry)
	if commons.InError(accessTokenQuery.Error) {
		return nil, accessTokenQuery.Error
	}
	return &accessTokenEntry, nil
}

func (dao AccessTokenDao) DeleteAccessTokenEntry(accessToken string) error {
	var accessTokenEntry dbModels.AccessToken
	return dao.DB.Where("access_token = ?", accessToken).Find(&accessTokenEntry).Unscoped().Delete(accessTokenEntry).Error
}
