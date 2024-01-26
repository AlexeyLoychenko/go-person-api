package logger

import (
	"log/slog"
	"os"
)

type Logger interface {
	Debug(msg string, keyvals ...interface{})
	Info(msg string, keyvals ...interface{})
	Error(msg string, keyvals ...interface{})
	With(keyvals ...interface{}) LoggerWrapper
}

type LoggerWrapper struct {
	logger *slog.Logger
}

func NewSlogWrapper(logLevel string, handlerType string) LoggerWrapper {
	var logger *slog.Logger
	slogLevel := slog.LevelInfo
	switch logLevel {
	case "debug":
		slogLevel = slog.LevelDebug
	case "info":
		slogLevel = slog.LevelInfo
	case "error":
		slogLevel = slog.LevelError
	}
	logger = slog.New(NewSlogHandler(handlerType, slogLevel))
	return LoggerWrapper{logger: logger}
}

func NewSlogHandler(handlerType string, level slog.Level) slog.Handler {
	if handlerType == "json" {
		return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
}

func (lw LoggerWrapper) Debug(msg string, args ...interface{}) {
	lw.logger.Debug(msg, args...)
}

func (lw LoggerWrapper) Info(msg string, args ...interface{}) {
	lw.logger.Info(msg, args...)
}

func (lw LoggerWrapper) Error(msg string, args ...interface{}) {
	lw.logger.Error(msg, args...)
}

func (lw LoggerWrapper) With(args ...interface{}) LoggerWrapper {
	lw.logger = lw.logger.With(args...)
	return lw
}
