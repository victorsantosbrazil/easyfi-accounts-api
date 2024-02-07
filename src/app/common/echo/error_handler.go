package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/errors"
)

const (
	_NOT_FOUND_ERROR = "Not found"
)

func HttpErrorHandler(err error, c echo.Context) {
	c.Logger().Error(err)

	switch err.(type) {
	case *echo.HTTPError:
		handleEchoHttpError(err.(*echo.HTTPError), c)
		return
	case errors.ApiError:
		handleApiError(err.(errors.ApiError), c)
		return
	default:
		handleError(err, c)
	}
}

func handleEchoHttpError(err *echo.HTTPError, c echo.Context) {
	switch err.Code {
	case http.StatusNotFound:
		c.JSON(err.Code, errors.NotFoundError(_NOT_FOUND_ERROR))
		return
	default:
		c.JSON(http.StatusInternalServerError, errors.InternalServerError())
	}
}

func handleApiError(err errors.ApiError, c echo.Context) {
	c.JSON(err.Status, err)
}

func handleError(err error, c echo.Context) {
	c.JSON(http.StatusInternalServerError, errors.InternalServerError())
}
