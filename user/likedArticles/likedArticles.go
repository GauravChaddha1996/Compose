package likedArticles

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
)

func getLikedArticles(model *RequestModel) (*ResponseModel, error) {
	likedArticleLimit := 3
	likeEntriesOfUser, err := userCommons.LikeService.GetAllLikeEntriesOfUser(model.UserId, *model.MaxLikedAt, likedArticleLimit)
	if commons.InError(err) {
		return nil, errors.New("Can't fetch posted articles")
	}
	likeEntriesLen := len(*likeEntriesOfUser)
	if likeEntriesLen == 0 {
		var message string
		if *model.MaxLikedAt == model.DefaultMaxLikedAt {
			message = "So empty... No liked articles."
		} else {
			message = "No more liked articles to show"
		}
		return &ResponseModel{
			Status:          commons.ResponseStatusWrapper{}.SUCCESS,
			Message:         message,
			HasMoreArticles: false,
		}, nil
	}

	var articleIdsArr = make([]string, likeEntriesLen)
	for index, likeEntry := range *likeEntriesOfUser {
		articleIdsArr[index] = likeEntry.ArticleId
	}

	articleArr, err := userCommons.ArticleService.GetAllArticles(articleIdsArr)
	if commons.InError(err) {
		return nil, errors.New("Can't fetch article details")
	}
	var likedArticleArr = make([]LikedArticle, likeEntriesLen)
	for index, article := range *articleArr {
		likedArticleArr[index] = LikedArticle{
			Id:          article.Id,
			Title:       article.Title,
			Description: article.Description,
		}
	}
	lastCreatedAt := (*likeEntriesOfUser)[likeEntriesLen-1].CreatedAt.Format("2 Jan 2006 15:04:05")
	return &ResponseModel{
		Status:          commons.ResponseStatusWrapper{}.SUCCESS,
		LikedArticles:   likedArticleArr,
		MaxLikedAt:      lastCreatedAt,
		HasMoreArticles: !(likeEntriesLen < likedArticleLimit),
	}, nil
}
