package userCommons

import (
	"compose/serviceContracts"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB
var ArticleService serviceContracts.ArticleServiceContract
