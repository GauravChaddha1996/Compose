package postedArticles

import (
	"compose/commons"
	"compose/user/userCommons"
	"errors"
)

func getPostedArticles(model *RequestModel) (*ResponseModel, error) {
	postedArticleLimit := 3
	articleArr, err := userCommons.ArticleService.GetAllArticlesOfUser(model.CommonModel.UserId, *model.MaxCreatedAt, postedArticleLimit)
	if commons.InError(err) {
		return nil, errors.New("Can't fetch posted articles")
	}
	articleArrLen := len(*articleArr)
	if articleArrLen == 0 {
		var message string
		if *model.MaxCreatedAt == model.DefaultMaxCreatedAt {
			message = "So empty... No posted articles."
		} else {
			message = "No more articles to show"
		}
		return &ResponseModel{
			Status:          commons.ResponseStatusWrapper{}.SUCCESS,
			Message:         message,
			HasMoreArticles: false,
		}, nil
	}

	var postedArticleArr = make([]PostedArticle, articleArrLen)
	for index, article := range *articleArr {
		postedArticleArr[index] = PostedArticle{
			Id:          article.Id,
			Title:       article.Title,
			Description: article.Description,
		}
	}
	lastCreatedAt := (*articleArr)[articleArrLen-1].CreatedAt.Format("2 Jan 2006 15:04:05")
	return &ResponseModel{
		Status:          commons.ResponseStatusWrapper{}.SUCCESS,
		PostedArticles:  postedArticleArr,
		MaxCreatedAt:    lastCreatedAt,
		HasMoreArticles: !(articleArrLen < postedArticleLimit),
	}, nil
}
