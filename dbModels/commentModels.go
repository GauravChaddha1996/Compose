package dbModels

import "time"

type Comment struct {
	CommentId     string
	UserId        string
	ArticleId     string
	MarkdownId    string
	ParentId      string
	RootCommentId string
	Level         int64
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}