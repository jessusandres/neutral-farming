package gorm

import (
	"gorm.io/gorm"
	"looker.com/neutral-farming/internal/repository"
)

type Repositories struct {
	Farm             repository.FarmRepository
	IrrigationSector repository.IrrigationSectorRepository
	IrrigationData   repository.IrrigationDataRepository
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Farm:             NewFarmRepo(db),
		IrrigationSector: NewIrrigationSectorRepo(db),
		IrrigationData:   NewIrrigationDataRepo(db),
	}
}
