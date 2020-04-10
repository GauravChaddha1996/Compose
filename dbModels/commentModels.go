package dbModels

import "time"

type CommentMarkdown struct {
	Id       string
	Markdown string
}

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