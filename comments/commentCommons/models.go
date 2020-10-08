package commentCommons

import (
	"compose/commons"
	"compose/dbModels"
)

type CommentEntityType int
type ReplyEntityType int

const CommentTypeEndId = "comment_type_end_id"
const ReplyTypeContinueId = "reply_type_continue_id"
const ReplyTypeEndId = "reply_type_end_id"

type CommentEntityTypeWrapper struct {
	CommentTypeNormal CommentEntityType
	CommentTypeEnd    CommentEntityType
}
type ReplyEntityTypeWrapper struct {
	ReplyTypeNormal   ReplyEntityType
	ReplyTypeContinue ReplyEntityType
	ReplyTypeEnd      ReplyEntityType
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
		ReplyTypeContinue: 1,
		ReplyTypeEnd:      2,
	}
}

type CommentEntity struct {
	CommentType  CommentEntityType `json:"comment_type,omitempty"`
	CommentId    string            `json:"comment_id,omitempty"`
	Markdown     string            `json:"markdown,omitempty"`
	PostedByUser *PostedByUser     `json:"user,omitempty"`
	PostedAt     string            `json:"posted_at,omitempty"`
	Replies      []*ReplyEntity    `json:"replies,omitempty"`
	ReplyCount   uint64            `json:"-"`
}

type PostedByUser struct {
	UserId   string `json:"user_id,omitempty"`
	PhotoUrl string `json:"photo_url,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ReplyEntity struct {
	ReplyType              ReplyEntityType `json:"reply_type,omitempty"`
	ReplyId                string          `json:"reply_id,omitempty"`
	Markdown               string          `json:"markdown,omitempty"`
	PostedByUser           *PostedByUser   `json:"user,omitempty"`
	PostedAt               string          `json:"posted_at,omitempty"`
	ContinuePostbackParams string          `json:"continue_postback_params,omitempty"`
	Replies                []*ReplyEntity  `json:"replies,omitempty"`
	ReplyCount             uint64          `json:"-"`
}

type ParentEntity struct {
	ParentId     string        `json:"parent_id,omitempty"`
	Markdown     string        `json:"markdown,omitempty"`
	PostedByUser *PostedByUser `json:"user,omitempty"`
	ReplyCount   uint64        `json:"total_replies,omitempty"`
}

func GetCommentEntityFromModel(comment *dbModels.Comment, user *PostedByUser) *CommentEntity {
	if comment == nil {
		return nil
	}
	return &CommentEntity{
		CommentType:  NewCommentEntityTypeWrapper().CommentTypeNormal,
		CommentId:    comment.CommentId,
		Markdown:     comment.Markdown,
		PostedByUser: user,
		PostedAt:     comment.CreatedAt.Format(commons.TimeFormat),
		Replies:      []*ReplyEntity{},
		ReplyCount:   comment.ReplyCount,
	}
}

func GetNoMoreCommentEntity(msg string) CommentEntity {
	return CommentEntity{
		CommentType:  NewCommentEntityTypeWrapper().CommentTypeEnd,
		CommentId:    CommentTypeEndId,
		Markdown:     msg,
		PostedByUser: nil,
	}
}

func GetReplyEntityFromModel(reply *dbModels.Reply, user *PostedByUser) *ReplyEntity {
	if reply == nil {
		return nil
	}
	return &ReplyEntity{
		ReplyType:    NewReplyEntityTypeWrapper().ReplyTypeNormal,
		ReplyId:      reply.ReplyId,
		Markdown:     reply.Markdown,
		PostedByUser: user,
		PostedAt:     reply.CreatedAt.Format(commons.TimeFormat),
		Replies:      []*ReplyEntity{},
		ReplyCount:   reply.ReplyCount,
	}
}

func GetContinueReplyEntity(continuePostbackParams string) *ReplyEntity {
	return &ReplyEntity{
		ReplyType:              NewReplyEntityTypeWrapper().ReplyTypeContinue,
		ReplyId:                ReplyTypeContinueId,
		Markdown:               "Continue reading this thread ...",
		PostedByUser:           nil,
		ContinuePostbackParams: continuePostbackParams,
	}
}

func GetNoMoreReplyEntity(msg string) ReplyEntity {
	return ReplyEntity{
		ReplyType:    NewReplyEntityTypeWrapper().ReplyTypeEnd,
		ReplyId:      ReplyTypeEndId,
		Markdown:     msg,
		PostedByUser: nil,
	}
}
