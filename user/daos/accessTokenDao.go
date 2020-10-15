package daos

import (
	"compose/commons"
	"compose/dbModels"
	"compose/user/userCommons"
	"gorm.io/gorm"
)

type AccessTokenDao struct {
	db *gorm.DB
}

func GetAccessTokenDao() AccessTokenDao {
	return AccessTokenDao{db: userCommons.Database}
}
func GetAccessTokenDaoUnderTransaction(db *gorm.DB) AccessTokenDao {
	return AccessTokenDao{db}
}

func (dao AccessTokenDao) CreateAccessTokenEntry(token dbModels.AccessToken) error {
	return dao.db.Create(token).Error
}

func (dao AccessTokenDao) FindAccessTokenEntryViaUserId(userId string) (*dbModels.AccessToken, error) {
	var accessTokenEntry dbModels.AccessToken
	accessTokenQuery := dao.db.Where("user_id = ?", userId).Find(&accessTokenEntry)
	if commons.InError(accessTokenQuery.Error) {
		return nil, accessTokenQuery.Error
	}
	return &accessTokenEntry, nil
}

func (dao AccessTokenDao) DeleteAccessTokenEntry(accessToken string) error {
	var accessTokenEntry dbModels.AccessToken
	return dao.db.Where("access_token = ?", accessToken).Find(&accessTokenEntry).Unscoped().Delete(accessTokenEntry).Error
}
