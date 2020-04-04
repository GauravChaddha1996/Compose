package user

import (
	"compose/commons"
	"compose/dbModels"
	"compose/user/daos"
	"errors"
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

func (impl ServiceContractImpl) ChangeArticleCount(userId string, change bool) error {
	user, err := impl.dao.FindUserViaId(userId)
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

	err = impl.dao.UpdateUser(changeMap, userId)
	if commons.InError(err) {
		return errors.New("User article count can't be updated")
	}
	return nil
}

func (impl ServiceContractImpl) ChangeLikeCount(userId string, change bool) error {
	user, err := impl.dao.FindUserViaId(userId)
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

	err = impl.dao.UpdateUser(changeMap, userId)
	if commons.InError(err) {
		return errors.New("User like count can't be updated")
	}
	return nil
}
