package user

import (
	"compose/commons"
	"compose/dbModels"
	"compose/user/daos"
	"errors"
	"github.com/jinzhu/gorm"
)

type ServiceContractImpl struct {
	dao *daos.UserDao
}

func GetServiceContractImpl() ServiceContractImpl {
	return ServiceContractImpl{
		dao: daos.GetUserDao(),
	}
}

func (impl ServiceContractImpl) GetUser(userId string) (*dbModels.User, error) {
	// Convert this into a handler call
	user, err := impl.dao.FindUserViaId(userId)
	if commons.InError(err) {
		return nil, err
	}
	return user, nil
}

func (impl ServiceContractImpl) GetUsers(userIds []string) ([]*dbModels.User, error) {
	var users = make([]*dbModels.User, len(userIds))

	for index := range userIds {
		user, err := impl.dao.FindUserViaId(userIds[index])
		if commons.InError(err) {
			return nil, err
		}
		users[index] = user
	}
	return users, nil
}

func (impl ServiceContractImpl) ChangeArticleCount(userId string, change bool, transaction *gorm.DB) error {
	userDao := impl.getUserDao(transaction)
	user, err := userDao.FindUserViaId(userId)
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

	err = userDao.UpdateUser(changeMap, userId)
	if commons.InError(err) {
		return errors.New("User article count can't be updated")
	}
	return nil
}

func (impl ServiceContractImpl) ChangeLikeCount(userId string, change bool, transaction *gorm.DB) error {
	userDao := impl.getUserDao(transaction)
	user, err := userDao.FindUserViaId(userId)
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

	err = userDao.UpdateUser(changeMap, userId)
	if commons.InError(err) {
		return errors.New("User like count can't be updated")
	}
	return nil
}

func (impl ServiceContractImpl) getUserDao(transaction *gorm.DB) *daos.UserDao {
	if transaction != nil {
		return daos.GetUserDaoUnderTransaction(transaction)
	} else {
		return impl.dao
	}
}
