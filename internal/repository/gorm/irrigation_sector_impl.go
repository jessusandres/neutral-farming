package gorm

import (
	"gorm.io/gorm"

	"looker.com/neutral-farming/internal/domain"
	"looker.com/neutral-farming/internal/model"
	"looker.com/neutral-farming/internal/repository"
)

type IrrigationSectorRepo struct {
	db *gorm.DB
}

func NewIrrigationSectorRepo(db *gorm.DB) repository.IrrigationSectorRepository {
	return &IrrigationSectorRepo{db}
}

func (repo *IrrigationSectorRepo) FindByID(id uint) (*domain.IrrigationSector, error) {
	instance := model.IrrigationSector{}

	result := repo.db.First(&instance, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &domain.IrrigationSector{
		ID:     instance.ID,
		Name:   instance.Name,
		FarmID: instance.FarmID,
	}, nil
}
