package dbModels

import "time"

type LikeEntry struct {
	Id        uint64
	UserId    string
	ArticleId string
	CreatedAt time.Time
}

func (LikeEntry) TableName() string {
	return "likes"
}
