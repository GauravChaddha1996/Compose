package daos

import (
	"compose/commons"
	"compose/dbModels"
	"compose/user/userCommons"
	"gorm.io/gorm"
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

func (dao UserDao) CreateUser(user dbModels.User) error {
	return dao.db.Create(user).Error
}

func (dao UserDao) FindUserViaEmail(email string) (*dbModels.User, error) {
	var user dbModels.User
	userDeletionResult := dao.db.Where("email = ?", email).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao UserDao) FindUserViaId(userId string) (*dbModels.User, error) {
	var user dbModels.User
	userDeletionResult := dao.db.Where("user_id = ?", userId).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao UserDao) FindUserViaIds(userIds []string) ([]*dbModels.User, error) {
	parentIdQuery := ""
	userIdsLen := len(userIds)
	if userIdsLen == 0 {
		return []*dbModels.User{}, nil
	}
	for index, parentId := range userIds {
		parentIdQuery += "\"" + parentId + "\""
		if index != userIdsLen-1 {
			parentIdQuery += ","
		}
	}
	whereQuery := "user_id IN (" + parentIdQuery + ")"

	var users []*dbModels.User
	queryResult := dao.db.Where(whereQuery).Find(&users)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return users, nil
}

func (dao UserDao) DoesUserIdExist(userId string) (bool, error) {
	_, err := dao.FindUserViaId(userId)
	if commons.InError(err) {
		return false, err
	} else {
		return true, nil
	}
}

func (dao UserDao) UpdateUser(changeMap map[string]interface{}, userId string) error {
	var user dbModels.User
	return dao.db.Model(user).Where("user_id = ?", userId).UpdateColumns(changeMap).Error
}

func (dao UserDao) DeleteUser(email string) error {
	var user dbModels.User
	return dao.db.Where("email = ?", email).Unscoped().Delete(&user).Error
}
