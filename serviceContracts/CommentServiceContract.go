package serviceContracts

import "github.com/jinzhu/gorm"

type CommentServiceContract interface {
	DeleteAssociatedCommentsAndReplies(articleId string, transaction *gorm.DB) error
}
