package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/log"
)

func LoggerContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := log.NewLogger()
		ctx := log.LoggerContext(c.Request().Context(), logger)
		request := c.Request().WithContext(ctx)
		c.SetRequest(request)
		return next(c)
	}
}
