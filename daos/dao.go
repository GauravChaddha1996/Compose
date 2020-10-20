package daos

import (
	"compose/commons"
	"compose/daos/article"
	"compose/daos/commentAndReply"
	"compose/daos/like"
	"compose/daos/user"
	"gorm.io/gorm"
)

func GetUserDao() *user.UserDao {
	return &user.UserDao{DB: commons.GetDB()}
}

func GetUserDaoUnderTransaction(db *gorm.DB) *user.UserDao {
	return &user.UserDao{DB: db}
}

func GetAccessTokenDao() user.AccessTokenDao {
	return user.AccessTokenDao{DB: commons.GetDB()}
}
func GetAccessTokenDaoUnderTransaction(db *gorm.DB) user.AccessTokenDao {
	return user.AccessTokenDao{DB: db}
}

func GetPasswordDao() user.PasswordDao {
	return user.PasswordDao{DB: commons.GetDB()}
}
func GetPasswordDaoUnderTransaction(db *gorm.DB) user.PasswordDao {
	return user.PasswordDao{DB: db}
}

func GetArticleMarkdownDao() *article.ArticleMarkdownDao {
	return &article.ArticleMarkdownDao{DB: commons.GetDB()}
}

func GetArticleMarkdownDaoDuringTransaction(db *gorm.DB) *article.ArticleMarkdownDao {
	return &article.ArticleMarkdownDao{DB: db}
}

func GetArticleDao() *article.ArticleDao {
	return &article.ArticleDao{DB: commons.GetDB()}
}

func GetArticleDaoDuringTransaction(db *gorm.DB) *article.ArticleDao {
	return &article.ArticleDao{DB: db}
}

func GetLikeDao() like.LikeDao {
	return like.LikeDao{DB: commons.GetDB()}
}

func GetLikeDaoDuringTransaction(db *gorm.DB) like.LikeDao {
	return like.LikeDao{DB: db}
}

func GetReplyDao() *commentAndReply.ReplyDao {
	return &commentAndReply.ReplyDao{DB: commons.GetDB()}
}

func GetReplyDaoDuringTransaction(db *gorm.DB) *commentAndReply.ReplyDao {
	return &commentAndReply.ReplyDao{DB: db}
}

func GetCommentDao() *commentAndReply.CommentDao {
	return &commentAndReply.CommentDao{DB: commons.GetDB()}
}

func GetCommentDaoDuringTransaction(db *gorm.DB) *commentAndReply.CommentDao {
	return &commentAndReply.CommentDao{DB: db}
}
