package article

import (
	"compose/commons"
	"compose/dataLayer/dbModels"
	"gorm.io/gorm"
)

type ArticleMarkdownDao struct {
	DB *gorm.DB
}

func (dao ArticleMarkdownDao) CreateArticleMarkdown(markdown dbModels.ArticleMarkdown) error {
	return dao.DB.Create(markdown).Error
}

func (dao ArticleMarkdownDao) GetArticleMarkdown(markdownId string) (*dbModels.ArticleMarkdown, error) {
	var markdown dbModels.ArticleMarkdown
	markdownQuery := dao.DB.Where("id = ?", markdownId).Find(&markdown)
	if commons.InError(markdownQuery.Error) {
		return nil, markdownQuery.Error
	}
	return &markdown, nil
}

func (dao ArticleMarkdownDao) UpdateArticleMarkdown(markdownId string, changeMap map[string]interface{}) error {
	var markdown dbModels.ArticleMarkdown
	return dao.DB.Model(markdown).Where("id = ?", markdownId).UpdateColumns(changeMap).Error
}

func (dao ArticleMarkdownDao) DeleteArticleMarkdown(markdownId string) error {
	var markdown dbModels.ArticleMarkdown
	return dao.DB.Where("id = ?", markdownId).Unscoped().Delete(&markdown).Error
}
