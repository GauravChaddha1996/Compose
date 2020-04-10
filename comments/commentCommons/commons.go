package commentCommons

import (
	"compose/serviceContracts"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB
var ArticleServiceContract serviceContracts.ArticleServiceContract
var UserServiceContract serviceContracts.UserServiceContract
