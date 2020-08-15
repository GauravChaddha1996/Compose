package commentCommons

type CommentEntity struct {
	CommentId    string        `json:"comment_id,omitempty"`
	Markdown     string        `json:"markdown,omitempty"`
	PostedByUser PostedByUser  `json:"user,omitempty"`
	Replies      []ReplyEntity `json:"replies,omitempty"`
}

type PostedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ReplyEntity struct {
	ReplyId      string        `json:"reply_id,omitempty"`
	Markdown     string        `json:"markdown,omitempty"`
	PostedByUser PostedByUser  `json:"user,omitempty"`
	Replies      []ReplyEntity `json:"replies,omitempty"`
}
