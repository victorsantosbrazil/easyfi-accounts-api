package middleware

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/financial-institutions-api/src/common/app/log"
)

func TestTraceMiddleware(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockLogger := log.NewMockLogger(mockCtrl)

	ctx := log.LoggerContext(context.Background(), mockLogger)
	request := httptest.NewRequest("GET", "/", nil).WithContext(ctx)
	recorder := httptest.NewRecorder()
	eCtx := echo.New().NewContext(request, recorder)

	handleFunc := TraceMiddleware(func(c echo.Context) error {
		return nil
	})

	mockLogger.EXPECT().With(_TRACE_ID_KEY, gomock.Any()).Return(mockLogger)

	err := handleFunc(eCtx)

	if assert.NoError(t, err) {
		actualLogger := log.FromContext(eCtx.Request().Context())
		assert.Equal(t, mockLogger, actualLogger)
	}
}
