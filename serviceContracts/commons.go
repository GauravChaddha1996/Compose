package serviceContracts

var userServiceContract UserServiceContract
var articleServiceContract ArticleServiceContract
var likeServiceContract LikeServiceContract
var commentServiceContract CommentServiceContract

func Init(userContract UserServiceContract, articleContract ArticleServiceContract, likeContract LikeServiceContract, commentContract CommentServiceContract) {
	userServiceContract = userContract
	articleServiceContract = articleContract
	likeServiceContract = likeContract
	commentServiceContract = commentContract
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

func GetCommentServiceContract() CommentServiceContract {
	return commentServiceContract
}
