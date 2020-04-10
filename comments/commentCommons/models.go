package commentCommons

type CommentEntity struct {
	CommentId     string          `json:"comment_id"`
	Markdown      string          `json:"markdown"`
	PostedByUser  PostedByUser    `json:"posted_by_user"`
	PostedAt      string          `json:"posted_at"`
	ChildComments []CommentEntity `json:"child_comments,omitempty"`
}

type PostedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	Name     string `json:"name,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
}
