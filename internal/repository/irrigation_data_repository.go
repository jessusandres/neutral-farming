package repository

import "looker.com/neutral-farming/internal/domain"

type IrrigationDataRepository interface {
	FindByID(irrigationID uint) (*domain.IrrigationData, error)
}
