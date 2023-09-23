package v1

import "github.com/labstack/echo/v4"

var _CONTROLLER_PATH = "/financial-institutions"

type InstitutionsController struct {
	*echo.Group
}

func (c *InstitutionsController) list(ec echo.Context) error {
	return ec.String(200, "Hello world")
}

func NewInstitutionsController(v1 V1Group) *InstitutionsController {

	c := &InstitutionsController{
		Group: (*echo.Group)(v1).Group(_CONTROLLER_PATH),
	}

	c.GET("/", c.list)

	return c
}
