package daos

import (
	"compose/article/articleCommons"
	"github.com/jinzhu/gorm"
)

type MarkdownDao struct {
	db *gorm.DB
}

func GetMarkdownDaoDuringTransaction(db *gorm.DB) *MarkdownDao {
	return &MarkdownDao{db: db}
}

func (dao MarkdownDao) CreateMarkdown(markdown articleCommons.Markdown) error {
	return dao.db.Create(markdown).Error
}
