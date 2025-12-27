package repository

import "looker.com/neutral-farming/internal/domain"

type IrrigationSectorRepository interface {
	FindByID(ID uint) (*domain.IrrigationSector, error)
}
