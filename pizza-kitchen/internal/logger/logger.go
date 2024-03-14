package logger

import "go.uber.org/zap"


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
