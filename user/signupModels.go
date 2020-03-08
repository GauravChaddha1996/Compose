package user

import "compose/commons"

const (
	SIGNUP_ERROR_SECURITY_MESSAGE = "Signup request not secure."
	SIGNUP_ERROR_PARSE_MESSAGE    = "Signup request cannot be parsed."

	SIGNUP_ERROR_VALIDITY_NAME_MESSAGE  = "Signup request not valid. Name cannot be empty."
	SIGNUP_ERROR_VALIDITY_EMAIL_MESSAGE = "Signup request not valid. Email isn't valid."

	SIGNUP_REQUEST_MODEL_NAME     = "name"
	SIGNUP_REQUEST_MODEL_EMAIL    = "email"
	SIGNUP_REQUEST_MODEL_PASSWORD = "password"

	SIGNUP_ERROR_USER_EMAIL_ALREADY_PRESENT      = "Signup request error: Email already present."
	SIGNUP_ERROR_USDERID_GENERATION_FAILURE      = "Signup request error: User id cannot be generated"
	SIGNUP_ERROR_USER_DB_SAVE_FAILURE      = "Signup request error: User cannot be saved"
	SIGNUP_ERROR_ACCESS_TOKEN_GENERATION_FAILURE = "Signup request error: Access token cannot be generated"
)

type UserSignupRequestModel struct {
	Email string
	Name  string
}

type UserSignupResponseModel struct {
	Status      commons.ResponseStatus `json:"Status"`
	Message     string                 `json:"Message"`
	AccessToken string                 `json:"AccessToken"`
}
