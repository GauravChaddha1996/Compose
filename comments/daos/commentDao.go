package daos

import (
	"compose/dbModels"
	"github.com/jinzhu/gorm"
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
