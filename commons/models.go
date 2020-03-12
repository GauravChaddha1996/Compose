package commons

type ResponseStatus string

type ResponseStatusWrapper struct {
	SUCCESS ResponseStatus
	FAILED  ResponseStatus
}

func NewResponseStatus() ResponseStatusWrapper {
	return ResponseStatusWrapper{
		SUCCESS: "success",
		FAILED:  "failed",
	}
}
