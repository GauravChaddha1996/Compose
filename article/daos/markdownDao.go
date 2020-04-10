package daos

import (
	"compose/article/articleCommons"
	"compose/commons"
	"compose/dbModels"
	"github.com/jinzhu/gorm"
)

type ArticleMarkdownDao struct {
	db *gorm.DB
}

func GetArticleMarkdownDaoDuringTransaction(db *gorm.DB) *ArticleMarkdownDao {
	return &ArticleMarkdownDao{db: db}
}

func GetArticleMarkdownDao() *ArticleMarkdownDao {
	return &ArticleMarkdownDao{db: articleCommons.Database}
}

func (dao ArticleMarkdownDao) CreateArticleMarkdown(markdown dbModels.ArticleMarkdown) error {
	return dao.db.Create(markdown).Error
}

func (dao ArticleMarkdownDao) GetArticleMarkdown(markdownId string) (*dbModels.ArticleMarkdown, error) {
	var markdown dbModels.ArticleMarkdown
	markdownQuery := dao.db.Where("id = ?", markdownId).Find(&markdown)
	if commons.InError(markdownQuery.Error) {
		return nil, markdownQuery.Error
	}
	return &markdown, nil
}

func (dao ArticleMarkdownDao) UpdateArticleMarkdown(markdownId string, changeMap map[string]interface{}) error {
	var markdown dbModels.ArticleMarkdown
	return dao.db.Model(markdown).Where("id = ?", markdownId).UpdateColumns(changeMap).Error
}

func (dao ArticleMarkdownDao) DeleteArticleMarkdown(markdownId string) error {
	var markdown dbModels.ArticleMarkdown
	return dao.db.Where("id = ?", markdownId).Unscoped().Delete(&markdown).Error
}
