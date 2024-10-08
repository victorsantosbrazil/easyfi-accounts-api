package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/app/usecase"
	"github.com/victorsantosbrazil/easyfi-accounts-api/src/common/app/model/pagination"
)

const _CONTROLLER_PATH = "/institutions"

type InstitutionController struct {
	*echo.Group
	listUseCase usecase.ListInstitutionsUseCase
}

// @Tags institutions
// @Summary list institutions
// @Router /v1/institutions [get]
// @Produce json
// @Param page query int false " "
// @Param size query int false " "
// @Success 200 {object} usecase.ListInstitutionsUseCaseResponse
// @Failure 400  {object}  errors.ApiError
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
