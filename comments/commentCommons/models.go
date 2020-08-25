package commentCommons

type CommentEntityType int
type ReplyEntityType int

const ReplyTypeErrorId = "reply_type_error_id"
const ReplyTypeContinueId = "reply_type_continue_id"
const CommentTypeEndId = "comment_type_end_id"

type CommentEntityTypeWrapper struct {
	CommentTypeNormal CommentEntityType
	CommentTypeEnd    CommentEntityType
}
type ReplyEntityTypeWrapper struct {
	ReplyTypeNormal   ReplyEntityType
	ReplyTypeError    ReplyEntityType
	ReplyTypeContinue ReplyEntityType
}

func NewCommentEntityTypeWrapper() CommentEntityTypeWrapper {
	return CommentEntityTypeWrapper{
		CommentTypeNormal: 0,
		CommentTypeEnd:    1,
	}
}

func NewReplyEntityTypeWrapper() ReplyEntityTypeWrapper {
	return ReplyEntityTypeWrapper{
		ReplyTypeNormal:   0,
		ReplyTypeError:    1,
		ReplyTypeContinue: 2,
	}
}

type CommentEntity struct {
	CommentType  CommentEntityType `json:"comment_type,omitempty"`
	CommentId    string            `json:"comment_id,omitempty"`
	Markdown     string            `json:"markdown,omitempty"`
	PostedByUser *PostedByUser     `json:"user,omitempty"`
	Replies      []ReplyEntity     `json:"replies,omitempty"`
}

type PostedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ReplyEntity struct {
	ReplyType    ReplyEntityType `json:"reply_type,omitempty"`
	ReplyId      string          `json:"reply_id,omitempty"`
	Markdown     string          `json:"markdown,omitempty"`
	PostedByUser *PostedByUser   `json:"user,omitempty"`
	Replies      []ReplyEntity   `json:"replies,omitempty"`
}

func GetErrorReplies() *[]ReplyEntity {
	return &[]ReplyEntity{GetErrorReplyEntity()}
}

func GetNoMoreCommentEntity(msg string) CommentEntity {
	return CommentEntity{
		CommentType:  NewCommentEntityTypeWrapper().CommentTypeEnd,
		CommentId:    CommentTypeEndId,
		Markdown:     msg,
		PostedByUser: nil,
		Replies:      nil,
	}
}

func GetErrorReplyEntity() ReplyEntity {
	return ReplyEntity{
		ReplyType:    NewReplyEntityTypeWrapper().ReplyTypeError,
		ReplyId:      ReplyTypeErrorId,
		Markdown:     "Error loading replies. Tap to try again.",
		PostedByUser: nil,
		Replies:      nil,
	}
}

func GetContinueReplyEntity() ReplyEntity {
	return ReplyEntity{
		ReplyType:    NewReplyEntityTypeWrapper().ReplyTypeContinue,
		ReplyId:      ReplyTypeContinueId,
		Markdown:     "Continue reading this thread ...",
		PostedByUser: nil,
		Replies:      nil,
	}
}
