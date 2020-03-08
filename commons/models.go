package commons

type ResponseStatus string

const (
	//todo convert this inside an enum
	RESPONSE_STATUS_SUCCESS ResponseStatus = "success"
	RESPONSE_STATUS_FAILED  ResponseStatus = "failed"
)