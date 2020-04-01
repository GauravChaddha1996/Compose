package daos

import (
	"compose/article/articleCommons"
	"compose/commons"
	"github.com/jinzhu/gorm"
)

type MarkdownDao struct {
	db *gorm.DB
}

func GetMarkdownDaoDuringTransaction(db *gorm.DB) *MarkdownDao {
	return &MarkdownDao{db: db}
}

func GetMarkdownDao() *MarkdownDao {
	return &MarkdownDao{db: articleCommons.Database}
}

func (dao MarkdownDao) CreateMarkdown(markdown articleCommons.Markdown) error {
	return dao.db.Create(markdown).Error
}

func (dao MarkdownDao) GetMarkdown(markdownId string) (*articleCommons.Markdown, error) {
	var markdown articleCommons.Markdown
	markdownQuery := dao.db.Where("id = ?", markdownId).Find(&markdown)
	if commons.InError(markdownQuery.Error) {
		return nil, markdownQuery.Error
	}
	return &markdown, nil
}

func (dao MarkdownDao) UpdateMarkdown(markdownId string, changeMap map[string]interface{}) error {
	var markdown articleCommons.Markdown
	return dao.db.Model(markdown).Where("id = ?", markdownId).UpdateColumns(changeMap).Error
}

func (dao MarkdownDao) DeleteMarkdown(markdownId string) error {
	var markdown articleCommons.Markdown
	return dao.db.Where("id = ?", markdownId).Unscoped().Delete(&markdown).Error
}
