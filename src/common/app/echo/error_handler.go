package echo

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/errors"
)

const (
	_NOT_FOUND_ERROR = "Not found"
)

func HttpErrorHandler(err error, ctx echo.Context) {
	EchoLogger(ctx).Error(err.Error())

	switch err.(type) {
	case *echo.HTTPError:
		handleEchoHttpError(err.(*echo.HTTPError), ctx)
		return
	case errors.ApiError:
		handleApiError(err.(errors.ApiError), ctx)
		return
	default:
		handleError(err, ctx)
	}
}

func handleEchoHttpError(err *echo.HTTPError, ctx echo.Context) {
	switch err.Code {
	case http.StatusNotFound:
		ctx.JSON(err.Code, errors.NotFoundError(_NOT_FOUND_ERROR))
		return
	default:
		ctx.JSON(http.StatusInternalServerError, errors.InternalServerError())
	}
}

func handleApiError(err errors.ApiError, ctx echo.Context) {
	ctx.JSON(err.Status, err)
}

func handleError(err error, ctx echo.Context) {
	ctx.JSON(http.StatusInternalServerError, errors.InternalServerError())
}
