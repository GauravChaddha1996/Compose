package commentAndReply

import (
	"compose/commons"
	"compose/dataLayer/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type CommentDao struct {
	DB *gorm.DB
}

func (dao CommentDao) CreateComment(comment models.Comment) error {
	return dao.DB.Create(&comment).Error
}

func (dao CommentDao) ReadComments(articleId string, maxCreatedAtTime time.Time, limit int) (*[]models.Comment, error) {
	var comments []models.Comment
	commentsQuery := dao.DB.
		Where("article_id = ? && created_at < ?", articleId, maxCreatedAtTime).
		Order("created_at desc").
		Limit(limit).
		Find(&comments)
	if commons.InError(commentsQuery.Error) {
		return nil, commentsQuery.Error
	}
	return &comments, nil
}

func (dao CommentDao) DoesCommentExist(commentId string) bool {
	var comment models.Comment
	queryResult := dao.DB.Select("comment_id").Where("comment_id = ?", commentId).Limit(1).Find(&comment)
	if commons.InError(queryResult.Error) {
		return false
	}
	return true
}

func (dao CommentDao) FindComment(commentId string) (*models.Comment, error) {
	var comment models.Comment
	queryResult := dao.DB.Where("comment_id = ?", commentId).Limit(1).Find(&comment)
	if commons.InError(queryResult.Error) {
		return nil, queryResult.Error
	}
	return &comment, nil
}

func (dao CommentDao) IncreaseReplyCount(commentId string) error {
	comment, err := dao.FindComment(commentId)
	if commons.InError(err) {
		return errors.New("Cannot find comment for this comment id")
	}
	var changeMap = make(map[string]interface{})
	changeMap["reply_count"] = (*comment).ReplyCount + 1
	return dao.UpdateComment(commentId, changeMap)
}

func (dao CommentDao) UpdateComment(commentId string, changeMap map[string]interface{}) error {
	var comment models.Comment
	return dao.DB.Model(comment).Where("comment_id = ?", commentId).UpdateColumns(changeMap).Error
}

func (dao CommentDao) DeleteCommentsForArticle(articleId string) error {
	var comments []models.Comment
	return dao.DB.Where("article_id = ?", articleId).Find(&comments).Unscoped().Delete(comments).Error
}
