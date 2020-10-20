package daos

import (
	"compose/commons"
	article2 "compose/dataLayer/daos/article"
	commentAndReply2 "compose/dataLayer/daos/commentAndReply"
	like2 "compose/dataLayer/daos/like"
	user2 "compose/dataLayer/daos/user"
	"gorm.io/gorm"
)

func GetUserDao() *user2.UserDao {
	return &user2.UserDao{DB: commons.GetDB()}
}

func GetUserDaoUnderTransaction(db *gorm.DB) *user2.UserDao {
	return &user2.UserDao{DB: db}
}

func GetAccessTokenDao() user2.AccessTokenDao {
	return user2.AccessTokenDao{DB: commons.GetDB()}
}
func GetAccessTokenDaoUnderTransaction(db *gorm.DB) user2.AccessTokenDao {
	return user2.AccessTokenDao{DB: db}
}

func GetPasswordDao() user2.PasswordDao {
	return user2.PasswordDao{DB: commons.GetDB()}
}
func GetPasswordDaoUnderTransaction(db *gorm.DB) user2.PasswordDao {
	return user2.PasswordDao{DB: db}
}

func GetArticleMarkdownDao() *article2.ArticleMarkdownDao {
	return &article2.ArticleMarkdownDao{DB: commons.GetDB()}
}

func GetArticleMarkdownDaoDuringTransaction(db *gorm.DB) *article2.ArticleMarkdownDao {
	return &article2.ArticleMarkdownDao{DB: db}
}

func GetArticleDao() *article2.ArticleDao {
	return &article2.ArticleDao{DB: commons.GetDB()}
}

func GetArticleDaoDuringTransaction(db *gorm.DB) *article2.ArticleDao {
	return &article2.ArticleDao{DB: db}
}

func GetLikeDao() like2.LikeDao {
	return like2.LikeDao{DB: commons.GetDB()}
}

func GetLikeDaoDuringTransaction(db *gorm.DB) like2.LikeDao {
	return like2.LikeDao{DB: db}
}

func GetReplyDao() *commentAndReply2.ReplyDao {
	return &commentAndReply2.ReplyDao{DB: commons.GetDB()}
}

func GetReplyDaoDuringTransaction(db *gorm.DB) *commentAndReply2.ReplyDao {
	return &commentAndReply2.ReplyDao{DB: db}
}

func GetCommentDao() *commentAndReply2.CommentDao {
	return &commentAndReply2.CommentDao{DB: commons.GetDB()}
}

func GetCommentDaoDuringTransaction(db *gorm.DB) *commentAndReply2.CommentDao {
	return &commentAndReply2.CommentDao{DB: db}
}
