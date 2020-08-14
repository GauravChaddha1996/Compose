package daos

import (
	"compose/commons"
	"compose/dbModels"
	"github.com/jinzhu/gorm"
	"time"
)

type CommentDao struct {
	db *gorm.DB
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
