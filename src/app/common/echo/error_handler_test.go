package echo

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	cmnErrors "github.com/victorsantosbrazil/financial-institutions-api/src/app/common/errors"
)

func TestErrorHandler(t *testing.T) {

	t.Run("should just return api errors when handling api errors", func(t *testing.T) {
		err := cmnErrors.BadRequestError("bad request")
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, rec)

		HttpErrorHandler(err, ctx)

		var actualErr cmnErrors.ApiError
		json.NewDecoder(rec.Body).Decode(&actualErr)
		assert.Equal(t, err, actualErr)
	})

	t.Run("should return internal server error when handling generic errors", func(t *testing.T) {
		err := errors.New("generic error")
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, rec)

		HttpErrorHandler(err, ctx)

		expectedErr := cmnErrors.InternalServerError()
		var actualErr cmnErrors.ApiError
		json.NewDecoder(rec.Body).Decode(&actualErr)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("should return not found error when handling echo http errors with status codes not found", func(t *testing.T) {
		err := echo.NewHTTPError(404, "not found error")
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, rec)

		HttpErrorHandler(err, ctx)

		expectedErr := cmnErrors.NotFoundError(_NOT_FOUND_ERROR)
		var actualErr cmnErrors.ApiError
		json.NewDecoder(rec.Body).Decode(&actualErr)
		assert.Equal(t, expectedErr, actualErr)
	})

	t.Run("should return internal server error when handling echo http errors with status codes not handled", func(t *testing.T) {
		err := echo.NewHTTPError(500, "generic error")
		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, rec)

		HttpErrorHandler(err, ctx)

		expectedErr := cmnErrors.InternalServerError()
		var actualErr cmnErrors.ApiError
		json.NewDecoder(rec.Body).Decode(&actualErr)
		assert.Equal(t, expectedErr, actualErr)
	})
}
