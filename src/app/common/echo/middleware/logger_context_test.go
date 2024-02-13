package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/log"
)

func TestLoggerContextMiddleware(t *testing.T) {
	expectedLogger := log.NewLogger()

	handleFunc := LoggerContextMiddleware(func(c echo.Context) error {
		return nil
	})

	request := httptest.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()
	eCtx := echo.New().NewContext(request, recorder)

	err := handleFunc(eCtx)

	if assert.NoError(t, err) {
		actualLogger := log.FromContext(eCtx.Request().Context())
		assert.Equal(t, expectedLogger, actualLogger)
	}
}
