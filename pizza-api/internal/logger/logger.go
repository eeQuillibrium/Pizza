package logger

import "go.uber.org/zap"

const (
	DEBUG      = "debug"
	PRODUCTION = "production"
)

type Logger struct {
	zap.Logger
	zap.SugaredLogger
}

func New() *Logger {
	logger := zap.NewExample()
	return &Logger{
		*logger,
		*logger.Sugar(),
	}
}
