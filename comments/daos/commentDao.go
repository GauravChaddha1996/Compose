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

func (dao CommentDao) GetNextRootCommentsOfArticle(articleId string, maxCreatedAt *time.Time, limit int) (*[]dbModels.Comment, error) {
	var comments []dbModels.Comment
	query := dao.db.Where("article_id = ? && created_at < ? && level = 1", articleId, maxCreatedAt).
		Order("created_at desc").
		Limit(limit).
		Find(&comments)
	if commons.InError(query.Error) {
		return nil, errors.New("Error in fetching root comments of article")
	}
	return &comments, nil
}

func (dao CommentDao) CreateComment(comment dbModels.Comment) error {
	return dao.db.Create(&comment).Error
}
