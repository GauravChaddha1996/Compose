package articleComments

import (
	"compose/comments/commentCommons"
	"compose/comments/daos"
	"compose/commons"
	"compose/dbModels"
	"errors"
	"fmt"
	"sync"
	"time"
)

func getArticleCommentsResponse(model *RequestModel) (*ResponseModel, error) {
	commentDao := daos.GetCommentDao()
	markdownDao := daos.GetCommentMarkdownDao()

	var articleRootCommentsLimit = 3
	rootComments, _ := commentDao.GetNextRootCommentsOfArticle(model.ArticleId, model.MaxCreatedAt, articleRootCommentsLimit)
	rootCommentsLen := len(*rootComments)
	if rootCommentsLen == 0 {
		return getNoCommentsResponseModel(model), nil
	}

	commentEntityArr := make([]commentCommons.CommentEntity, rootCommentsLen)

	var rootCommentWaitGroup sync.WaitGroup
	var err error
	err = nil
	for index, comment := range *rootComments {
		rootCommentWaitGroup.Add(1)
		go func(index int, comment dbModels.Comment) {
			postedByUser, err1 := getPostedByUser(comment.UserId)
			markdown, err2 := getCommentMarkdown(comment.MarkdownId, markdownDao)
			if commons.InError(err1) {
				err = err1
			} else if commons.InError(err2) {
				err = err2
			} else {
				commentEntityArr[index] = commentCommons.CommentEntity{
					CommentId:     comment.CommentId,
					Markdown:      markdown,
					PostedByUser:  *postedByUser,
					PostedAt:      getPostedAtTime(comment),
					ChildComments: nil,
				}
			}
			rootCommentWaitGroup.Done()
		}(index, comment)
	}
	rootCommentWaitGroup.Wait()
	if err != nil {
		return nil, err
	}
	lastCreatedAt := (*rootComments)[rootCommentsLen-1].CreatedAt.Format("2 Jan 2006 15:04:05")
	return &ResponseModel{
		Status:          commons.NewResponseStatus().SUCCESS,
		Comments:        commentEntityArr,
		MaxCreatedAt:    lastCreatedAt,
		HasMoreComments: !(rootCommentsLen < articleRootCommentsLimit),
	}, nil
}

func getNoCommentsResponseModel(model *RequestModel) *ResponseModel {
	var message string
	if *model.MaxCreatedAt == model.DefaultMaxCreatedAt {
		message = "No comments to show"
	} else {
		message = "No more comments to show"
	}
	return &ResponseModel{
		Status:          commons.NewResponseStatus().SUCCESS,
		Message:         message,
		HasMoreComments: false,
	}
}

func getPostedByUser(userId string) (*commentCommons.PostedByUser, error) {
	user, err := commentCommons.UserServiceContract.GetUser(userId)
	if commons.InError(err) {
		return nil, errors.New("Error in fetching user details")
	}
	return &commentCommons.PostedByUser{
		UserId:   user.UserId,
		Name:     user.Name,
		PhotoUrl: user.PhotoUrl,
	}, nil
}

func getCommentMarkdown(markdownId string, dao *daos.CommentMarkdownDao) (string, error) {
	markdown, err := dao.GetCommentMarkdown(markdownId)
	if commons.InError(err) {
		return "", errors.New("Error in fetching markdown for comment")
	}
	return markdown.Markdown, nil
}

func getPostedAtTime(comment dbModels.Comment) string {
	var t time.Time
	var postedAt string
	if comment.UpdatedAt != nil {
		t = *comment.UpdatedAt
		postedAt = fmt.Sprint("Updated at ")
	} else {
		t = comment.CreatedAt
		postedAt = fmt.Sprint("Created at ")
	}
	return fmt.Sprint(postedAt, t.Hour(), ":", t.Minute(), ":", t.Second(), t.Day(), t.Month(), t.Year())
}
