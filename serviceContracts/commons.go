package serviceContracts

var userServiceContract UserServiceContract

func Init(userContract UserServiceContract) {
	userServiceContract = userContract
}

func GetUserServiceContract() UserServiceContract {
	return userServiceContract
}
