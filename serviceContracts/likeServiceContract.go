package serviceContracts

import (
	"compose/dbModels"
	"gorm.io/gorm"
	"time"
)

type LikeServiceContract interface {
	GetAllLikeEntriesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.LikeEntry, error)
	DeleteAllLikeEntriesOfArticle(articleId string, transaction *gorm.DB) error
}
