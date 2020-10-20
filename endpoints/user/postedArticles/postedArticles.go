package postedArticles

import (
	"compose/commons"
	"compose/dataLayer/daos"
	"errors"
)

func getPostedArticles(model *RequestModel) (*ResponseModel, error) {
	postedArticleLimit := 3
	articleDao := daos.GetArticleDao()
	articleArr, err := articleDao.GetArticlesOfUser(model.UserId, *model.MaxCreatedAt, postedArticleLimit)
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
			Status:          commons.NewResponseStatus().SUCCESS,
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
	lastCreatedAt := (*articleArr)[articleArrLen-1].CreatedAt.Format(commons.TimeFormat)
	return &ResponseModel{
		Status:          commons.NewResponseStatus().SUCCESS,
		PostedArticles:  postedArticleArr,
		MaxCreatedAt:    lastCreatedAt,
		HasMoreArticles: !(articleArrLen < postedArticleLimit),
	}, nil
}
