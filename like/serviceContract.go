package like

import (
	"compose/dbModels"
	"compose/like/daos"
	"gorm.io/gorm"
	"time"
)

type ServiceContractImpl struct {
	dao daos.LikeDao
}

func GetServiceContractImpl() ServiceContractImpl {
	return ServiceContractImpl{dao: daos.GetLikeDao()}
}

func (impl ServiceContractImpl) GetAllLikeEntriesOfUser(userId string, maxCreatedAtTime time.Time, limit int) (*[]dbModels.LikeEntry, error) {
	return impl.dao.GetUserLikes(userId, maxCreatedAtTime, limit)
}
func (impl ServiceContractImpl) DeleteAllLikeEntriesOfArticle(articleId string, transaction *gorm.DB) error {
	return daos.GetLikeDaoDuringTransaction(transaction).DeleteAllLikesOfArticle(articleId)
}
