package log

import "context"

const (
	_CONTEXT_KEY = "LOGGER"
)

func LoggerContext(parentCtx context.Context, logger Logger) context.Context {
	return context.WithValue(parentCtx, _CONTEXT_KEY, logger)
}

func FromContext(ctx context.Context) Logger {
	if logger, ok := ctx.Value(_CONTEXT_KEY).(Logger); ok {
		return logger
	}
	return NewLogger()
}
