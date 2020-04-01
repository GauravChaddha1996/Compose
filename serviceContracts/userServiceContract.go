package serviceContracts

import "compose/user/userCommons"

type UserServiceContract interface {
	GetUser(userId string) (*userCommons.User, error)
}
