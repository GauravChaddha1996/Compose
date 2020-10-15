package userCommons

import (
	"compose/serviceContracts"
	"gorm.io/gorm"
)

var Database *gorm.DB
var ArticleService serviceContracts.ArticleServiceContract
var LikeService serviceContracts.LikeServiceContract
