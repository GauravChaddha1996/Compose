package serviceContracts

import (
	"compose/dbModels"
	"time"
)

type LikeServiceContract interface {
	GetAllLikeEntriesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.LikeEntry, error)
}
