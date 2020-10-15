package likeCommons

import (
	"compose/serviceContracts"
	"gorm.io/gorm"
)

var Database *gorm.DB
var ArticleServiceContract serviceContracts.ArticleServiceContract
var UserServiceContract serviceContracts.UserServiceContract
