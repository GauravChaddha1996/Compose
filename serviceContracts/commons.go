package serviceContracts

var userServiceContract UserServiceContract
var articleServiceContract ArticleServiceContract
var likeServiceContract LikeServiceContract

func Init(userContract UserServiceContract, articleContract ArticleServiceContract, likeContract LikeServiceContract) {
	userServiceContract = userContract
	articleServiceContract = articleContract
	likeServiceContract = likeContract
}

func GetUserServiceContract() UserServiceContract {
	return userServiceContract
}

func GetArticleServiceContract() ArticleServiceContract {
	return articleServiceContract
}

func GetLikeServiceContract() LikeServiceContract {
	return likeServiceContract
}
