package api

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/common/domain/model/pagination"
	"github.com/victorsantosbrazil/financial-institutions-api/src/app/domain/usecase"
)

const _CONTROLLER_PATH = "/financial-institutions"

type InstitutionController struct {
	*echo.Group
	listUseCase usecase.ListInstitutionsUseCase
}

func (c *InstitutionController) list(eCtx echo.Context) error {
	ctx := eCtx.Request().Context()

	pageParams, err := pagination.NewPageParams(eCtx.QueryParams())
	if err != nil {
		return err
	}

	response, err := c.listUseCase.Run(ctx, pageParams)
	if err != nil {
		return err
	}

	return eCtx.JSON(200, response)
}

func NewInstitutionController(v1 V1Group, listUseCase usecase.ListInstitutionsUseCase) *InstitutionController {

	c := &InstitutionController{
		Group:       v1.Group(_CONTROLLER_PATH),
		listUseCase: listUseCase,
	}

	c.GET("", c.list)

	return c
}
