package logger

import (
	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger
var RequestResponseLogger *zerolog.Logger

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
