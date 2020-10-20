package apiEntity

import (
	"compose/commons"
	"compose/dataLayer/dbModels"
)

type ReplyEntityType int

const ReplyTypeContinueId = "reply_type_continue_id"
const ReplyTypeEndId = "reply_type_end_id"

type ReplyEntityTypeWrapper struct {
	ReplyTypeNormal   ReplyEntityType
	ReplyTypeContinue ReplyEntityType
	ReplyTypeEnd      ReplyEntityType
	ReplyTypeDeleted  ReplyEntityType
}

func NewReplyEntityTypeWrapper() ReplyEntityTypeWrapper {
	return ReplyEntityTypeWrapper{
		ReplyTypeNormal:   0,
		ReplyTypeContinue: 1,
		ReplyTypeEnd:      2,
		ReplyTypeDeleted:  3,
	}
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

func GetReplyEntityFromModel(reply *dbModels.Reply, user *PostedByUser) *ReplyEntity {
	if reply == nil {
		return nil
	}
	if reply.IsDeleted == 1 {
		return GetDeletedReplyEntity(reply, user)
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

func GetDeletedReplyEntity(reply *dbModels.Reply, user *PostedByUser) *ReplyEntity {
	return &ReplyEntity{
		ReplyType:    NewReplyEntityTypeWrapper().ReplyTypeDeleted,
		ReplyId:      reply.ReplyId,
		Markdown:     "This reply has been deleted",
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
