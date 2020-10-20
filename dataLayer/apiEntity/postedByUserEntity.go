package apiEntity

import "compose/dataLayer/dbModels"

type PostedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

func GetPostedByUser(user *dbModels.User) *PostedByUser {
	return &PostedByUser{
		UserId:   user.UserId,
		PhotoUrl: user.PhotoUrl,
		Name:     user.Name,
	}
}
