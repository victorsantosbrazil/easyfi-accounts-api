package middleware

import (
	"math/rand"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/log"
)

const (
	_TRACE_ID_KEY = "traceId"
)

func TraceMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logTraceId(c)
		return next(c)
	}
}

func logTraceId(c echo.Context) {
	logger := log.FromContext(c.Request().Context()).With(_TRACE_ID_KEY, rand.Int())
	ctx := log.LoggerContext(c.Request().Context(), logger)
	request := c.Request().WithContext(ctx)
	c.SetRequest(request)
}
