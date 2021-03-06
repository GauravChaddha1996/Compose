package apiEntity

import "compose/dataLayer/dbModels"

type SmallUserEntity struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

func GetSmallUserEntity(user *dbModels.User) *SmallUserEntity {
	return &SmallUserEntity{
		UserId:   user.UserId,
		PhotoUrl: user.PhotoUrl,
		Name:     user.Name,
	}
}
