package logger

import "go.uber.org/zap"

type Logger struct {
	zap.SugaredLogger
	zap.Logger
}

func New() *Logger {
	logger := zap.NewExample()
	return &Logger{
		*logger.Sugar(),
		*logger,
	}
}

