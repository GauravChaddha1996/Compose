package likeCommons

import (
	"compose/serviceContracts"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB
var ArticleServiceContract serviceContracts.ArticleServiceContract
