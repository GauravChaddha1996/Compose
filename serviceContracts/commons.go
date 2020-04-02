package serviceContracts

var userServiceContract UserServiceContract
var articleServiceContract ArticleServiceContract

func Init(userContract UserServiceContract, articleContract ArticleServiceContract) {
	userServiceContract = userContract
	articleServiceContract = articleContract
}

func GetUserServiceContract() UserServiceContract {
	return userServiceContract
}

func GetArticleServiceContract() ArticleServiceContract {
	return articleServiceContract
}
