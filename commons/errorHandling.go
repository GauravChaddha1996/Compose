package commons

import (
	"compose/commons/logger"
)

func InError(err error) bool {
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		return true
	}
	return false
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
