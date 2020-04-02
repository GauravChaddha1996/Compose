package articleLikes

import (
	"compose/commons"
	"compose/like/likeCommons"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	requestModel, err := getRequestModel(request)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}

	articleExists := likeCommons.ArticleServiceContract.DoesArticleExist(requestModel.ArticleId)
	if articleExists == false {
		writeFailedResponse(errors.New("No such article id exists"), writer)
		return
	}

	likedByUserArr, lastLikeId, err := getArticleLikeUserList(requestModel)
	if commons.InError(err) {
		writeFailedResponse(err, writer)
		return
	}
	var response ResponseModel
	if likedByUserArr == nil {
		response = ResponseModel{
			Status:       commons.NewResponseStatus().SUCCESS,
			Message:      "No more likes to show",
			LikedByUsers: []LikedByUser{},
		}
	} else {
		response = ResponseModel{
			Status:       commons.NewResponseStatus().SUCCESS,
			LikedByUsers: *likedByUserArr,
			LastLikeId:   strconv.FormatUint(*lastLikeId, 10),
		}
	}

	jsonResponse, err := json.Marshal(response)
	commons.PanicIfError(err)
	_, err = writer.Write(jsonResponse)
	commons.PanicIfError(err)
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
