package log

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromContext(t *testing.T) {
	t.Run("should return context logger when context has logger", func(t *testing.T) {
		ctx := context.Background()
		expectedLogger := NewLogger().With("traceId", "123")
		ctx = LoggerContext(ctx, expectedLogger)
		actualLogger := FromContext(ctx)
		assert.Equal(t, expectedLogger, actualLogger)
	})

	t.Run("should return default logger when context does not have logger", func(t *testing.T) {
		ctx := context.Background()
		expectedLogger := NewLogger()
		actualLogger := FromContext(ctx)
		assert.Equal(t, expectedLogger, actualLogger)
	})
}
