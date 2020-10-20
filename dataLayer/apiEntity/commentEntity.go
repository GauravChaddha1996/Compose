package apiEntity

import (
	"compose/commons"
	"compose/dataLayer/dbModels"
)

type CommentEntityType int

const CommentTypeEndId = "comment_type_end_id"

type CommentEntityTypeWrapper struct {
	CommentTypeNormal  CommentEntityType
	CommentTypeEnd     CommentEntityType
	CommentTypeDeleted CommentEntityType
}

func NewCommentEntityTypeWrapper() CommentEntityTypeWrapper {
	return CommentEntityTypeWrapper{
		CommentTypeNormal:  0,
		CommentTypeEnd:     1,
		CommentTypeDeleted: 2,
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

func GetCommentEntityFromModel(comment *dbModels.Comment, user *PostedByUser) *CommentEntity {
	if comment == nil {
		return nil
	}
	if comment.IsDeleted == 1 {
		return GetDeletedCommentEntity(comment, user)
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

func GetDeletedCommentEntity(comment *dbModels.Comment, user *PostedByUser) *CommentEntity {
	return &CommentEntity{
		CommentType:  NewCommentEntityTypeWrapper().CommentTypeDeleted,
		CommentId:    comment.CommentId,
		Markdown:     "This comment has been deleted",
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

