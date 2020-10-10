package daos

import (
	"compose/comments/commentCommons"
	"compose/commons"
	"compose/dbModels"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type CommentDao struct {
	db *gorm.DB
}

func GetCommentDao() *CommentDao {
	return &CommentDao{commentCommons.Database}
}

func GetCommentDaoDuringTransaction(db *gorm.DB) *CommentDao {
	return &CommentDao{db}
}

func (dao CommentDao) CreateComment(comment dbModels.Comment) error {
	return dao.db.Create(&comment).Error
}

func (dao CommentDao) ReadComments(articleId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.Comment, error) {
	var comments []dbModels.Comment
	commentsQuery := dao.db.
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
	var comment dbModels.Comment
	queryResult := dao.db.Select("comment_id").Where("comment_id = ?", commentId).Limit(1).Find(&comment)
	if commons.InError(queryResult.Error) {
		return false
	}
	return true
}

func (dao CommentDao) FindComment(commentId string) (*dbModels.Comment, error) {
	var comment dbModels.Comment
	queryResult := dao.db.Where("comment_id = ?", commentId).Limit(1).Find(&comment)
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
	var comment dbModels.Comment
	return dao.db.Model(comment).Where("comment_id = ?", commentId).UpdateColumns(changeMap).Error
}

func (dao CommentDao) DeleteCommentsForArticle(articleId string) error {
	var comments []dbModels.Comment
	return dao.db.Where("article_id = ?", articleId).Find(&comments).Unscoped().Delete(comments).Error
}
