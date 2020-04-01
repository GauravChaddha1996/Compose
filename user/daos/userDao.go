package daos

import (
	"compose/commons"
	"compose/user/userCommons"
	"github.com/jinzhu/gorm"
)

func GetUserDao() *UserDao {
	return &UserDao{userCommons.Database}
}

func GetUserDaoUnderTransaction(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

type UserDao struct {
	db *gorm.DB
}

func (dao UserDao) CreateUser(user userCommons.User) error {
	return dao.db.Create(user).Error
}

func (dao UserDao) FindUserViaEmail(email string) (*userCommons.User, error) {
	var user userCommons.User
	userDeletionResult := dao.db.Where("email = ?", email).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao UserDao) FindUserViaId(userId string) (*userCommons.User, error) {
	var user userCommons.User
	userDeletionResult := dao.db.Where("user_id = ?", userId).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao UserDao) UpdateUser(changeMap map[string]interface{}, userId string) error {
	var user userCommons.User
	return dao.db.Model(user).Where("user_id = ?", userId).UpdateColumns(changeMap).Error
}

func (dao UserDao) DeleteUser(email string) error {
	var user userCommons.User
	return dao.db.Where("email = ?", email).Unscoped().Delete(&user).Error
}
