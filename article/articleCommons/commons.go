package articleCommons

import (
	"compose/serviceContracts"
	"gorm.io/gorm"
)

var Database *gorm.DB
var UserServiceContract serviceContracts.UserServiceContract
var CommentServiceContract serviceContracts.CommentServiceContract
var LikeServiceContract serviceContracts.LikeServiceContract
