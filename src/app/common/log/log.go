package log

import (
	"log/slog"
	"os"
)

type Logger interface {
	With(args ...any) Logger
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
}

type loggerImpl struct {
	*slog.Logger
}

func (l *loggerImpl) With(args ...any) Logger {
	return &loggerImpl{
		Logger: l.Logger.With(args...),
	}
}

func (l *loggerImpl) Fatal(msg string, args ...any) {
	l.Error(msg, args...)
	os.Exit(1)
}

func NewLogger() Logger {
	baseLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return &loggerImpl{
		Logger: baseLogger,
	}
}
