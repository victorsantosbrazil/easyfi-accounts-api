package echo

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func New() *echo.Echo {
	e := echo.New()

	e.HTTPErrorHandler = HttpErrorHandler

	serveSwagger(e)

	return e
}

func serveSwagger(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
