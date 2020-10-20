package article

import (
	"compose/commons"
	"compose/dataLayer/models"
	"gorm.io/gorm"
)

type ArticleMarkdownDao struct {
	DB *gorm.DB
}

func (dao ArticleMarkdownDao) CreateArticleMarkdown(markdown models.ArticleMarkdown) error {
	return dao.DB.Create(markdown).Error
}

func (dao ArticleMarkdownDao) GetArticleMarkdown(markdownId string) (*models.ArticleMarkdown, error) {
	var markdown models.ArticleMarkdown
	markdownQuery := dao.DB.Where("id = ?", markdownId).Find(&markdown)
	if commons.InError(markdownQuery.Error) {
		return nil, markdownQuery.Error
	}
	return &markdown, nil
}

func (dao ArticleMarkdownDao) UpdateArticleMarkdown(markdownId string, changeMap map[string]interface{}) error {
	var markdown models.ArticleMarkdown
	return dao.DB.Model(markdown).Where("id = ?", markdownId).UpdateColumns(changeMap).Error
}

func (dao ArticleMarkdownDao) DeleteArticleMarkdown(markdownId string) error {
	var markdown models.ArticleMarkdown
	return dao.DB.Where("id = ?", markdownId).Unscoped().Delete(&markdown).Error
}
