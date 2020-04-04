package dbModels

import (
	"time"
)

type User struct {
	UserId       string
	Email        string
	Name         string
	Description  string
	PhotoUrl     string
	ArticleCount uint64
	LikeCount    uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Password struct {
	UserId       string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AccessToken struct {
	UserId      string
	AccessToken string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
