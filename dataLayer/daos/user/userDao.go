package user

import (
	"compose/commons"
	"compose/dataLayer/models"
	"errors"
	"gorm.io/gorm"
)

type UserDao struct {
	DB *gorm.DB
}

func (dao UserDao) CreateUser(user models.User) error {
	return dao.DB.Create(user).Error
}

func (dao UserDao) FindUserViaEmail(email string) (*models.User, error) {
	var user models.User
	userDeletionResult := dao.DB.Where("email = ?", email).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao UserDao) FindUserViaId(userId string) (*models.User, error) {
	var user models.User
	userDeletionResult := dao.DB.Where("user_id = ?", userId).Find(&user)
	if commons.InError(userDeletionResult.Error) {
		return nil, userDeletionResult.Error
	}
	return &user, nil
}

func (dao UserDao) FindUserViaIds(userIds []string) ([]*models.User, error) {
	parentIdQuery := ""
	userIdsLen := len(userIds)
	if userIdsLen == 0 {
		return []*models.User{}, nil
	}
	for index, parentId := range userIds {
		parentIdQuery += "\"" + parentId + "\""
		if index != userIdsLen-1 {
			parentIdQuery += ","
		}
	}
	whereQuery := "user_id IN (" + parentIdQuery + ")"

	var users []*models.User
	queryResult := dao.DB.Where(whereQuery).Find(&users)
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
	var user models.User
	return dao.DB.Model(user).Where("user_id = ?", userId).UpdateColumns(changeMap).Error
}

func (dao UserDao) DeleteUser(email string) error {
	var user models.User
	return dao.DB.Where("email = ?", email).Unscoped().Delete(&user).Error
}

/* Helper f()s */

func (dao UserDao) ChangeArticleCount(userId string, change bool) error {
	user, err := dao.FindUserViaId(userId)
	if commons.InError(err) {
		return errors.New("Can't find any such user")
	}
	if change {
		user.ArticleCount += 1
	} else {
		user.ArticleCount -= 1
	}

	var changeMap = make(map[string]interface{})
	changeMap["article_count"] = user.ArticleCount

	err = dao.UpdateUser(changeMap, userId)
	if commons.InError(err) {
		return errors.New("User article count can't be updated")
	}
	return nil
}

func (dao UserDao) ChangeLikeCount(userId string, change bool) error {
	user, err := dao.FindUserViaId(userId)
	if commons.InError(err) {
		return errors.New("Can't find any such user")
	}
	if change {
		user.LikeCount += 1
	} else {
		user.LikeCount -= 1
	}

	var changeMap = make(map[string]interface{})
	changeMap["like_count"] = user.LikeCount

	err = dao.UpdateUser(changeMap, userId)
	if commons.InError(err) {
		return errors.New("User like count can't be updated")
	}
	return nil
}
