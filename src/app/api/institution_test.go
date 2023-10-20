package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	cmnmodel "github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/usecase"
)

func TestListInstitutions(t *testing.T) {
	e := echo.New()
	g := e.Group("")

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	listUseCase := usecase.NewMockListInstitutionsUseCase(mockCtrl)
	institutionController := NewInstitutionController(g, listUseCase)

	req := httptest.NewRequest(http.MethodGet, _CONTROLLER_PATH, nil)
	rec := httptest.NewRecorder()
	eCtx := e.NewContext(req, rec)

	page := 2
	size := 10

	queryParams := eCtx.QueryParams()
	queryParams.Add("page", strconv.Itoa(page))
	queryParams.Add("size", strconv.Itoa(size))

	pageRequest := cmnmodel.PageRequest{Page: page, Size: size}
	expected := usecase.ListInstitutionsUseCaseResponse{Pagination: cmnmodel.Pagination{}, Items: []usecase.ListInstitutionsUseCaseResponseItem{
		{Id: 1, Name: "Nubank"},
		{Id: 2, Name: "Brazil"},
	}}

	listUseCase.EXPECT().Run(eCtx.Request().Context(), pageRequest).Return(expected)

	if assert.NoError(t, institutionController.list(eCtx)) {
		var actual usecase.ListInstitutionsUseCaseResponse
		err := json.NewDecoder(rec.Body).Decode(&actual)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, actual)
	}

}
