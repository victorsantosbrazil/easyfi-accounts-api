package echo

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = HttpErrorHandler

	e.Use(middleware.LoggerContextMiddleware)
	e.Use(middleware.TraceMiddleware)

	serveSwagger(e)

	return e
}

func serveSwagger(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
