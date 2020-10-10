package logout

import (
	"compose/commons"
	"compose/user/daos"
	"errors"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	commonModel := commons.GetCommonModel(request)
	accessTokenDao := daos.GetAccessTokenDao()
	err := accessTokenDao.DeleteAccessTokenEntry(commonModel.AccessToken)
	if commons.InError(err) {
		commons.WriteFailedResponse(errors.New("Error deleting access token entry"), writer)
		return
	}
	responseModel := ResponseModel{
		Status:  commons.NewResponseStatus().SUCCESS,
		Message: "Succcessfully logged out",
	}
	commons.WriteSuccessResponse(responseModel, writer)
}

type ResponseModel struct {
	Status  commons.ResponseStatus `json:"status,omitempty"`
	Message string                 `json:"message,omitempty"`
}
