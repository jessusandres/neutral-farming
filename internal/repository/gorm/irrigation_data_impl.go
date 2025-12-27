package gorm

import (
	"gorm.io/gorm"
	
	"looker.com/neutral-farming/internal/domain"
	"looker.com/neutral-farming/internal/model"
	"looker.com/neutral-farming/internal/repository"
)

type IrrigationDataRepo struct {
	db *gorm.DB
}

func NewIrrigationDataRepo(db *gorm.DB) repository.IrrigationDataRepository {
	return &IrrigationDataRepo{db}
}

func (repository *IrrigationDataRepo) FindByID(irrigationID uint) (*domain.IrrigationData, error) {
	instance := &model.IrrigationData{}

	result := repository.db.First(instance, irrigationID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.IrrigationData{
		ID:                 instance.ID,
		FarmID:             instance.FarmID,
		IrrigationSectorID: instance.IrrigationSectorID,
		StartTime:          instance.StartTime,
		EndTime:            instance.EndTime,
		NominalAmount:      instance.NominalAmount,
		RealAmount:         instance.RealAmount,
	}, nil
}
