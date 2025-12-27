package controller

import "looker.com/neutral-farming/internal/service"

type Controllers struct {
	FarmController       *FarmController
	SectorController     *SectorController
	IrrigationController *IrrigationController
}

func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		FarmController:       NewFarmController(services.FarmService),
		SectorController:     NewSectorController(services.SectorService),
		IrrigationController: NewIrrigationController(services.IrrigationService),
	}
}
