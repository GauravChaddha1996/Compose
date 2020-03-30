package articleCommons

import (
	"time"
)

type Markdown struct {
	Id       string
	Markdown string
}

type Article struct {
	Id          string
	UserId      string
	Title       string
	Description string
	MarkdownId  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
