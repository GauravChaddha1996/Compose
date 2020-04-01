package daos

import (
	"compose/commons"
	"compose/user/userCommons"
	"github.com/jinzhu/gorm"
)

type PasswordDao struct {
	db *gorm.DB
}

func GetPasswordDao() PasswordDao {
	return PasswordDao{db: userCommons.Database}
}
func GetPasswordDaoUnderTransaction(db *gorm.DB) PasswordDao {
	return PasswordDao{db}
}

func (dao PasswordDao) CreatePasswordEntry(password userCommons.Password) error {
	return dao.db.Create(password).Error
}

func (dao PasswordDao) FindPasswordEntryViaUserId(userId string) (*userCommons.Password, error) {
	var passwordEntry userCommons.Password
	passwordEntryQuery := dao.db.Where("user_id = ?", userId).Find(&passwordEntry)
	if commons.InError(passwordEntryQuery.Error) {
		return nil, passwordEntryQuery.Error
	}
	return &passwordEntry, nil
}
