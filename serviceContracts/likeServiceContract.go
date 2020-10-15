package serviceContracts

import (
	"compose/dbModels"
	"github.com/jinzhu/gorm"
	"time"
)

type LikeServiceContract interface {
	GetAllLikeEntriesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.LikeEntry, error)
	DeleteAllLikeEntriesOfArticle(articleId string, transaction *gorm.DB) error
}
