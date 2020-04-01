package daos

import (
	"compose/commons"
	"compose/user/userCommons"
	"github.com/jinzhu/gorm"
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

func (dao AccessTokenDao) CreateAccessTokenEntry(token userCommons.AccessToken) error {
	return dao.db.Create(token).Error
}

func (dao AccessTokenDao) FindAccessTokenEntryViaUserId(userId string) (*userCommons.AccessToken, error) {
	var accessTokenEntry userCommons.AccessToken
	accessTokenQuery := dao.db.Where("user_id = ?", userId).Find(&accessTokenEntry)
	if commons.InError(accessTokenQuery.Error) {
		return nil, accessTokenQuery.Error
	}
	return &accessTokenEntry, nil
}
