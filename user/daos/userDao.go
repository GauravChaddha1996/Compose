package daos

import (
	"compose/commons"
	"compose/user/userCommons"
	"github.com/jinzhu/gorm"
)

func GetUserDao() *Dao {
	return &Dao{userCommons.GetDB()}
}

func GetUserDaoUnderTransation(db *gorm.DB) *Dao {
	return &Dao{db}
}

type Dao struct {
	db *gorm.DB
}

func (dao Dao) CreateUser(user userCommons.User) error {
	return dao.db.Create(user).Error
}

func (dao Dao) FindUserViaEmail(email string) (*userCommons.User, error) {
	var user userCommons.User
	userDeletionResult := dao.db.Where("email = ?", email).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao Dao) FindUserViaId(userId string) (*userCommons.User, error) {
	var user userCommons.User
	userDeletionResult := dao.db.Where("user_id = ?", userId).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao Dao) UpdateUser(changeMap map[string]interface{}, user *userCommons.User) error {
	return dao.db.Model(&user).Updates(changeMap).Error
}

func (dao Dao) DeleteUser(email string) error {
	var user userCommons.User
	return dao.db.Where("email = ?", email).Unscoped().Delete(&user).Error
}
