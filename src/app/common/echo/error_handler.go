package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/errors"
)

const _INTERNAL_SERVER_ERROR = "internal server error"

func ErrorHandler(err error, c echo.Context) {
	if apiErr, ok := err.(errors.ApiError); ok {
		c.JSON(apiErr.Status, apiErr)
		return
	}
	c.JSON(http.StatusInternalServerError, errors.InternalServerError(_INTERNAL_SERVER_ERROR))
}
