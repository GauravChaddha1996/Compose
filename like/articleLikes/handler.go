package articleLikes

import (
	"compose/commons"
	"compose/like/likeCommons"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}

	articleExists := likeCommons.ArticleServiceContract.DoesArticleExist(requestModel.ArticleId)
	if articleExists == false {
		commons.WriteFailedResponse(errors.New("No such article id exists"), writer)
		return
	}

	response, err := getArticleLikesResponse(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
		return
	}
	commons.WriteSuccessResponse(response, writer)
}
