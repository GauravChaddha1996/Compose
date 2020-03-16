package userCommons

import "time"

type User struct {
	UserId      string
	Email       string
	Name        string
	Description string
	PhotoUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

type Password struct {
	UserId       string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}

type AccessToken struct {
	UserId      string
	AccessToken string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
