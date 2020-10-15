package serviceContracts

import (
	"compose/dbModels"
	"gorm.io/gorm"
)

type UserServiceContract interface {
	GetUser(userId string) (*dbModels.User, error)
	GetUsers(userIds []string) ([]*dbModels.User, error)
	ChangeArticleCount(userId string, change bool, transaction *gorm.DB) error // send change true to increase and false to decrease
	ChangeLikeCount(userId string, change bool, transaction *gorm.DB) error    // send change true to increase and false to decrease
}
