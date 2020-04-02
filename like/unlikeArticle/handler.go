package unlikeArticle

import (
	"compose/commons"
	"compose/like/likeCommons"
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

	articleUserId := likeCommons.ArticleServiceContract.GetArticleAuthorId(requestModel.ArticleId)
	if articleUserId == nil {
		writeFailedResponse(errors.New("No such article id exists"), writer)
		return
	}
	err = securityClearance(requestModel, articleUserId)
	if commons.InError(err) {
		writer.WriteHeader(http.StatusForbidden)
		writeFailedResponse(err, writer)
		return
	}

	err = unlikeArticle(requestModel)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	response := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Article unliked successfully",
	}

	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)
}

func securityClearance(model *RequestModel, articleUserId *string) error {
	if model.CommonModel.UserId == *articleUserId {
		return errors.New("You cannot unlike your own article")
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
