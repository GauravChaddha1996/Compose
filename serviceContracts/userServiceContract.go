package serviceContracts

import "compose/user/userCommons"

type UserServiceContract interface {
	GetUser(userId string) (*userCommons.User, error)
	ChangeArticleCount(userId string, change bool) error // send change true to increase and false to decrease
}
