package user

import (
	"time"
)

type User struct {
	UserId      string `gorm:"PRIMARY_KEY;NOT NULL"`
	Email       string `gorm:"NOT NULL;UNIQUE"`
	Name        string `gorm:"NOT NULL"`
	IsActive    int    `gorm:"NOT NULL"`
	Description string
	PhotoUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
