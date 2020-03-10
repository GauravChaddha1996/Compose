package user

import "compose/commons"

const (
	// todo convert this all to enums
	ERROR_PARSE_MESSAGE = "Signup request cannot be parsed."

	ERROR_VALIDITY_NAME_MESSAGE     = "Signup request not valid. Name should be between 1 and 255 characters."
	ERROR_VALIDITY_EMAIL_MESSAGE    = "Signup request not valid. Email isn't valid."
	ERROR_VALIDITY_PASSWORD_MESSAGE = "Signup request not valid. Password isn't valid."

	REQUEST_MODEL_NAME     = "name"
	REQUEST_MODEL_EMAIL    = "email"
	REQUEST_MODEL_PASSWORD = "password"

	ERROR_USER_EMAIL_ALREADY_PRESENT       = "Signup request error: Email already present."
	ERROR_USDERID_GENERATION_FAILURE       = "Signup request error: User id cannot be generated"
	ERROR_USER_DB_SAVE_FAILURE             = "Signup request error: User cannot be saved"
	ERROR_USER_PASSWORD_GENERATION_FAILURE = "Signup request error: Password cannot be generated"
	ERROR_USER_PASSWORD_SAVE_FAILURE       = "Signup request error: Password cannot be saved"
	ERROR_ACCESS_TOKEN_GENERATION_FAILURE  = "Signup request error: Access token cannot be generated"
	ERROR_ACCESS_TOKEN_SAVE_FAILURE        = "Signup request error: Access token cannot be saved"
)

type SignupRequestModel struct {
	Email    string
	Name     string
	Password string
}

// todo add rules here for null
type SignupResponseModel struct {
	Status      commons.ResponseStatus `json:"Status"`
	Message     string                 `json:"Message"`
	AccessToken string                 `json:"AccessToken"`
}
