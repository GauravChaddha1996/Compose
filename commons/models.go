package commons

const CommonModelKey = "common_model"

type ResponseStatus string

type ResponseStatusWrapper struct {
	SUCCESS ResponseStatus
	FAILED  ResponseStatus
}

type GenericErrorResponseModel struct {
	Status  ResponseStatus `json:"status,omitempty"`
	Message string         `json:"message,omitempty"`
}

type CommonRequestModel struct {
	AccessToken string
	UserId      string
	UserEmail   string
}

type EndpointSecurityConfig struct {
	CheckAccessToken bool
	CheckUserId      bool
	CheckUserEmail   bool
}