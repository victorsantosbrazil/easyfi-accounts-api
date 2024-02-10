package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/log"
)

func EchoLogger(ctx echo.Context) log.Logger {
	return log.NewLogger().With(
		"host", ctx.Request().Host,
		"path", ctx.Request().RequestURI,
		"method", ctx.Request().Method,
	)
}
