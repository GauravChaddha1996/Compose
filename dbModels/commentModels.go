package dbModels

import "time"

type Comment struct {
	CommentId string
	ArticleId string
	UserId    string
	Markdown  string
	LikeCount uint64
	IsDeleted uint
	CreatedAt time.Time
	UpdatedAt time.Time
}