package daos

import (
	"compose/article/articleCommons"
	"compose/commons"
	"compose/dbModels"
	"github.com/jinzhu/gorm"
)

type CommentMarkdownDao struct {
	db *gorm.DB
}

func GetCommentMarkdownDaoDuringTransaction(db *gorm.DB) *CommentMarkdownDao {
	return &CommentMarkdownDao{db: db}
}

func GetCommentMarkdownDao() *CommentMarkdownDao {
	return &CommentMarkdownDao{db: articleCommons.Database}
}

func (dao CommentMarkdownDao) CreateCommentMarkdown(markdown dbModels.CommentMarkdown) error {
	return dao.db.Create(markdown).Error
}

func (dao CommentMarkdownDao) GetCommentMarkdown(markdownId string) (*dbModels.CommentMarkdown, error) {
	var markdown dbModels.CommentMarkdown
	markdownQuery := dao.db.Where("id = ?", markdownId).Find(&markdown)
	if commons.InError(markdownQuery.Error) {
		return nil, markdownQuery.Error
	}
	return &markdown, nil
}

func (dao CommentMarkdownDao) UpdateCommentMarkdown(markdownId string, changeMap map[string]interface{}) error {
	var markdown dbModels.CommentMarkdown
	return dao.db.Model(markdown).Where("id = ?", markdownId).UpdateColumns(changeMap).Error
}

func (dao CommentMarkdownDao) DeleteCommentMarkdown(markdownId string) error {
	var markdown dbModels.CommentMarkdown
	return dao.db.Where("id = ?", markdownId).Unscoped().Delete(&markdown).Error
}
