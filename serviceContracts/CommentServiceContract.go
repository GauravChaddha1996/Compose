package serviceContracts

import "gorm.io/gorm"

type CommentServiceContract interface {
	DeleteAssociatedCommentsAndReplies(articleId string, transaction *gorm.DB) error
}
