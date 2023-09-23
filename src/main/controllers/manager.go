package controllers

import v1 "github.com/victorsantosbrazil/financial-institutions-api/src/main/controllers/v1"

type ControllerManager struct {
	institutionsController *v1.InstitutionsController
}

func NewControllerManager(institutionsController *v1.InstitutionsController) *ControllerManager {
	return &ControllerManager{
		institutionsController: institutionsController,
	}
}
