package delete

import (
	"compose/article/articleCommons"
	"compose/article/daos"
	"compose/commons"
	"encoding/json"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	article, err := daos.GetArticleDao().GetArticle(requestModel.ArticleId)
	if commons.InError(err) {
		writeFailedResponse(errors.New("No such article id exists"), writer)
		return
	}
	err = securityClearance(requestModel, article)
	if commons.InError(err) {
		writer.WriteHeader(http.StatusForbidden)
		writeFailedResponse(err, writer)
		return
	}

	err = deleteArticle(article)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Article deleted successfully",
	}

	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)

}

func securityClearance(model *RequestModel, article *articleCommons.Article) error {
	if model.CommonModel.UserId != article.UserId {
		return errors.New("Article not posted by this user")
	}
	return nil
}

func writeFailedResponse(err error, writer http.ResponseWriter) {
	failedResponse := ResponseModel{
		Status:  commons.NewResponseStatus().FAILED,
		Message: err.Error(),
	}
	failedResponseJson, err := json.Marshal(failedResponse)
	commons.PanicIfError(err)
	_, err = writer.Write(failedResponseJson)
	commons.PanicIfError(err)
}
