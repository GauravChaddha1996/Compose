package user

import (
	"compose/commons"
	"compose/dataLayer/models"
	"gorm.io/gorm"
)

type PasswordDao struct {
	DB *gorm.DB
}

func (dao PasswordDao) CreatePasswordEntry(password models.Password) error {
	return dao.DB.Create(password).Error
}

func (dao PasswordDao) FindPasswordEntryViaUserId(userId string) (*models.Password, error) {
	var passwordEntry models.Password
	passwordEntryQuery := dao.DB.Where("user_id = ?", userId).Find(&passwordEntry)
	if commons.InError(passwordEntryQuery.Error) {
		return nil, passwordEntryQuery.Error
	}
	return &passwordEntry, nil
}

func (dao PasswordDao) DeletePasswordEntryViaUserId(userId string) error {
	var passwordEntry models.Password
	return dao.DB.Where("user_id = ?", userId).Unscoped().Delete(&passwordEntry).Error
}
