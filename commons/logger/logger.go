package logger

import (
	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger
var RequestResponseLogger *zerolog.Logger

var FETCH = "fetch"
var ACTION = "action"
var USER_ID = "user_id"
var ARTICLE_ID = "article_id"

func InitLogger() {
	zerolog.TimeFieldFormat = loggerTimeFormat
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	baseLogger := zerolog.New(NewComposeConsoleWriter()).With().Timestamp().Logger()
	requestLogger := zerolog.New(NewComposeConsoleWriter(func(w *ComposeConsoleWriter) {
		w.FieldsInDifferentLines = true
	})).With().Timestamp().Logger()

	RequestResponseLogger = &requestLogger
	Logger = &baseLogger
}

func Info(msg string) {
	Logger.Info().Msg(msg)
}

func InfoPreNewLine(msg string) {
	println()
	Logger.Info().Msg(msg)
}

func InfoPostNewLine(msg string) {
	Logger.Info().Msg(msg)
	println()
}
