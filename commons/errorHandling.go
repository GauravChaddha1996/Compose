package commons

import (
	"compose/commons/logger"
	"github.com/rs/zerolog"
)

func InError(err error) bool {
	if err != nil {
		logger.Logger.Error().Msg(err.Error())
		return true
	}
	return false
}

func InError2(err error, myLogger *zerolog.Logger) bool {
	if myLogger == nil {
		myLogger = logger.Logger
	}
	if err != nil {
		myLogger.Error().Msg(err.Error())
		return true
	}
	return false
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
