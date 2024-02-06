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

	t.Run("should just return api errors", func(t *testing.T) {
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

		expectedErr := cmnErrors.InternalServerError(_INTERNAL_SERVER_ERROR)
		var actualErr cmnErrors.ApiError
		json.NewDecoder(rec.Body).Decode(&actualErr)
		assert.Equal(t, expectedErr, actualErr)
	})

}
