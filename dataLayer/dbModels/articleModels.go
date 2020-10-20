package dbModels

import (
	"time"
)

type ArticleMarkdown struct {
	Id       string
	Markdown string
}

type Article struct {
	Id                string
	UserId            string
	Title             string
	Description       string
	MarkdownId        string
	LikeCount         uint64
	TopCommentCount   uint64
	TotalCommentCount uint64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
