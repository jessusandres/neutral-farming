package service

import (
	"looker.com/neutral-farming/internal/repository/gorm"
)

type Services struct {
	FarmService       *FarmService
	SectorService     *SectorService
	IrrigationService *IrrigationService
}

func NewServices(repos *gorm.Repositories) *Services {
	return &Services{
		FarmService:       NewFarmService(repos.Farm),
		SectorService:     NewSectorService(repos.IrrigationSector),
		IrrigationService: NewIrrigationService(repos.IrrigationData),
	}
}
