package echo

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/log"
)

func EchoLogger(eCtx echo.Context) log.Logger {
	request := eCtx.Request()

	logger := log.FromContext(eCtx.Request().Context())

	return logger.With(
		"host", request.Host,
		"path", request.RequestURI,
		"method", request.Method,
	)
}
