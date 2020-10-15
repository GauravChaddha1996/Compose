package articleCommons

import (
	"compose/serviceContracts"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB
var UserServiceContract serviceContracts.UserServiceContract
var CommentServiceContract serviceContracts.CommentServiceContract
var LikeServiceContract serviceContracts.LikeServiceContract
