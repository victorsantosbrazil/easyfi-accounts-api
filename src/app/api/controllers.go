package api

type Controllers struct {
	institutionsController *InstitutionController
}

func NewControllers(institutionsController *InstitutionController) *Controllers {
	return &Controllers{
		institutionsController: institutionsController,
	}
}
