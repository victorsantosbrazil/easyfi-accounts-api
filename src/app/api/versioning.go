package api

import "github.com/labstack/echo/v4"

type V1Group = *echo.Group

func NewV1Group(e *echo.Echo) V1Group {
	return e.Group("/v1")
}
