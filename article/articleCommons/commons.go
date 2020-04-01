package articleCommons

import (
	"compose/serviceContracts"
	"github.com/jinzhu/gorm"
)

var Database *gorm.DB
var UserServiceContract serviceContracts.UserServiceContract
