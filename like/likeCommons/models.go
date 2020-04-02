package likeCommons

type LikeEntry struct {
	Id        uint64
	UserId    string
	ArticleId string
}

func (LikeEntry) TableName() string {
	return "likes"
}
