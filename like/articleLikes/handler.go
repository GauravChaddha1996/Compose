package articleLikes

import (
	"compose/commons"
	"compose/like/likeCommons"
	"errors"
	"net/http"
	"strconv"
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

	likedByUserArr, lastLikeId, err := getArticleLikeUserList(requestModel)
	if commons.InError(err) {
		commons.WriteFailedResponse(err, writer)
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

	commons.WriteSuccessResponse(response, writer)
}